package function

import (
	faasflow "github.com/s8sg/faas-flow"
	"os"
)

func getResizerFunction() string {
	resizer := os.Getenv("resizer-function")
	if resizer == "" {
		resizer = "resizeimagemagick"
	}
	return resizer
}

func getFilterFunction(context *faasflow.Context) string {
	filter := context.Query.Get("filter")
	if filter == "" {
		filter = "coherent-line-drawing"
	}
	return filter
}

// Define provide definiton of the workflow
func Define(flow *faasflow.Workflow, context *faasflow.Context) (err error) {
	resizer := getResizerFunction()
	filter := getFilterFunction(context)
	// resize and apply filter
	flow.SyncNode().Apply(resizer).Apply(filter)
	return
}

// DefineStateStore provides the override of the default StateStore
func DefineStateStore() (faasflow.StateStore, error) {
	return nil, nil
}

// ProvideDataStore provides the override of the default DataStore
func DefineDataStore() (faasflow.DataStore, error) {
	return nil, nil
}
