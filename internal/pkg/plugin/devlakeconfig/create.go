package devlakeconfig

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/log"
)

func Create(options map[string]interface{}) (statemanager.ResourceStatus, error) {
	// Initialize Operator with Operations
	operator := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			validate,
			RenderAuthConfig,
		},
		ExecuteOperations: plugininstaller.ExecuteOperations{
			ApplyConfig,
		},
		TerminateOperations: plugininstaller.TerminateOperations{
			// TODO(dtm): Add your TerminateOperations here.
		},
		GetStatusOperation: GetStatus,
	}

	// Execute all Operations in Operator
	status, err := operator.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return nil, err
	}
	log.Debugf("Return map: %v", status)
	return status, nil
}
