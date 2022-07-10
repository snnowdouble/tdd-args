package tdd_args

import (
	"fmt"
	"strings"
)

type StringParser struct {
}

func (s *StringParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	if len(argsList) > idx+1 {
		parser.Directory = argsList[idx+1]
	}
	if s.containsInvalidArgs(argsList, idx) {
		return fmt.Errorf("too many args error")
	}

	return nil
}

func (s *StringParser) containsInvalidArgs(argsList []string, idx int) bool {
	return len(argsList) > idx+2 && !strings.Contains(argsList[idx+2], "-")
}
