package internal

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/goast"
	"github.com/dave/dst/decorator/resolver/guess"
	"github.com/dave/dst/dstutil"
	"github.com/skamensky/printfdebug/internal/options"
	"go/parser"
	"go/token"
)

type funcInfo struct {
	FuncIdentifier          string
	Body                    *dst.BlockStmt
	MustHaveReturnStatement bool
	DoesHaveReturnStatement bool
}

var DecorationMessage = "// automatically added by printf-debugger. Do not change this comment. It is an identifier."

type funcInfoStack []*funcInfo

func (s funcInfoStack) Push(v *funcInfo) funcInfoStack {
	return append(s, v)
}

func (s funcInfoStack) Pop() (funcInfoStack, *funcInfo) {
	l := len(s)
	if l == 0 {
		panic(errors.New("attempted to access empty stack"))
	}
	return s[:l-1], s[l-1]
}

func (s funcInfoStack) Peek() *funcInfo {
	l := len(s)
	if l == 0 {
		panic(errors.New("attempted to access empty stack"))
	}
	return s[l-1]
}

func getPrintStatement(message string) *dst.ExprStmt {
	e := &dst.ExprStmt{
		X: &dst.CallExpr{
			Fun: &dst.Ident{Name: "Println", Path: "fmt"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: message,
				},
			},
		},
	}
	e.Decs.End.Append(DecorationMessage)
	return e
}

func newFuncInfo(functionName string, body *dst.BlockStmt, funcType *dst.FuncType) *funcInfo {

	funcInfo := &funcInfo{
		FuncIdentifier:          functionName,
		Body:                    body,
		MustHaveReturnStatement: funcType.Results != nil,
	}

	for _, nd := range funcInfo.Body.List {
		if _, ok := nd.(*dst.ReturnStmt); ok {
			funcInfo.DoesHaveReturnStatement = true
		}
	}
	return funcInfo
}

func (f *funcInfo) modifyFuncBody() {
	// handle adding entering print statements and leaving print statements for implicit returns
	lst := f.Body.List
	if len(lst) == 0 {
		enterLeaveFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Entering and leaving empty function \"%v\"\n"`, f.FuncIdentifier))
		lst = append(lst, enterLeaveFunctionStmnt)
		f.Body.List = lst
		return
	} else {
		enterFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Entering \"%v\"\n"`, f.FuncIdentifier))
		lst = append([]dst.Stmt{enterFunctionStmnt}, lst...)
		f.Body.List = lst
	}

	if !f.DoesHaveReturnStatement {
		leaveFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Leaving \"%v\"\n"`, f.FuncIdentifier))
		f.Body.List = append(f.Body.List, leaveFunctionStmnt)
	}
}

func indexOfExpr(slice []dst.Expr, node dst.Node) int {
	for idx, elem := range slice {
		if elem == node {
			return idx
		}
	}
	return -1
}

func getAnonymousFunctionName(c *dstutil.Cursor) string {
	node := c.Node()
	switch parent := c.Parent().(type) {
	case *dst.AssignStmt:
		idx := indexOfExpr(parent.Rhs, node)
		lhs := parent.Lhs[idx]
		varNode, _ := lhs.(*dst.Ident)
		return varNode.Name
	case *dst.ValueSpec:
		idx := indexOfExpr(parent.Values, node)
		return parent.Names[idx].Name
	default:
		return "anonymous-function " + c.Name()
	}
}

func AddPrintDebugging(options *options.Options, codeBytes *bytes.Buffer) (*bytes.Buffer, error) {
	fset := token.NewFileSet()
	dec := decorator.NewDecoratorWithImports(fset, "src.go", goast.New())
	file, err := dec.Parse(codeBytes)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	infoStack := make(funcInfoStack, 0)

	// todo print line number?
	// todo go fmt
	// todo optional message formatting (before and after function with pos and name as template vars)
	// todo optional list of excluded or included function names

	preApply := func(c *dstutil.Cursor) bool {
		switch x := c.Node().(type) {
		case *dst.FuncDecl:
			infoStack = infoStack.Push(
				newFuncInfo(x.Name.Name, x.Body, x.Type),
			)
		case *dst.FuncLit:
			funcName := getAnonymousFunctionName(c)
			infoStack = infoStack.Push(newFuncInfo(funcName, x.Body, x.Type))
		}
		return true
	}

	postApply := func(c *dstutil.Cursor) bool {

		switch c.Node().(type) {
		case *dst.ReturnStmt:
			leavingFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Leaving \"%v\"\n"`, infoStack.Peek().FuncIdentifier))
			c.InsertBefore(leavingFunctionStmnt)
		case *dst.FuncDecl:
			infoStack.Peek().modifyFuncBody()
			infoStack, _ = infoStack.Pop()
		case *dst.FuncLit:
			infoStack.Peek().modifyFuncBody()
			infoStack, _ = infoStack.Pop()
		}
		return true
	}

	dstutil.Apply(file, preApply, postApply)
	codeBuffer := &bytes.Buffer{}
	err = populateBuffer(file, codeBuffer)
	if err != nil {
		return &bytes.Buffer{}, err
	} else {
		return codeBuffer, nil
	}
}

func RemovePrintDebugging(options *options.Options, codeBytes *bytes.Buffer) (*bytes.Buffer, error) {
	fset := token.NewFileSet()
	file, err := decorator.ParseFile(fset, "src.go", codeBytes, parser.ParseComments)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	preApply := func(c *dstutil.Cursor) bool {
		if c.Node() != nil {
			end := c.Node().Decorations().End
			if len(end) == 1 && end[0] == DecorationMessage {
				c.Delete()
			}
		}
		return true
	}
	postApply := func(c *dstutil.Cursor) bool { return true }
	dstutil.Apply(file, preApply, postApply)
	codeBuffer := &bytes.Buffer{}
	err = populateBuffer(file, codeBuffer)
	if err != nil {
		return &bytes.Buffer{}, err
	} else {
		return codeBuffer, nil
	}
}

func populateBuffer(file *dst.File, buffer *bytes.Buffer) error {
	res := decorator.NewRestorerWithImports("main", guess.New())
	return res.Fprint(buffer, file)
}
