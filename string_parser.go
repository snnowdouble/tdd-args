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
	if len(argsList) > idx+2 && !strings.Contains(argsList[idx+2], "-") {
		return fmt.Errorf("too many args error")
	}

	return nil
}
