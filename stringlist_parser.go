package tdd_args

import (
	"regexp"
	"strings"
)

type StringListParser struct {
	IntParser
}

func (sl *StringListParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	argsList = argsList[idx+1:]
	stringList := make([]string, 0)
	var flagPattern = regexp.MustCompile(`^-([a-z]|[A-Z])+$`)
	for _, args := range argsList {
		if flagPattern.MatchString(args) {
			break
		}
		stringList = append(stringList, args)
	}
	parser.StringList = stringList

	return nil
}

func (sl *StringListParser) containsInvalidArgs(argsList []string, idx int) bool {
	return len(argsList) > idx+1 && !strings.Contains(argsList[idx+1], "-")
}
