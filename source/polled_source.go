package source

import (
	"github.com/nienie/manyfaced/poll"
)

//PolledConfigurationSource the definition of configuration source that brings dynamic changes to
// the configuration via polling.
type PolledConfigurationSource interface {

	//Poll the configuration source to get the latest content.
	//initial true if the operation is the first pool
	//checkPoint that is used to determine the starting point if the result returned is incremental.
	//nil if there is no check point or the caller wishes to get the full content.
	//the return value represents the content of the configuration which may be full or incremental.
	Poll(initial bool, checkPoint interface{}) (*poll.PolledResult, error)
}
