package plugin

import (
	"testing"
)

func TestFunc(t *testing.T) {
	root := Root{}

	// Prepare correct mock inputs
	resource := map[string]interface{}{"type": "aws_instance"} // Correct type for plan, config, state
	mockInputMap := map[string]interface{}{"key": resource}    // Correct type for plan, config, state

	// Test for "plan"
	planFunc := root.Func("test")
	if planFunc == nil {
		t.Fatal("Expected function for 'test', got nil")
	}
	result := planFunc.(func(interface{}) interface{})(mockInputMap)
	namespaceTime := result.(*namespaceTime)
	println(namespaceTime.Message)
	println(namespaceTime.Time.GoString())
	// if _, ok := result.(*ResourcesNs); !ok {
	// 	t.Fatalf("Expected *main.ResourcesNs for 'plan', got %T", result)
	// }
}
