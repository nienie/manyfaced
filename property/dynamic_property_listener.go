package property

//DynamicPropertyListener ...
type DynamicPropertyListener struct {}

//ConfigSourceLoaded ...
func (l *DynamicPropertyListener)ConfigurationSourceLoaded(source interface{}) {
    updateAllProperties()
}

//AddProperty ...
func (l *DynamicPropertyListener)AddProperty(source interface{}, name string, value interface{}, beforeUpdate bool) error {
    return updateProperty(name, value)
}

//SetProperty ...
func (l *DynamicPropertyListener)SetProperty(source interface{}, name string, value interface{}, beforeUpdate bool) error{
    return updateProperty(name, value)
}

//ClearProperty ...
func (l *DynamicPropertyListener)ClearProperty(source interface{}, name string, value interface{}, beforeUpdate bool) error {
    if !beforeUpdate {
        return updateProperty(name, value)
    }
    return nil
}

//Clear ...
func (l *DynamicPropertyListener)Clear(source interface{}, beforeUpdate bool) {
    if !beforeUpdate {
        updateAllProperties()
    }
}
