package tdd_args

type StringParser struct {
}

func (s *StringParser) parser(parser *ArgsParser, value string) error {
	parser.Directory = value

	return nil
}
