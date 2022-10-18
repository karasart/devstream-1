package main

import (
	"github.com/devstream-io/devstream/internal/pkg/plugin/devlake"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// NAME is the name of this DevStream plugin.
const NAME = "devlake"

// Plugin is the type used by DevStream core. It's a string.
type Plugin string

// Create implements the create of devlake.
func (p Plugin) Create(options map[string]interface{}) (statemanager.ResourceStatus, error) {
	return devlake.Create(options)
}

// Update implements the update of devlake.
func (p Plugin) Update(options map[string]interface{}) (statemanager.ResourceStatus, error) {
	return devlake.Update(options)
}

// Delete implements the delete of devlake.
func (p Plugin) Delete(options map[string]interface{}) (bool, error) {
	return devlake.Delete(options)
}

// Read implements the read of devlake.
func (p Plugin) Read(options map[string]interface{}) (statemanager.ResourceStatus, error) {
	return devlake.Read(options)
}

// DevStreamPlugin is the exported variable used by the DevStream core.
var DevStreamPlugin Plugin

func main() {
	log.Infof("%T: %s is a plugin for DevStream. Use it with DevStream.\n", DevStreamPlugin, NAME)
}
