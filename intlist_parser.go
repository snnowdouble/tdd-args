package tdd_args

import (
	"fmt"
	"strconv"
	"strings"
)

type IntListParser struct {
	IntParser
}

func (il *IntListParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	argsList = argsList[1:]
	intList := make([]int, 0)
	for _, args := range argsList {
		argsInt, err := strconv.ParseInt(args, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid args error")
		}
		intList = append(intList, int(argsInt))
	}
	parser.IntList = intList

	return nil
}

func (il *IntListParser) containsInvalidArgs(argsList []string, idx int) bool {
	return len(argsList) > idx+1 && !strings.Contains(argsList[idx+1], "-")
}
