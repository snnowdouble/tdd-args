package tdd_args

import (
	"regexp"
)

type StringListParser struct {
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
