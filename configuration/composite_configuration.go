package configuration

import "fmt"

var (
    //ErrorIndexOutOfBound ...
    ErrorIndexOutOfBound = fmt.Errorf("index out of bound")
    //ErrorConfigExists ...
    ErrorConfigExists = fmt.Errorf("config exists")
)

//CompositeConfiguration ...
type CompositeConfiguration struct {
    *MapConfiguration
    namedConfigurations          map[string]Configuration
    configList                   []Configuration
}

//NewCompositeConfiguration ...
func NewCompositeConfiguration() *CompositeConfiguration {
    cfg := &CompositeConfiguration{
        MapConfiguration:       NewMapConfiguration(nil),
        namedConfigurations:    make(map[string]Configuration),
        configList:             make([]Configuration, 0),
    }
    cfg.MapConfiguration.SetConfigurationAccessor(cfg)
    return cfg
}

//AddConfiguration ...
func (c *CompositeConfiguration)AddConfiguration(config Configuration) {
    c.AddNamedConfiguration(config, "")
}

//AddNamedConfiguration ...
func (c *CompositeConfiguration)AddNamedConfiguration(config Configuration, name string) {
    c.addConfigurationAtIndex(config, name, len(c.configList))
}

func (c *CompositeConfiguration)getIndexOfConfiguration(config Configuration) int {
    for i, cfg := range c.configList{
        if cfg == config {
            return i
        }
    }
    return len(c.configList)
}

//GetConfigurations ...
func (c *CompositeConfiguration)GetConfigurations() []Configuration {
    num := len(c.configList)
    configList := make([]Configuration, num)
    for i, config := range c.configList {
        configList[i] = config
    }
    return configList
}

//AddConfigurationAtIndex add a configuration with a name at a particular index.
func (c *CompositeConfiguration)addConfigurationAtIndex(config Configuration, name string, index int) error{
    err := c.checkIndex(index)
    if err != nil {
        return err
    }
    for _, conf := range c.configList {
        if conf == config {
            return ErrorConfigExists
        }
    }
    length := len(c.configList)
    c.configList = append(c.configList, nil)
    if index < length {
        c.configList = append(c.configList[:index + 1], c.configList[index:length]...)
        if len(name) != 0 {
            c.namedConfigurations[name] = config
        }
    }
    for _, l := range c.MapConfiguration.GetConfigurationListeners() {
        config.AddConfigurationListener(l)
    }
    c.configList[index] = config
    return nil
}

//AddConfigurationAtFront ...
func (c *CompositeConfiguration)AddConfigurationAtFront(config Configuration, name string) error {
    return c.addConfigurationAtIndex(config, name, 0)
}

func (c *CompositeConfiguration)checkIndex(index int) error {
    if index < 0 || index > len(c.configList) {
        return ErrorIndexOutOfBound
    }
    return nil
}

//GetNumberOfConfigurations ...
func (c *CompositeConfiguration)GetNumberOfConfigurations()int {
    return len(c.configList)
}

//GetConfigurationByName ...
func (c *CompositeConfiguration)GetConfigurationByName(name string) Configuration {
    return c.namedConfigurations[name]
}

//RemoveConfigurationByName ...
func (c *CompositeConfiguration)RemoveConfigurationByName(name string)Configuration {
    config := c.GetConfigurationByName(name)
    if config != nil {
        index := c.getIndexOfConfiguration(config)
        if index != len(c.configList) {
            c.configList = append(c.configList[:index], c.configList[index + 1:]...)
        }
    }
    return config
}

//RemoveConfiguration ...
func (c *CompositeConfiguration)RemoveConfiguration(config Configuration) bool {
    name := c.getNameForConfiguration(config)
    if len(name) != 0 {
        delete(c.namedConfigurations, name)
    }
    index := c.getIndexOfConfiguration(config)
    c.configList = append(c.configList[:index], c.configList[index + 1:]...)
    return true
}

func (c *CompositeConfiguration)getNameForConfiguration(config Configuration) string {
    for name, conf := range c.namedConfigurations {
        if conf == config {
            return name
        }
    }
    return ""
}

//Clear ...
func (c *CompositeConfiguration)Clear() {
    for _, config := range c.configList {
        config.Clear()
    }
    c.MapConfiguration.Clear()
    c.configList = make([]Configuration, 0)
}

//GetProperty ...
func (c *CompositeConfiguration)GetProperty(key string) interface{} {
    for _, config := range c.configList {
        if config.GetProperty(key) != nil {
            return config.GetProperty(key)
        }
    }
    return c.MapConfiguration.GetProperty(key)
}

//GetKeys ...
func (c *CompositeConfiguration)GetKeys() []string {
    keys := make([]string, 0)
    for _, config := range c.configList {
        keys = append(keys, config.GetKeys()...)
    }
    keys = append(keys, c.MapConfiguration.GetKeys()...)
    return keys
}

//ContainsKeys ...
func (c *CompositeConfiguration)ContainsKeys(key string) bool {
    for _, config := range c.configList {
        if config.ContainsKey(key) {
            return true
        }
    }
    return c.MapConfiguration.ContainsKey(key)
}

//IsEmpty ...
func (c *CompositeConfiguration)IsEmpty() bool {
    for _, config := range c.configList {
        if config.IsEmpty() == false {
            return false
        }
    }
    return c.MapConfiguration.IsEmpty()
}

//AddConfigurationListener ...
func (c *CompositeConfiguration)AddConfigurationListener(l ConfigurationListener) {
    for _, config := range c.configList {
        config.AddConfigurationListener(l)
    }
    c.MapConfiguration.AddConfigurationListener(l)
}

//RemoveConfigurationListener ...
func (c *CompositeConfiguration)RemoveConfigurationListener(l ConfigurationListener) {
    for _, config := range c.configList {
        config.RemoveConfigurationListener(l)
    }
    c.MapConfiguration.RemoveConfigurationListener(l)
}