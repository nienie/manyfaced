package configuration

//ConfigurationListener ...
type ConfigurationListener interface {

    //ConfigurationChanged notifies this listener about a change on a monitored configuration object.
    ConfigurationChanged(*ConfigurationEvent)
}