package parser

import (
	"github.com/magiconair/properties"
)

//PropertiesFileParser ...
type PropertiesFileParser struct {
	parserType string
}

//NewPropertiesFileParser ...
func NewPropertiesFileParser() *PropertiesFileParser {
	return &PropertiesFileParser{
		parserType: "properties",
	}
}

//GetParserType ...
func (parser *PropertiesFileParser) GetParserType() string {
	return parser.parserType
}

//Parse ...
func (parser *PropertiesFileParser) Parse(file string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	prop, err := properties.LoadFile(file, properties.UTF8)
	if err != nil {
		return result, err
	}
	m := prop.Map()
	for key, val := range m {
		result[key] = val
	}
	return result, nil
}
