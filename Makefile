.PHONY: build
build: 
	mkdir bin && go build -o bin/module module/main.go	
