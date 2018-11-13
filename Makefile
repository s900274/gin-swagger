GOPATH:=$(CURDIR)
export GOPATH

all: clean dep bld

bld: gin-test

gin-test:
	go build -o bin/gin-test gintest

clean:
	@rm -f bin/gin-test
	@rm -f ./db/*.db
	@rm -rf ./pkg
	@rm -rf status
	@rm -f  log/*log*
	@rm -rf ./output

cleanlog:
	@rm -f log/*log*

dep:
	go get github.com/gin-gonic/gin
	go get github.com/swaggo/gin-swagger
	go get github.com/swaggo/gin-swagger/swaggerFiles
	go get github.com/alecthomas/template
	go get github.com/mattn/go-sqlite3
	go get github.com/BurntSushi/toml
	go get github.com/shengkehua/xlog4go