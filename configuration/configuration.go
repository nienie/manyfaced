package configuration

//Accessor ...
type Accessor interface {
    //AddProperty add a property to the configuration. If it already exists then the value stated
    //here will be added to configuration entry.
    AddProperty(key string, value interface{})

    //SetProperty set a property, this will replace any previously set values.
    SetProperty(key string, value interface{})

    //GetProperty gets a property from the configuration.
    GetProperty(key string) interface{}

    //ClearProperty removes a property from the configuration.
    ClearProperty(key string)

    //Clear removes all properties from the configuration.
    Clear()
}

//Configuration The main Configuration interface.
type Configuration interface {
    Accessor

    //Subset returns a decorator configuration containing every key from the current
    //configuration that stats with the specified prefix. The prefix is removed from
    //the keys in the subset.
    Subset(prefix string) Configuration

    //IsEmpty check if the configuration is empty.
    IsEmpty() bool

    //ContainsKey check if the configuration contains the specified key.
    ContainsKey(key string) bool

    //GetKeys gets the list of the keys contained in the configuration.
    GetKeys() []string

    //GetBool gets a bool associated with the given configuration key.
    GetBool(key string, defaultValue bool) bool

    //MustGetBool gets a bool associated with the given configuration key. If failed, a panic is thrown.
    MustGetBool(key string) bool

    //GetFloat64 gets a float64 associated with the given configuration key.
    GetFloat64(key string, defaultValue float64) float64

    //MustGetFloat64 gets a float64 associated with the given configuration key. If failed, a panic is thrown.
    MustGetFloat64(key string) float64

    //GetInt gets a int associated with the given configuration key.
    GetInt(key string, defaultValue int) int

    //MustGetInt gets a int associated with the given configuration key. If failed, a panic is thrown.
    MustGetInt(key string) int

    //GetInt32 gets a int32 associated with the given configuration key.
    GetInt32(key string, defaultValue int32) int32

    //MustGetInt32 gets a int32 associated with the given configuration key.
    MustGetInt32(key string) int32

    //GetInt64 gets a int64 associated with the given configuration key.
    GetInt64(key string, defaultValue int64) int64

    //MustGetInt64 gets a int64 associated with the given configuration key. If failed, a panic is thrown.
    MustGetInt64(key string) int64

    //GetString gets a string associated with the given configuration key.
    GetString(key string, defaultValue string) string

    //MustGetString gets a string associated with the given configuration key. If failed, a panic is thrown.
    MustGetString(key string) string

    //AddConfigurationListener add a configuration listener.
    AddConfigurationListener(Listener)

    //RemoveConfigurationListener remove configuration listener.
    RemoveConfigurationListener(Listener)

    //GetConfigurationListeners get all configuration listeners.
    GetConfigurationListeners() []Listener
}