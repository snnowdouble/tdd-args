package tdd_args

type StringParser struct {
}

func (s *StringParser) parser(parser *ArgsParser, value string) {
	parser.Directory = value
}
