package tdd_args

type ArgsParser struct {
	Logging   bool   `tag:"l"`
	Port      int    `tag:"p"`
	Directory string `tag:"d"`
}

type ArgsSchema struct {
	Name     string
	DataType string
	Tag      string
}

func Parse(argsParser *ArgsParser, argsList ...string) {
	argsParser.Logging = true

}
