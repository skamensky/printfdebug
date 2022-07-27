rm -rf dist
mkdir dist

function do_build(){
  echo "building for ${GOOS}/${GOARCH}"
  go build -o "./dist/printfdebug_${GOOS}_${GOARCH}$1"
}

export GOOS=darwin
export GOARCH=amd64
do_build


export GOOS=darwin
export GOARCH=arm64
do_build

export GOOS=linux
export GOARCH=386
do_build

export GOOS=linux
export GOARCH=amd64
do_build

export GOOS=linux
export GOARCH=arm
do_build

export GOOS=linux
export GOARCH=arm64
do_build

export GOOS=windows
export GOARCH=386
do_build ".exe"

export GOOS=windows
export GOARCH=amd64
do_build ".exe"