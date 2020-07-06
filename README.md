# JSON Tools by Language Recognition

Inspire by [stedolan/jq](https://github.com/stedolan/jq), Created with [antlr4](https://github.com/antlr/antlr4).

### Try it

> go get -u github.com/xiaoler/jtlr-go/cmd/jtlr

### Feature

basic usage:
> jtlr '{"a": 1}'

interactive mode:
> jtlr -a

read line from stdin:
> bash examples/stdout.sh | jtlr -s

read line from file:
> jtlr -f examples/json.txt

set indent
> jtlr -t "  " '{"a": 1}'

### todo:

- 上色
- unicode转义 ✔︎
- cli flag ✔︎
- stdin读取、交互模式 ✔︎
- interactive模式下特殊字符处理 ✔︎
- 缩进中的 \t 处理
- print改io.Writer
- 子树解析
- Marshal && Unmarshal
