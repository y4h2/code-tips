#!/bin/sh

go get -u github.com/dvyukov/go-fuzz/go-fuzz \ 
github.com/dvyukov/go-fuzz/go-fuzz-build

go-fuzz-build ./api

go-fuzz -bin=./api-fuzz.zip -workdir=workdir/corpus