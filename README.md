[![Build Status](https://travis-ci.org/cloudfoundry-incubator/api.png)](https://travis-ci.org/cloudfoundry-incubator/api)
api
===

Experimental CF API server written in GO


===
#Developer Setup

First, install Go from http://golang.org/doc/install (not from homebrew)

```
$ go get github.com/kr/godep
$ mkdir -p $GOPATH/src/github.com/cloudfoundry-incubator
$ cd $GOPATH/src/github.com/cloudfoundry-incubator
$ git clone git@github.com:cloudfoundry-incubator/api.git
$ cd api
$ bin/test_setup
$ bin/test
```
