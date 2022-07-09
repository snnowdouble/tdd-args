package tdd_args

import (
	"reflect"
)

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
	argsType := reflect.TypeOf(ArgsParser{})
	argsSchemaMap := make(map[string]*ArgsSchema)
	for i := 0; i < argsType.NumField(); i++ {
		argsSchemaMap["-"+argsType.Field(i).Tag.Get("tag")] = &ArgsSchema{
			Name:     argsType.Field(i).Name,
			DataType: argsType.Field(i).Type.String(),
			Tag:      argsType.Field(i).Tag.Get("tag"),
		}
	}
	for _, args := range argsList {
		if argsSchema, ok := argsSchemaMap[args]; ok {
			if argsSchema.DataType == "bool" {
				argsParser.Logging = true
			}
		}
	}

}
