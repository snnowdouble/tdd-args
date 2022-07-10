package tdd_args

import (
	"fmt"
	"regexp"
	"strconv"
)

type IntListParser struct {
}

func (il *IntListParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	argsList = argsList[idx+1:]
	intList := make([]int, 0)
	var flagPattern = regexp.MustCompile(`^-([a-z]|[A-Z])+$`)
	for _, args := range argsList {
		if flagPattern.MatchString(args) {
			break
		}
		argsInt, err := strconv.ParseInt(args, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid args error")
		}
		intList = append(intList, int(argsInt))
	}
	parser.IntList = intList

	return nil
}
