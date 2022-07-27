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
	"path"
	"regexp"
	"strconv"
	"strings"
)

type funcInfo struct {
	FuncIdentifier          string
	Body                    *dst.BlockStmt
	MustHaveReturnStatement bool
	DoesHaveReturnStatement bool
}

var DecorationMessage = "// automatically added by printf-debugger. Do not change this comment. It is an identifier."
var functionPrefix = "printfdebug_Printf_"

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

func getForceImportVariables() []*dst.GenDecl {
	// will add something like "var _ = runtime.Caller"  to the bottom of the file. This is to force dst to import functions we use in the printf function
	pkgToSampleFunc := map[string]string{
		"runtime":       "Caller",
		"path/filepath": "Clean",
		"fmt":           "Println",
		"strings":       "Split",
		"os":            "PathSeparator",
	}

	valueSpecs := []*dst.GenDecl{}
	for pkg, sampleFunc := range pkgToSampleFunc {
		rawValuSpec :=
			&dst.GenDecl{
				Tok: token.VAR,
				Specs: []dst.Spec{
					&dst.ValueSpec{
						Names: []*dst.Ident{
							{
								Name: "_",
								Obj: &dst.Object{
									Kind: dst.Var,
									Name: "_",
								},
								Path: "",
							},
						},
						Values: []dst.Expr{
							&dst.Ident{
								Name: sampleFunc,
								Path: pkg,
							},
						},
					},
				},
				Rparen: false,
			}

		rawValuSpec.Decs.End.Append(DecorationMessage)
		valueSpecs = append(valueSpecs, rawValuSpec)
	}

	return valueSpecs
}

func getFileSpecificFunctionName(filename string) string {

	// tries to produce a valid function name. This could result in duplicate function names within the same package down the road
	noExt := strings.Replace(filename, ".go", "", 1)
	validVarNameRegEx := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	funcName := validVarNameRegEx.ReplaceAll([]byte(noExt), []byte{})

	return fmt.Sprintf("%v%v", functionPrefix, string(funcName))
}
func getRuntimeFuncAsString(functionName string) string {
	return fmt.Sprintf(`func %v(message string, pathDepthFromEnd int) {
	minInt := func(first int, second int) (min int) {
		if first < second {
			return first
		} else {
			return second
		}
	}

	_, file, line, ok := runtime.Caller(1)
	if ok {
		fileParts := strings.Split(file, string(os.PathSeparator))
		pathFromEndSafe := minInt(len(fileParts), pathDepthFromEnd)
		limited := filepath.Join(fileParts[len(fileParts)-pathFromEndSafe:]...)
		limitedCleaned := "??"
		if limited != "" {
			limitedCleaned = limited
		}
		fmt.Printf("%%v:%%v %%v\n", limitedCleaned, line, message)
	} else {
		fmt.Printf("unkown_file:? %%v\n", message)
	}
}
`, functionName)
}

func getPrintStatement(message string, options *options.Options) (result *dst.ExprStmt) {
	if options.NoRuntime {
		//outputs something like `fmt.Printf("Entering \"Func\"\n")`
		result = &dst.ExprStmt{
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
	} else {
		// outputs something like `printfdebug_Printf_Func("Entering \"Func\"\n",1)`
		name := getFileSpecificFunctionName(path.Base(options.FilePath))
		result = &dst.ExprStmt{
			X: &dst.CallExpr{
				Fun: &dst.Ident{Name: name, Path: ""},
				Args: []dst.Expr{
					&dst.BasicLit{
						Kind:  token.STRING,
						Value: message,
					},
					&dst.BasicLit{
						Kind:  token.INT,
						Value: strconv.Itoa(options.PathDepth),
					},
				},
			},
		}
	}

	result.Decs.End.Append(DecorationMessage)
	return result
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

func (f *funcInfo) modifyFuncBody(options *options.Options) {

	enterFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Entering \"%v\"\n"`, f.FuncIdentifier), options)
	f.Body.List = append([]dst.Stmt{enterFunctionStmnt}, f.Body.List...)

	if !f.DoesHaveReturnStatement {
		// handles adding entering print statements and leaving print statements for implicit returns
		leaveFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Leaving \"%v\"\n"`, f.FuncIdentifier), options)
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

func getAnonymousFunctionName(c *dstutil.Cursor, parent dst.Node, grandParent dst.Node) string {
	node := c.Node()
	resolvedName := ""
	switch parentTyped := parent.(type) {
	case *dst.AssignStmt:
		// we want to extract varName from something like "varName:=func(){}"
		idx := indexOfExpr(parentTyped.Rhs, node)
		lhs := parentTyped.Lhs[idx]
		varNode, _ := lhs.(*dst.Ident)
		resolvedName = varNode.Name
	case *dst.ValueSpec:
		// we want to extract varName from something like "var varName =func(){}"
		idx := indexOfExpr(parentTyped.Values, node)
		resolvedName = parentTyped.Names[idx].Name

	}
	if resolvedName != "" {
		return resolvedName
	} else {
		// we want to extract varName from something like "var varName =func(){}()"
		switch grandParentTyped := grandParent.(type) {
		case *dst.ValueSpec:
			idx := indexOfExpr(grandParentTyped.Values, parent)
			resolvedName = grandParentTyped.Names[idx].Name
		}
	}
	if resolvedName == "" {
		resolvedName = "anonymous-function " + c.Name()
	}

	return resolvedName
}

func AddPrintDebugging(options *options.Options, codeBytes *bytes.Buffer) (*bytes.Buffer, error) {
	fset := token.NewFileSet()
	dec := decorator.NewDecoratorWithImports(fset, "src.go", goast.New())
	file, err := dec.Parse(codeBytes)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	infoStack := make(funcInfoStack, 0)

	nodeToParent := make(map[dst.Node]dst.Node)

	preApply := func(c *dstutil.Cursor) bool {
		nodeToParent[c.Node()] = c.Parent()
		switch x := c.Node().(type) {
		case *dst.FuncDecl:
			infoStack = infoStack.Push(
				newFuncInfo(x.Name.Name, x.Body, x.Type),
			)
		case *dst.FuncLit:
			parent := nodeToParent[x]
			grandParent := nodeToParent[parent]
			funcName := getAnonymousFunctionName(c, parent, grandParent)
			infoStack = infoStack.Push(newFuncInfo(funcName, x.Body, x.Type))
		case *dst.File:
			if !options.NoRuntime {
				for _, nd := range getForceImportVariables() {
					// for some reason the type checker threw a "Cannot use 'getForceImportVariables()' (type []*dst.GenDecl) as the type []Decl"
					// when writing "x.Decls=append(x.Decls,getForceImportVariables()...)" but allowed us to append it in a loop...
					x.Decls = append(x.Decls, nd)
				}
			}
		}

		return true
	}

	postApply := func(c *dstutil.Cursor) bool {

		switch x := c.Node().(type) {
		case *dst.ReturnStmt:
			leavingFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Leaving \"%v\"\n"`, infoStack.Peek().FuncIdentifier), options)
			c.InsertBefore(leavingFunctionStmnt)
		case *dst.ExprStmt:
			callExpr, ok := x.X.(*dst.CallExpr)
			if !ok {
				break
			}
			name := ""
			_path := ""

			// package qualified function call
			selectorExpr, ok := callExpr.Fun.(*dst.SelectorExpr)
			if ok {
				ident, ok := selectorExpr.X.(*dst.Ident)
				if ok {
					name = selectorExpr.Sel.Name
					_path = ident.Name

				}
			}

			// function call
			ident, ok := callExpr.Fun.(*dst.Ident)
			if ok {
				_path = ident.Path
				name = ident.Name
			}
			if _path == "" && name == "" {
				break
			}

			earlyProgramExit := false
			if _path == "" && name == "panic" {
				earlyProgramExit = true
			}
			if _path == "os" && name == "Exit" {
				earlyProgramExit = true
			}
			if _path == "log" && name == "Fatalf" {
				earlyProgramExit = true
			}
			if _path == "log" && name == "Fatal" {
				earlyProgramExit = true
			}
			if _path == "log" && name == "Fatalln" {
				earlyProgramExit = true
			}
			if earlyProgramExit {
				leavingFunctionStmnt := getPrintStatement(fmt.Sprintf(`"Leaving \"%v\"\n"`, infoStack.Peek().FuncIdentifier), options)
				c.InsertBefore(leavingFunctionStmnt)
			}
		case *dst.FuncDecl:
			infoStack.Peek().modifyFuncBody(options)
			infoStack, _ = infoStack.Pop()
		case *dst.FuncLit:
			infoStack.Peek().modifyFuncBody(options)
			infoStack, _ = infoStack.Pop()

		}
		return true
	}

	dstutil.Apply(file, preApply, postApply)
	codeBuffer := &bytes.Buffer{}
	err = populateBuffer(file, codeBuffer)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	if !options.NoRuntime {
		funcName := getFileSpecificFunctionName(path.Base(options.FilePath))
		runtimeFunc := getRuntimeFuncAsString(funcName)
		codeBuffer.WriteString(runtimeFunc)
	}

	return codeBuffer, nil

}

func RemovePrintDebugging(options *options.Options, codeBytes *bytes.Buffer) (*bytes.Buffer, error) {
	fset := token.NewFileSet()
	dec := decorator.NewDecoratorWithImports(fset, "src.go", goast.New())
	file, err := dec.ParseFile("src.go", codeBytes, parser.ParseComments)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	preApply := func(c *dstutil.Cursor) bool {
		if c.Node() != nil {
			end := c.Node().Decorations().End
			if len(end) == 1 && end[0] == DecorationMessage {
				c.Delete()
			}
			funcDecl, ok := c.Node().(*dst.FuncDecl)
			if ok && strings.Contains(funcDecl.Name.Name, functionPrefix) {
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
