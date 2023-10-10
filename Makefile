.PHONY: build
build: 
	-rm -rf bin && mkdir bin && go build -o bin/module main.go

package:
	tar -czf module.tar.gz bin/module
