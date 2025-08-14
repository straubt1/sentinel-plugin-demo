package plugin

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestGetAllEnvs(t *testing.T) {
	// Create a new Root instance
	root := &Root{}

	// Test that getallenvs function exists
	t.Run("FunctionExists", func(t *testing.T) {
		fn := root.Func("getallenvs")
		if fn == nil {
			t.Fatal("getallenvs function should not be nil")
		}
	})

	// Test that getallenvs returns a callable function
	t.Run("ReturnsCallableFunction", func(t *testing.T) {
		fn := root.Func("getallenvs")
		if fn == nil {
			t.Fatal("getallenvs function should not be nil")
		}

		// Call the function
		callable, ok := fn.(func() interface{})
		if !ok {
			t.Fatal("getallenvs should return a callable function")
		}

		result := callable()
		if result == nil {
			t.Fatal("getallenvs function call should not return nil")
		}
	})

	// Test that getallenvs returns a map pointer
	t.Run("ReturnsMapPointer", func(t *testing.T) {
		fn := root.Func("getallenvs")
		callable := fn.(func() interface{})
		result := callable()

		envMapPtr, ok := result.(*map[string]string)
		if !ok {
			t.Fatal("getallenvs should return a pointer to map[string]string")
		}

		if envMapPtr == nil {
			t.Fatal("returned map pointer should not be nil")
		}
	})

	// Test that getallenvs contains expected environment variables
	t.Run("ContainsEnvironmentVariables", func(t *testing.T) {
		// Set a test environment variable
		testKey := "TEST_GETALLENVS_VAR"
		testValue := "test_value_123"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		fn := root.Func("getallenvs")
		callable := fn.(func() interface{})
		result := callable()
		envMapPtr := result.(*map[string]string)
		envMap := *envMapPtr

		// Check that our test variable is present
		if value, exists := envMap[testKey]; !exists {
			t.Errorf("Expected environment variable %s to be present", testKey)
		} else if value != testValue {
			t.Errorf("Expected %s=%s, got %s=%s", testKey, testValue, testKey, value)
		}

		// Check that the map is not empty (should contain system env vars)
		if len(envMap) == 0 {
			t.Error("Environment map should not be empty")
		}
	})

	// Test that getallenvs handles environment variables with equals signs in values
	t.Run("HandlesEqualsInValues", func(t *testing.T) {
		testKey := "TEST_EQUALS_VAR"
		testValue := "value=with=equals=signs"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		fn := root.Func("getallenvs")
		callable := fn.(func() interface{})
		result := callable()
		envMapPtr := result.(*map[string]string)
		envMap := *envMapPtr

		if value, exists := envMap[testKey]; !exists {
			t.Errorf("Expected environment variable %s to be present", testKey)
		} else if value != testValue {
			t.Errorf("Expected %s=%s, got %s=%s", testKey, testValue, testKey, value)
		}
	})

	// Test that getallenvs handles empty environment variables
	t.Run("HandlesEmptyValues", func(t *testing.T) {
		testKey := "TEST_EMPTY_VAR"
		testValue := ""
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		fn := root.Func("getallenvs")
		callable := fn.(func() interface{})
		result := callable()
		envMapPtr := result.(*map[string]string)
		envMap := *envMapPtr

		if value, exists := envMap[testKey]; !exists {
			t.Errorf("Expected environment variable %s to be present", testKey)
		} else if value != testValue {
			t.Errorf("Expected %s=%s, got %s=%s", testKey, testValue, testKey, value)
		}
	})

	// Test consistency across multiple calls
	t.Run("ConsistentAcrossCalls", func(t *testing.T) {
		testKey := "TEST_CONSISTENCY_VAR"
		testValue := "consistent_value"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		fn := root.Func("getallenvs")
		callable := fn.(func() interface{})

		// First call
		result1 := callable()
		envMapPtr1 := result1.(*map[string]string)
		envMap1 := *envMapPtr1

		// Second call
		result2 := callable()
		envMapPtr2 := result2.(*map[string]string)
		envMap2 := *envMapPtr2

		// Both calls should return the same value for our test variable
		if envMap1[testKey] != envMap2[testKey] {
			t.Error("getallenvs should return consistent results across calls")
		}
	})
}

func TestGetEnv(t *testing.T) {
	// Create a new Root instance
	root := &Root{}

	// Test that getenv function exists
	t.Run("FunctionExists", func(t *testing.T) {
		fn := root.Func("getenv")
		if fn == nil {
			t.Fatal("getenv function should not be nil")
		}
	})

	// Test that getenv returns a callable function
	t.Run("ReturnsCallableFunction", func(t *testing.T) {
		fn := root.Func("getenv")
		if fn == nil {
			t.Fatal("getenv function should not be nil")
		}

		// Call the function
		callable, ok := fn.(func(string) interface{})
		if !ok {
			t.Fatal("getenv should return a callable function that takes a string parameter")
		}

		result := callable("PATH")
		if result == nil {
			t.Fatal("getenv function call should not return nil")
		}
	})

	// Test that getenv returns a string pointer
	t.Run("ReturnsStringPointer", func(t *testing.T) {
		fn := root.Func("getenv")
		callable := fn.(func(string) interface{})
		result := callable("PATH")

		valuePtr, ok := result.(*string)
		if !ok {
			t.Fatal("getenv should return a pointer to string")
		}

		if valuePtr == nil {
			t.Fatal("returned string pointer should not be nil")
		}
	})

	// Test that getenv retrieves existing environment variables
	t.Run("RetrievesExistingVariable", func(t *testing.T) {
		testKey := "TEST_GETENV_VAR"
		testValue := "test_value_456"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		fn := root.Func("getenv")
		callable := fn.(func(string) interface{})
		result := callable(testKey)
		valuePtr := result.(*string)

		if *valuePtr != testValue {
			t.Errorf("Expected %s, got %s", testValue, *valuePtr)
		}
	})

	// Test that getenv returns empty string for non-existent variables
	t.Run("ReturnsEmptyForNonExistent", func(t *testing.T) {
		nonExistentKey := "NON_EXISTENT_VAR_12345"

		fn := root.Func("getenv")
		callable := fn.(func(string) interface{})
		result := callable(nonExistentKey)
		valuePtr := result.(*string)

		if *valuePtr != "" {
			t.Errorf("Expected empty string for non-existent variable, got %s", *valuePtr)
		}
	})

	// Test that getenv handles empty environment variables
	t.Run("HandlesEmptyVariables", func(t *testing.T) {
		testKey := "TEST_EMPTY_GETENV_VAR"
		os.Setenv(testKey, "")
		defer os.Unsetenv(testKey)

		fn := root.Func("getenv")
		callable := fn.(func(string) interface{})
		result := callable(testKey)
		valuePtr := result.(*string)

		if *valuePtr != "" {
			t.Errorf("Expected empty string, got %s", *valuePtr)
		}
	})

	// Test consistency across multiple calls
	t.Run("ConsistentAcrossCalls", func(t *testing.T) {
		testKey := "TEST_CONSISTENCY_GETENV_VAR"
		testValue := "consistent_getenv_value"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		fn := root.Func("getenv")
		callable := fn.(func(string) interface{})

		// First call
		result1 := callable(testKey)
		valuePtr1 := result1.(*string)

		// Second call
		result2 := callable(testKey)
		valuePtr2 := result2.(*string)

		if *valuePtr1 != *valuePtr2 {
			t.Error("getenv should return consistent results across calls")
		}
	})
}

func TestGetFile(t *testing.T) {
	// Create a new Root instance
	root := &Root{}

	// Test that getfile function exists
	t.Run("FunctionExists", func(t *testing.T) {
		fn := root.Func("getfile")
		if fn == nil {
			t.Fatal("getfile function should not be nil")
		}
	})

	// Test that getfile returns a callable function
	t.Run("ReturnsCallableFunction", func(t *testing.T) {
		fn := root.Func("getfile")
		if fn == nil {
			t.Fatal("getfile function should not be nil")
		}

		// Call the function
		callable, ok := fn.(func(string) interface{})
		if !ok {
			t.Fatal("getfile should return a callable function that takes a string parameter")
		}

		// Test with a non-existent file (should return nil)
		result := callable("/non/existent/file")
		if result != nil {
			t.Log("getfile returns nil for non-existent files, which is expected")
		}
	})

	// Test that getfile reads existing files
	t.Run("ReadsExistingFile", func(t *testing.T) {
		// Create a temporary file
		tempDir := os.TempDir()
		tempFile := filepath.Join(tempDir, "test_getfile.txt")
		testContent := "Hello, this is test content for getfile!"

		err := os.WriteFile(tempFile, []byte(testContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(tempFile)

		fn := root.Func("getfile")
		callable := fn.(func(string) interface{})
		result := callable(tempFile)

		if result == nil {
			t.Fatal("getfile should not return nil for existing file")
		}

		contentPtr, ok := result.(*string)
		if !ok {
			t.Fatal("getfile should return a pointer to string")
		}

		if *contentPtr != testContent {
			t.Errorf("Expected %s, got %s", testContent, *contentPtr)
		}
	})

	// Test that getfile returns nil for non-existent files
	t.Run("ReturnsNilForNonExistentFile", func(t *testing.T) {
		nonExistentFile := "/absolutely/non/existent/file/path/12345.txt"

		fn := root.Func("getfile")
		callable := fn.(func(string) interface{})
		result := callable(nonExistentFile)

		if result != nil {
			t.Error("getfile should return nil for non-existent files")
		}
	})

	// Test that getfile handles empty files
	t.Run("HandlesEmptyFiles", func(t *testing.T) {
		// Create an empty temporary file
		tempDir := os.TempDir()
		tempFile := filepath.Join(tempDir, "test_empty_getfile.txt")

		err := os.WriteFile(tempFile, []byte(""), 0644)
		if err != nil {
			t.Fatalf("Failed to create empty test file: %v", err)
		}
		defer os.Remove(tempFile)

		fn := root.Func("getfile")
		callable := fn.(func(string) interface{})
		result := callable(tempFile)

		if result == nil {
			t.Fatal("getfile should not return nil for existing empty file")
		}

		contentPtr, ok := result.(*string)
		if !ok {
			t.Fatal("getfile should return a pointer to string")
		}

		if *contentPtr != "" {
			t.Errorf("Expected empty string, got %s", *contentPtr)
		}
	})

	// Test that getfile handles binary content
	t.Run("HandlesBinaryContent", func(t *testing.T) {
		// Create a file with binary content
		tempDir := os.TempDir()
		tempFile := filepath.Join(tempDir, "test_binary_getfile.bin")
		binaryContent := []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD}

		err := os.WriteFile(tempFile, binaryContent, 0644)
		if err != nil {
			t.Fatalf("Failed to create binary test file: %v", err)
		}
		defer os.Remove(tempFile)

		fn := root.Func("getfile")
		callable := fn.(func(string) interface{})
		result := callable(tempFile)

		if result == nil {
			t.Fatal("getfile should not return nil for existing binary file")
		}

		contentPtr, ok := result.(*string)
		if !ok {
			t.Fatal("getfile should return a pointer to string")
		}

		expectedString := string(binaryContent)
		if *contentPtr != expectedString {
			t.Error("getfile should handle binary content correctly")
		}
	})

	// Test that getfile returns nil for directories
	t.Run("ReturnsNilForDirectories", func(t *testing.T) {
		// Use a known directory (temp directory)
		tempDir := os.TempDir()

		fn := root.Func("getfile")
		callable := fn.(func(string) interface{})
		result := callable(tempDir)

		if result != nil {
			t.Error("getfile should return nil when trying to read a directory")
		}
	})

	// Test consistency across multiple calls
	t.Run("ConsistentAcrossCalls", func(t *testing.T) {
		// Create a temporary file
		tempDir := os.TempDir()
		tempFile := filepath.Join(tempDir, "test_consistency_getfile.txt")
		testContent := "Consistency test content"

		err := os.WriteFile(tempFile, []byte(testContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
		defer os.Remove(tempFile)

		fn := root.Func("getfile")
		callable := fn.(func(string) interface{})

		// First call
		result1 := callable(tempFile)
		contentPtr1 := result1.(*string)

		// Second call
		result2 := callable(tempFile)
		contentPtr2 := result2.(*string)

		if *contentPtr1 != *contentPtr2 {
			t.Error("getfile should return consistent results across calls")
		}
	})
}

func TestGetEnvs(t *testing.T) {
	// Create a new Root instance
	root := &Root{}

	// Test that envs property exists
	t.Run("PropertyExists", func(t *testing.T) {
		result, err := root.Get("envs")
		if err != nil {
			t.Fatalf("envs property should not return error: %v", err)
		}
		if result == nil {
			t.Fatal("envs property should not return nil")
		}
	})

	// Test that envs returns a map
	t.Run("ReturnsMap", func(t *testing.T) {
		result, err := root.Get("envs")
		if err != nil {
			t.Fatalf("envs property should not return error: %v", err)
		}

		envMap, ok := result.(map[string]string)
		if !ok {
			t.Fatal("envs should return a map[string]string")
		}

		// Check that the map is not empty (should contain system env vars)
		if len(envMap) == 0 {
			t.Error("Environment map should not be empty")
		}
	})

	// Test that envs contains expected environment variables
	t.Run("ContainsEnvironmentVariables", func(t *testing.T) {
		// Set a test environment variable
		testKey := "TEST_GET_ENVS_VAR"
		testValue := "test_value_789"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		result, err := root.Get("envs")
		if err != nil {
			t.Fatalf("envs property should not return error: %v", err)
		}

		envMap := result.(map[string]string)

		// Check that our test variable is present
		if value, exists := envMap[testKey]; !exists {
			t.Errorf("Expected environment variable %s to be present", testKey)
		} else if value != testValue {
			t.Errorf("Expected %s=%s, got %s=%s", testKey, testValue, testKey, value)
		}
	})

	// Test that envs handles environment variables with equals signs in values
	t.Run("HandlesEqualsInValues", func(t *testing.T) {
		testKey := "TEST_GET_ENVS_EQUALS_VAR"
		testValue := "value=with=multiple=equals"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		result, err := root.Get("envs")
		if err != nil {
			t.Fatalf("envs property should not return error: %v", err)
		}

		envMap := result.(map[string]string)

		if value, exists := envMap[testKey]; !exists {
			t.Errorf("Expected environment variable %s to be present", testKey)
		} else if value != testValue {
			t.Errorf("Expected %s=%s, got %s=%s", testKey, testValue, testKey, value)
		}
	})

	// Test that envs handles empty environment variables
	t.Run("HandlesEmptyValues", func(t *testing.T) {
		testKey := "TEST_GET_ENVS_EMPTY_VAR"
		testValue := ""
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		result, err := root.Get("envs")
		if err != nil {
			t.Fatalf("envs property should not return error: %v", err)
		}

		envMap := result.(map[string]string)

		if value, exists := envMap[testKey]; !exists {
			t.Errorf("Expected environment variable %s to be present", testKey)
		} else if value != testValue {
			t.Errorf("Expected %s=%s, got %s=%s", testKey, testValue, testKey, value)
		}
	})

	// Test consistency across multiple calls
	t.Run("ConsistentAcrossCalls", func(t *testing.T) {
		testKey := "TEST_GET_ENVS_CONSISTENCY_VAR"
		testValue := "consistent_envs_value"
		os.Setenv(testKey, testValue)
		defer os.Unsetenv(testKey)

		// First call
		result1, err1 := root.Get("envs")
		if err1 != nil {
			t.Fatalf("First envs call should not return error: %v", err1)
		}
		envMap1 := result1.(map[string]string)

		// Second call
		result2, err2 := root.Get("envs")
		if err2 != nil {
			t.Fatalf("Second envs call should not return error: %v", err2)
		}
		envMap2 := result2.(map[string]string)

		// Both calls should return the same value for our test variable
		if envMap1[testKey] != envMap2[testKey] {
			t.Error("envs should return consistent results across calls")
		}
	})
}

func TestGetNow(t *testing.T) {
	// Create a new Root instance
	root := &Root{}

	// Test that now property exists
	t.Run("PropertyExists", func(t *testing.T) {
		result, err := root.Get("now")
		if err != nil {
			t.Fatalf("now property should not return error: %v", err)
		}
		if result == nil {
			t.Fatal("now property should not return nil")
		}
	})

	// Test that now returns a testTime pointer
	t.Run("ReturnsTestTimePointer", func(t *testing.T) {
		result, err := root.Get("now")
		if err != nil {
			t.Fatalf("now property should not return error: %v", err)
		}

		timePtr, ok := result.(*testTime)
		if !ok {
			t.Fatal("now should return a pointer to testTime")
		}

		if timePtr == nil {
			t.Fatal("returned testTime pointer should not be nil")
		}
	})

	// Test that now returns current time
	t.Run("ReturnsCurrentTime", func(t *testing.T) {
		beforeCall := time.Now()

		result, err := root.Get("now")
		if err != nil {
			t.Fatalf("now property should not return error: %v", err)
		}

		afterCall := time.Now()
		timePtr := result.(*testTime)

		// The returned time should be between beforeCall and afterCall
		if timePtr.Time.Before(beforeCall) || timePtr.Time.After(afterCall) {
			t.Error("now should return a time between the call boundaries")
		}
	})

	// Test that now has empty message (based on implementation)
	t.Run("HasEmptyMessage", func(t *testing.T) {
		result, err := root.Get("now")
		if err != nil {
			t.Fatalf("now property should not return error: %v", err)
		}

		timePtr := result.(*testTime)

		// Based on the implementation, Get("now") doesn't set Message
		if timePtr.Message != "" {
			t.Errorf("Expected empty message, got %s", timePtr.Message)
		}
	})

	// Test that multiple calls return different times
	t.Run("MultipleCallsReturnDifferentTimes", func(t *testing.T) {
		// First call
		result1, err1 := root.Get("now")
		if err1 != nil {
			t.Fatalf("First now call should not return error: %v", err1)
		}
		timePtr1 := result1.(*testTime)

		// Small delay to ensure different timestamps
		time.Sleep(1 * time.Millisecond)

		// Second call
		result2, err2 := root.Get("now")
		if err2 != nil {
			t.Fatalf("Second now call should not return error: %v", err2)
		}
		timePtr2 := result2.(*testTime)

		// Times should be different (second call should be after first)
		if !timePtr2.Time.After(timePtr1.Time) {
			t.Error("Second call to now should return a later time than first call")
		}
	})
}

func TestGetPwd(t *testing.T) {
	// Create a new Root instance
	root := &Root{}

	// Test that pwd property exists
	t.Run("PropertyExists", func(t *testing.T) {
		result, err := root.Get("pwd")
		if err != nil {
			t.Fatalf("pwd property should not return error: %v", err)
		}
		if result == nil {
			t.Fatal("pwd property should not return nil")
		}
	})

	// Test that pwd returns a string pointer
	t.Run("ReturnsStringPointer", func(t *testing.T) {
		result, err := root.Get("pwd")
		if err != nil {
			t.Fatalf("pwd property should not return error: %v", err)
		}

		dirPtr, ok := result.(*string)
		if !ok {
			t.Fatal("pwd should return a pointer to string")
		}

		if dirPtr == nil {
			t.Fatal("returned string pointer should not be nil")
		}
	})

	// Test that pwd returns valid directory path
	t.Run("ReturnsValidDirectory", func(t *testing.T) {
		result, err := root.Get("pwd")
		if err != nil {
			t.Fatalf("pwd property should not return error: %v", err)
		}

		dirPtr := result.(*string)
		dir := *dirPtr

		// Check that it's not empty
		if dir == "" {
			t.Error("pwd should not return empty string")
		}

		// Check that the directory exists
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("pwd returned directory %s does not exist", dir)
		}
	})

	// Test that pwd matches os.Getwd()
	t.Run("MatchesOsGetwd", func(t *testing.T) {
		// Get expected working directory
		expectedDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get working directory: %v", err)
		}

		result, err := root.Get("pwd")
		if err != nil {
			t.Fatalf("pwd property should not return error: %v", err)
		}

		dirPtr := result.(*string)
		actualDir := *dirPtr

		if actualDir != expectedDir {
			t.Errorf("Expected pwd to return %s, got %s", expectedDir, actualDir)
		}
	})

	// Test consistency across multiple calls
	t.Run("ConsistentAcrossCalls", func(t *testing.T) {
		// First call
		result1, err1 := root.Get("pwd")
		if err1 != nil {
			t.Fatalf("First pwd call should not return error: %v", err1)
		}
		dirPtr1 := result1.(*string)

		// Second call
		result2, err2 := root.Get("pwd")
		if err2 != nil {
			t.Fatalf("Second pwd call should not return error: %v", err2)
		}
		dirPtr2 := result2.(*string)

		// Both calls should return the same directory
		if *dirPtr1 != *dirPtr2 {
			t.Error("pwd should return consistent results across calls")
		}
	})

	// Test pwd after changing directory
	t.Run("ReflectsDirectoryChanges", func(t *testing.T) {
		// Get original directory
		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get original working directory: %v", err)
		}

		// Get pwd before change
		result1, err1 := root.Get("pwd")
		if err1 != nil {
			t.Fatalf("First pwd call should not return error: %v", err1)
		}
		dirPtr1 := result1.(*string)

		// Change to temp directory
		tempDir := os.TempDir()
		err = os.Chdir(tempDir)
		if err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}

		// Restore original directory at the end
		defer func() {
			os.Chdir(originalDir)
		}()

		// Get pwd after change
		result2, err2 := root.Get("pwd")
		if err2 != nil {
			t.Fatalf("Second pwd call should not return error: %v", err2)
		}
		dirPtr2 := result2.(*string)

		// pwd should reflect the directory change
		if *dirPtr1 == *dirPtr2 {
			t.Error("pwd should reflect directory changes")
		}

		// Resolve symbolic links for both paths to ensure accurate comparison
		expectedDir, err := filepath.EvalSymlinks(tempDir)
		if err != nil {
			// If EvalSymlinks fails, fall back to the original path
			expectedDir = tempDir
		}

		actualDir, err := filepath.EvalSymlinks(*dirPtr2)
		if err != nil {
			// If EvalSymlinks fails, use the path as-is
			actualDir = *dirPtr2
		}

		// New pwd should match the resolved temp directory
		if actualDir != expectedDir {
			t.Errorf("Expected pwd to return %s after directory change, got %s", expectedDir, actualDir)
		}
	})
}
