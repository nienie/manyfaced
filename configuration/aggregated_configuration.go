package configuration

type AggregatedConfiguration interface {

    Configuration

    //AddConfiguration ...
    AddConfiguration(Configuration)

    //AddNamedConfiguration ...
    AddNamedConfiguration(config Configuration, name string)

    //GetConfigurationByName ...
    GetConfigurationByName(name string) Configuration

    //GetNumberOfConfigurations ...
    GetNumberOfConfigurations() int

    //GetConfigurations ...
    GetConfigurations() []Configuration

    //RemoveConfigurationByName ...
    RemoveConfigurationByName(name string) Configuration

    //RemoveConfiguration ...
    RemoveConfiguration(Configuration) bool
}
