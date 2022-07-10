package tdd_args

import (
	"fmt"
	"strconv"
	"strings"
)

type IntParser struct {
}

func (i *IntParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	if len(argsList) > idx+1 {
		argsInt, _ := strconv.ParseInt(argsList[idx+1], 10, 64)
		parser.Port = int(argsInt)
	}
	if len(argsList) > idx+2 && !strings.Contains(argsList[idx+2], "-") {
		return fmt.Errorf("too many args error")
	}
	return nil

}
