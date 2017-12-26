package parser

import (
    "fmt"
    "path"
)

var (
    parserRegister *fileParserRegister
)

func init() {
    parserRegister = newFileParserRegister()
}

//ConfigFileParser ...
type ConfigFileParser interface {
    //GetParserType ...
    GetParserType() string

    //Parse ...
    Parse(file string) (map[string]interface{}, error)
}

type fileParserRegister struct {
    parsers     map[string]ConfigFileParser
}

func newFileParserRegister() *fileParserRegister {
    register := &fileParserRegister{
        parsers:    make(map[string]ConfigFileParser, 0),
    }
    register.AddParser(NewPropertiesFileParser())
    return register
}

//AddParser ...
func (register *fileParserRegister)AddParser(parser ConfigFileParser) {
    if parser != nil {
        register.parsers[parser.GetParserType()] = parser
    }
}

//AddParsers ...
func (register *fileParserRegister)AddParsers(parsers []ConfigFileParser) {
    for _, parser := range parsers {
        register.AddParser(parser)
    }
}

func (register *fileParserRegister)parse(file string) (map[string]interface{}, error) {
    suffix := path.Ext(file)
    for _, parser := range register.parsers {
        if parser.GetParserType() != suffix[1:] {
            continue
        }
        return parser.Parse(file)
    }
    return nil, fmt.Errorf("no parser support to parse file:%s", file)
}

//ParseConfigFile ...
func ParseConfigFile(file string) (map[string]interface{}, error) {
    return parserRegister.parse(file)
}

//RegisterParser ...
func RegisterParser(parser ConfigFileParser) {
    parserRegister.AddParser(parser)
}

//RegisterParsers ...
func RegisterParsers(parsers []ConfigFileParser) {
    parserRegister.AddParsers(parsers)
}