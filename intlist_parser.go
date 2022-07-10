package tdd_args

import (
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
		argsInt, _ := strconv.ParseInt(args, 10, 64)
		intList = append(intList, int(argsInt))
	}
	parser.IntList = intList

	return nil
}

func (il *IntListParser) containsInvalidArgs(argsList []string, idx int) bool {
	return len(argsList) > idx+1 && !strings.Contains(argsList[idx+1], "-")
}
