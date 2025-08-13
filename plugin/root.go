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

type testTime struct {
	Time    time.Time
	Message string
}

type envVars struct {
	All []string
}

// Func - Implement framework.Call interface
func (r *Root) Func(key string) interface{} {
	switch key {
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
	case "getenv":
		return func(key string) interface{} {
			value := os.Getenv(key)
			return &value
		}
	case "test":
		return func() interface{} {
			return &testTime{Time: time.Now(), Message: "Test message"}
		}
	}
	return nil
}

// Required Implementation - not used
func (r *Root) Configure(m map[string]interface{}) error {
	return nil
}

// Required Implementation - not used
func (r *Root) Get(key string) (interface{}, error) {
	switch key {

	case "envs":
		envMap := make(map[string]string)
		for _, env := range os.Environ() {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				envMap[parts[0]] = parts[1]
			}
		}
		return envMap, nil
		// return &envVars{All: []string{"ENV1", "ENV2"}}, nil
	case "now":
		return &testTime{Time: time.Now()}, nil
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		return &dir, nil
	case "plan_json":
		contents, err := os.ReadFile("../../subjects/schema.json")
		if err != nil {
			return nil, err
		}
		contentsStr := string(contents)
		return &contentsStr, nil

	}
	return nil, nil
}

// Required Implementation - not used
func (r *Root) New(data map[string]interface{}) (framework.Namespace, error) {
	return nil, nil
}
