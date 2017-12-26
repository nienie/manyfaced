package source

import (
    "github.com/magiconair/properties"
    "github.com/nienie/manyfaced/poll"
)

//URLConfigurationSource ...
type URLConfigurationSource struct {
    configURLs []string
}

//NewURLConfigurationSource ...
func NewURLConfigurationSource(configUrls []string) *URLConfigurationSource {
    return &URLConfigurationSource{
        configURLs:         configUrls,
    }
}

//Poll ...
func (source *URLConfigurationSource)Poll(initial bool, checkPoint interface{}) (*poll.PolledResult, error) {
    if source.configURLs == nil || len(source.configURLs) == 0 {
        return poll.NewFullPolledResult(nil), nil
    }
    complete := map[string]interface{}{}
    props, err := properties.LoadURLs(source.configURLs, false)
    if err != nil {
        return nil, err
    }
    m := props.Map()
    for key, val := range m {
        complete[key] = val
    }
    return poll.NewFullPolledResult(complete), nil
}