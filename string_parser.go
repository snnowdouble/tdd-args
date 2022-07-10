package tdd_args

type StringParser struct {
}

func (s *StringParser) parser(parser *ArgsParser, argsList []string, idx int) error {
	if len(argsList) > idx+1 {
		parser.Directory = argsList[idx+1]
	}

	return nil
}
