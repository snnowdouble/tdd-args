package tdd_args

import (
	"strconv"
)

type IntParser struct {
}

func (i *IntParser) parser(parser *ArgsParser, value string) {
	argsInt, _ := strconv.ParseInt(value, 10, 64)
	parser.Port = int(argsInt)

}
