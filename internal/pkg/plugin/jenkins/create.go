package jenkins

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/helm"
	"github.com/devstream-io/devstream/pkg/util/log"
)

// Create creates jenkins with provided options.
func Create(options map[string]interface{}) (map[string]interface{}, error) {
	// 1. config install operations
	runner := &plugininstaller.Runner{
		PreExecuteOperations: []plugininstaller.MutableOperation{
			helm.Validate,
			replaceStroageClass,
		},
		ExecuteOperations: []plugininstaller.BaseOperation{
			helm.DealWithNsWhenInstall,
			preCreate,
			helm.InstallOrUpdate,
			// show how to get pwd of the admin user
			howToGetPasswdOfAdmin,
			// show jenkins url
			showJenkinsUrl,
		},
		TermateOperations: []plugininstaller.BaseOperation{
			helm.DealWithNsWhenInterruption,
		},
		GetStatusOperation: wrapperHelmResourceAndCustomResource(helm.GetPluginStaticStateByReleaseNameWrapper(defaultStatefulsetTplList)),
	}

	// 2. execute installer get status and error
	status, err := runner.Execute(plugininstaller.RawOptions(options))
	if err != nil {
		return nil, err
	}
	log.Debugf("Return map: %v", status)
	return status, nil
}
