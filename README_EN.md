# GoSimhash: Simhash Lib Powered by Golang [中文](README.md)

[![Build Status](https://travis-ci.org/yanyiwu/gosimhash.png?branch=master)](https://travis-ci.org/yanyiwu/gosimhash) 
[![Author](https://img.shields.io/badge/author-@yanyiwu-blue.svg?style=flat)](http://yanyiwu.com/) 
[![GoDoc](https://godoc.org/github.com/yanyiwu/gosimhash?status.svg)](https://godoc.org/github.com/yanyiwu/gosimhash)
[![Coverage Status](https://coveralls.io/repos/yanyiwu/gosimhash/badge.svg?branch=master&service=github)](https://coveralls.io/github/yanyiwu/gosimhash?branch=master)
[![License](https://img.shields.io/badge/license-MIT-yellow.svg?style=flat)](http://yanyiwu.mit-license.org)

## Usage

```
go get github.com/yanyiwu/gosimhash
```

See example in [example/demo.go](example/demo.go)

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


## Contact

+ Email: `i@yanyiwu.com`
+ QQ: 64162451
+ WeChat: ![image](http://7viirv.com1.z0.glb.clouddn.com/5a7d1b5c0d_yanyiwu_personal_qrcodes.jpg)

[simhash]:http://github.com/yanyiwu/simhash
