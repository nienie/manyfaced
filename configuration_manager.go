package manyfaced

import "github.com/nienie/manyfaced/configuration"

var (
    configManager *configurationManager
)

func init() {
    configManager = newConfigurationManager()
    propertyFactory = newDynamicPropertyFactory()
    propertyFactory.initWithConfigurationSource(configManager.getCompositeConfiguration())
}

type configurationManager struct {
    configContainer     *configuration.CompositeConfiguration
}

func newConfigurationManager() *configurationManager {
    return &configurationManager{
        configContainer:     configuration.NewCompositeConfiguration(),
    }
}

func (c *configurationManager)addNamedConfiguration(config configuration.Configuration, name string) {
    c.configContainer.AddNamedConfiguration(config, name)
}

func (c *configurationManager)removeNamedConfiguration(name string) configuration.Configuration {
    return c.configContainer.RemoveConfigurationByName(name)
}

func (c *configurationManager)getNamedConfiguration(name string)configuration.Configuration {
    return c.configContainer.GetConfigurationByName(name)
}

func (c *configurationManager)addConfiguration(config configuration.Configuration) {
    c.configContainer.AddConfiguration(config)
}

func (c *configurationManager)getConfigurations()[]configuration.Configuration {
    return c.configContainer.GetConfigurations()
}

func (c *configurationManager)removeConfiguration(config configuration.Configuration) {
    c.configContainer.RemoveConfiguration(config)
}

func (c *configurationManager)getCompositeConfiguration()configuration.Configuration {
    return c.configContainer
}

//AddNamedConfiguration ...
func AddNamedConfiguration(config configuration.Configuration, name string) {
    configManager.addNamedConfiguration(config, name)
}

//RemoveNamedConfiguration ...
func RemoveNamedConfiguration(name string) configuration.Configuration {
    return configManager.removeNamedConfiguration(name)
}

//AddConfiguration ....
func AddConfiguration(config configuration.Configuration) {
    configManager.addConfiguration(config)
}

//RemoveConfiguration ...
func RemoveConfiguration(config configuration.Configuration) {
    configManager.removeConfiguration(config)
}

//GetConfigurations ...
func GetConfigurations() []configuration.Configuration {
    return configManager.getConfigurations()
}
