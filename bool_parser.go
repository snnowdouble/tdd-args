package tdd_args

import (
	"fmt"
	"strings"
)

type BoolParser struct {
}

func (b *BoolParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	parser.Logging = true
	if b.containsInvalidArgs(argsList, idx) {
		return fmt.Errorf("too many args error")
	}

	return nil
}

func (b *BoolParser) containsInvalidArgs(argsList []string, idx int) bool {
	return len(argsList) > idx+1 && !strings.Contains(argsList[idx+1], "-")
}
