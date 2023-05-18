export tag=v1.0

root:
    export ROOT=github.com/feiwang137/golang

build:
	echo "${ROOT}"
	echo "building"

push:
	echo "push feiwang137/golang"

test:
	echo "testing"

test2: test
	echo "test2"