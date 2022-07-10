package tdd_args

type SpecificParser interface {
	parser(parser *ArgsParser, value string) error
}
