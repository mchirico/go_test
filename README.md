[![Build Status](https://travis-ci.org/mchirico/go_test.svg?branch=master)](https://travis-ci.org/mchirico/go_test)
[![Go Report Card](https://goreportcard.com/badge/github.com/mchirico/go_test)](https://goreportcard.com/report/github.com/mchirico/go_test)
[![codecov](https://codecov.io/gh/mchirico/go_test/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/go_test)
[![Maintainability](https://api.codeclimate.com/v1/badges/437d440d9cf5fbe592a6/maintainability)](https://codeclimate.com/github/mchirico/go_test/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/437d440d9cf5fbe592a6/test_coverage)](https://codeclimate.com/github/mchirico/go_test/test_coverage)

## Testing in Go

Testing main exit status err, based
on flag input.

## Creating Work Environment

Run the following command to create `go_test`, which will become your work environment.

```bash
#!/bin/bash
mkdir -p go_test
cd go_test
wget https://gist.githubusercontent.com/mchirico/177447481b3eeab4f38e0071f26b816f/raw/90b4beb073a803d8a6aee863fc9bcfa945811c17/setpath
chmod 700 setpath
./setpath
cd src/github.com/mchirico
git clone git@github.com:mchirico/go_test.git
```



```bash
git checkout tags/v1.1

```

```bash
# Install
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen

# If you want, you can build mockgen in the current directory

go build github.com/golang/mock/mockgen
```


```bash
mockgen net/http Get

```






```bash
mockgen -destination=./mocks/mock_stuff.go -package=mocks database/sql/driver Conn,Driver

```