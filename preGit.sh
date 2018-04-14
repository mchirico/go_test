#!/bin/bash

#  git push --delete origin tagName
#  git tag -d tagName
#
# git tag -a v1.0 -m 'simple main test with flags'
# git checkout tags/v1.0



gofmt -s -w hello/hello.go
gofmt -s -w hello/hello_test.go
go test -v -coverprofile=c0.out -covermode=atomic github.com/mchirico/go_test/hello
go vet -v github.com/mchirico/go_test/hello
