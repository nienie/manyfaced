package source

import (
	"fmt"
	"github.com/nienie/manyfaced/parser"
	"github.com/nienie/manyfaced/poll"
	"time"
)

//FileConfigurationSource ...
type FileConfigurationSource struct {
	configFiles []string
	pollTimeout time.Duration
}

//NewFileConfigurationSource ...
func NewFileConfigurationSource(configFiles []string, pollTimeout time.Duration) *FileConfigurationSource {
	return &FileConfigurationSource{
		configFiles: configFiles,
		pollTimeout: pollTimeout,
	}
}

//Poll ...
func (source *FileConfigurationSource) Poll(initial bool, checkPoint interface{}) (*poll.PolledResult, error) {
	num := len(source.configFiles)
	completes := make([]map[string]interface{}, num)
	errors := make([]error, num)
	ch := make(chan struct{}, num)
	for i, configFile := range source.configFiles {
		go func(index int, configFile string) {
			defer func() {
				if r := recover(); r != nil {
					errors[index] = fmt.Errorf("%+v", r)
				}
				ch <- struct{}{}
			}()
			completes[index], errors[index] = parser.ParseConfigFile(configFile)
		}(i, configFile)
	}
	var finishedNum int
	ticker := time.NewTicker(source.pollTimeout)
	for {
		select {
		case <-ch:
			finishedNum++
			if finishedNum >= num {
				allCompletes := make(map[string]interface{})
				for i := 0; i < num; i++ {
					if errors[i] != nil {
						continue
					}
					for key, val := range completes[i] {
						allCompletes[key] = val
					}
				}
				return poll.NewFullPolledResult(allCompletes), nil
			}
		case <-ticker.C:
			return nil, fmt.Errorf("Poll timeout")
		}
	}
}
