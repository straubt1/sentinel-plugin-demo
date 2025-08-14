package plugin

import (
	"os"
	"strings"
	"time"

	sdk "github.com/hashicorp/sentinel-sdk"
	"github.com/hashicorp/sentinel-sdk/framework"
)

type Root struct {
}

func New() sdk.Plugin {
	return &framework.Plugin{
		Root: &Root{},
	}
}

// Return structs
type testTime struct {
	Time    time.Time
	Message string
}

type envVars struct {
	All []string
}

// This is where we can add new functions
// Example in Sentinel, key == "getallenvs":
//
// import "plugin-demo" as pd
// pd.getallenvs()
func (r *Root) Func(key string) interface{} {
	switch key {
	// Get all environment variables, return a map
	case "getallenvs":
		return func() interface{} {
			envMap := make(map[string]string)
			for _, env := range os.Environ() {
				parts := strings.SplitN(env, "=", 2)
				if len(parts) == 2 {
					envMap[parts[0]] = parts[1]
				}
			}
			return &envMap
		}
	// Get a specific environment variable, return its value or empty if not found
	case "getenv":
		return func(key string) interface{} {
			value := os.Getenv(key)
			return &value
		}
	case "getfile":
		return func(path string) interface{} {
			contents, err := os.ReadFile(path)
			if err != nil {
				return nil // File not found or inaccessible
			}
			contentsStr := string(contents)
			return &contentsStr
		}
	// Test function, return current time and a message
	case "test":
		return func() interface{} {
			return &testTime{Time: time.Now(), Message: "Test message"}
		}
	}
	return nil
}

// This is where we can add new properties
// Example in Sentinel, key == "now":
//
// import "plugin-demo" as pd
// pd.now
func (r *Root) Get(key string) (interface{}, error) {
	switch key {
	// Get all environment variables as a property, return a map
	case "envs":
		envMap := make(map[string]string)
		for _, env := range os.Environ() {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				envMap[parts[0]] = parts[1]
			}
		}
		return envMap, nil
	// Get current time as a property
	case "now":
		return &testTime{Time: time.Now()}, nil
	// Get current working directory as a property
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		return &dir, nil
	}
	return nil, nil
}

// Required Implementation - not used
func (r *Root) Configure(m map[string]interface{}) error {
	return nil
}

// Required Implementation - not used
func (r *Root) New(data map[string]interface{}) (framework.Namespace, error) {
	return nil, nil
}
