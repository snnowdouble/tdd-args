package tdd_args

type BoolParser struct {
}

func (b *BoolParser) parser(parser *ArgsParser, value string) error {
	parser.Logging = true

	return nil
}
