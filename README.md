# JSON Tools by Language Recognition

Inspire by [stedolan/jq](https://github.com/stedolan/jq), Created with [antlr4](https://github.com/antlr/antlr4).

### try it

> go get -u github.com/xiaoler/jtlr-go/cmd/jtlr

### feature

basic usage:
> jtlr {"a": 1}

interactive mode:
> jtlr -a

readline from stdin:
> echo stdout.sh | jtlr -s

### todo:

- 上色
- ~~unicode转义~~
- ~~cli flag~~
- ~~stdin读取、交互模式~~
- interactive模式下特殊支付处理
- print改io.Writer
- 子树解析
- Marshal && Unmarshal
