package tdd_args

import (
	"reflect"
)

type ArgsParser struct {
	Logging    bool     `tag:"l"`
	Port       int      `tag:"p"`
	Directory  string   `tag:"d"`
	IntList    []int    `tag:"t"`
	StringList []string `tag:"g"`
}

type ArgsSchema struct {
	Name     string
	DataType string
	Tag      string
}

type ArgsParserFactoryReq struct {
	DataType   string
	argsParser *ArgsParser
	valueList  []string
	idx        int
}

var initParserMap map[string]SpecificParser

func registerParseMap() map[string]SpecificParser {
	initParserMap := make(map[string]SpecificParser)
	initParserMap["bool"] = new(BoolParser)
	initParserMap["int"] = new(IntParser)
	initParserMap["string"] = new(StringParser)

	return initParserMap
}

func Parse(argsParser *ArgsParser, argsList ...string) error {
	argsSchemaMap := argsSchemaMapBuilder()
	for idx, args := range argsList {
		if argsSchema, ok := argsSchemaMap[args]; ok {
			err := parserFactory(&ArgsParserFactoryReq{
				DataType:   argsSchema.DataType,
				argsParser: argsParser,
				valueList:  argsList,
				idx:        idx,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil

}

func argsSchemaMapBuilder() map[string]*ArgsSchema {
	argsType := reflect.TypeOf(ArgsParser{})
	argsSchemaMap := make(map[string]*ArgsSchema)
	for i := 0; i < argsType.NumField(); i++ {
		argsSchemaMap[getArgsSchemaKey(argsType, i)] = &ArgsSchema{
			Name:     argsType.Field(i).Name,
			DataType: argsType.Field(i).Type.String(),
			Tag:      argsType.Field(i).Tag.Get("tag"),
		}
	}
	return argsSchemaMap
}

func getArgsSchemaKey(argsType reflect.Type, i int) string {
	return "-" + argsType.Field(i).Tag.Get("tag")
}

func parserFactory(req *ArgsParserFactoryReq) error {
	initParserMap := registerParseMap()
	err := initParserMap[req.DataType].parser(req.argsParser, req.valueList, req.idx)
	if err != nil {
		return err
	}

	return nil
}
