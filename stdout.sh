#!/bin/bash
n=1
while [ $n -le 10 ]
do
    echo '{"a": [1, 2], "b": {"a":1, "b":2}}'
    let n++
    sleep 1s
done
