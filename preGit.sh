#!/bin/bash

gofmt -s -w hello/hello.go
gofmt -s -w hello/hello_test.go
go test -v -coverprofile=c0.out -covermode=atomic github.com/mchirico/go_test/hello
go vet -v github.com/mchirico/go_test/hello
