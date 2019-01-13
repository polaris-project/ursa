package vm

import "testing"

// TestNewEnvironment - test functionality of environment initializer
func TestNewEnvironment(t *testing.T) {
	env := NewEnvironment(10, 10, 10, 10, 10, 10, 10, false, true) // Init env

	if env == nil { // Check nil env
		t.Fatal("nil environment configuration") // Panic
	}

	t.Log(env) // Log success
}

// TestString - test functionality of environment stringer
func TestString(t *testing.T) {
	env := NewEnvironment(10, 10, 10, 10, 10, 10, 10, false, true) // Init env

	if env == nil || env.String() == "" { // Check nil env
		t.Fatal("nil environment configuration") // Panic
	}

	t.Log(env.String()) // Log success
}

// TestBytes - test functionality of environment byte serialization
func TestBytes(t *testing.T) {
	env := NewEnvironment(10, 10, 10, 10, 10, 10, 10, false, true) // Init env

	if env == nil || env.Bytes() == nil { // Check nil env
		t.Fatal("nil environment configuration") // Panic
	}

	t.Log(env.Bytes()) // Log success
}

// TestNewEnvironmentFromBytes - test functionality of environment byteval marshaling
func TestEnvironmentFromBytes(t *testing.T) {
	env := NewEnvironment(10, 10, 10, 10, 10, 10, 10, false, true) // Init env

	if env == nil || env.Bytes() == nil { // Check nil env
		t.Fatal("nil environment configuration") // Panic
	}

	env, err := EnvironmentFromBytes(env.Bytes()) // Marshal from byte value

	if err != nil { // Check for errors
		t.Fatal(err) // Panic
	}

	t.Log(env) // Log success
}
