package tdd_args

type SpecificParser interface {
	parser(parser *ArgsParser, argsList []string, idx int) error
}
