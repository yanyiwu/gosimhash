# GoSimhash

[![Build Status](https://travis-ci.org/yanyiwu/gosimhash.png?branch=master)](https://travis-ci.org/yanyiwu/gosimhash) 
[![GoDoc](https://godoc.org/github.com/yanyiwu/gosimhash?status.svg)](https://godoc.org/github.com/yanyiwu/gosimhash)

GoSimhash 是 中文 [simhash] 库的 Go 语言版本库。

## 用法

```
go get github.com/yanyiwu/gosimhash
```

示例代码请见 example/demo.go

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

之所以需要先 cd 到 example 目录下，是因为 demo.go 里面有写死的字典相对路径。

## 客服

```
i@yanyiwu.com
```

[simhash]:http://github.com/yanyiwu/simhash
