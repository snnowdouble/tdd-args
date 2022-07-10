package tdd_args

import (
	"strings"
)

type StringListParser struct {
	IntParser
}

func (sl *StringListParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	argsList = argsList[1:]
	stringList := make([]string, 0)
	for _, args := range argsList {
		stringList = append(stringList, args)
	}
	parser.StringList = stringList

	return nil
}

func (sl *StringListParser) containsInvalidArgs(argsList []string, idx int) bool {
	return len(argsList) > idx+1 && !strings.Contains(argsList[idx+1], "-")
}
