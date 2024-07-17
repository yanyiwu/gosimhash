# GoSimhash: Simhash Lib Powered by Golang

[![Test](https://github.com/yanyiwu/gosimhash/actions/workflows/go.yml/badge.svg)](https://github.com/yanyiwu/gosimhash/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/yanyiwu/gosimhash?status.svg)](https://godoc.org/github.com/yanyiwu/gosimhash)
[![Coverage Status](https://coveralls.io/repos/yanyiwu/gosimhash/badge.svg?branch=master&service=github)](https://coveralls.io/github/yanyiwu/gosimhash?branch=master)
[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](http://yanyiwu.mit-license.org)

## Introduction

GoSimhash is a Go version of [simhash] algorithm.

## Usage

```
go get github.com/yanyiwu/gosimhash
```

The sample code can be found at [example/demo.go](example/demo.go)

```
cd example
go build

./example -help

# Usage of ./example:
#     -sentence="我来到北京清华大学":
#     -top_n=5:

./example

# 我来到北京清华大学 simhash: feb6372a8750eb1d

./example -sentence="南京市长江大桥" -top_n=5

# 南京市长江大桥 simhash: b2c6a622481d8eb2
```

[simhash]:http://github.com/yanyiwu/simhash

