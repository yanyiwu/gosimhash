# GoSimhash: Simhash Lib Powered by Golang [English](README_EN.md)

[![Build Status](https://travis-ci.org/yanyiwu/gosimhash.png?branch=master)](https://travis-ci.org/yanyiwu/gosimhash) 
[![Author](https://img.shields.io/badge/author-@yanyiwu-blue.svg?style=flat)](http://yanyiwu.com/) 
[![GoDoc](https://godoc.org/github.com/yanyiwu/gosimhash?status.svg)](https://godoc.org/github.com/yanyiwu/gosimhash)
[![Coverage Status](https://coveralls.io/repos/yanyiwu/gosimhash/badge.svg?branch=master&service=github)](https://coveralls.io/github/yanyiwu/gosimhash?branch=master)
[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](http://yanyiwu.mit-license.org)

[![logo](http://7viirv.com1.z0.glb.clouddn.com/GoSimhashLogo-v1.png)](https://github.com/yanyiwu/gosimhash)

## 简介

GoSimhash 是 中文 [simhash] 去重算法库，Golang版本。

## 用法

```
go get github.com/yanyiwu/gosimhash
```

示例代码请见 [example/demo.go](example/demo.go)

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

+ Email: `i@yanyiwu.com`
+ QQ: 64162451
+ WeChat: ![image](http://7viirv.com1.z0.glb.clouddn.com/5a7d1b5c0d_yanyiwu_personal_qrcodes.jpg)

[simhash]:http://github.com/yanyiwu/simhash


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/yanyiwu/gosimhash/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

