package tdd_args

import (
	"fmt"
	"strings"
)

type BoolParser struct {
}

func (b *BoolParser) parser(parser *ArgsParser, value string) error {
	parser.Logging = true
	if !strings.Contains(value, "-") {
		return fmt.Errorf("too many args error")
	}

	return nil
}
