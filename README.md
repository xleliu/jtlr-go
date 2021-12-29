# JSON Tools by Language Recognition

Inspire by [stedolan/jq](https://github.com/stedolan/jq), Created with [antlr4](https://github.com/antlr/antlr4).

### Try it

> go install github.com/xiaoler/jtlr-go/cmd/jtlr

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
> jtlr -t "\t" '{"a": 1}'
