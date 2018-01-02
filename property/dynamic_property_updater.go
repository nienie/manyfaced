package property

import (
	"github.com/nienie/manyfaced/configuration"
	"github.com/nienie/manyfaced/poll"
)

//DynamicPropertyUpdater apply the DynamicUpdatedResult to the configuration.
type DynamicPropertyUpdater struct{}

//UpdateProperties updates the properties in the config param given the contents of the result param.
func (updater *DynamicPropertyUpdater) UpdateProperties(result poll.DynamicUpdatedResult, config configuration.Configuration) {
	if result == nil || result.HasChanged() == false {
		return
	}

	if result.IsIncremental() {
		addedProps := result.GetAdded()
		for key, val := range addedProps {
			updater.addOrChangedProperty(key, val, config)
		}
		changedProps := result.GetChanged()
		for key, val := range changedProps {
			updater.addOrChangedProperty(key, val, config)
		}
		deletedProps := result.GetDeleted()
		for key := range deletedProps {
			updater.deleteProperty(key, config)
		}
		return
	}

	completeProps := result.GetComplete()
	if completeProps == nil {
		return
	}

	for key, val := range completeProps {
		updater.addOrChangedProperty(key, val, config)
	}

	existingKeys := config.GetKeys()
	for _, existingKey := range existingKeys {
		if _, ok := completeProps[existingKey]; !ok {
			updater.deleteProperty(existingKey, config)
		}
	}
	return
}

func (updater *DynamicPropertyUpdater) addOrChangedProperty(name string, newValue interface{}, config configuration.Configuration) {
	config.SetProperty(name, newValue)
}

func (updater *DynamicPropertyUpdater) deleteProperty(key string, config configuration.Configuration) {
	config.ClearProperty(key)
}
