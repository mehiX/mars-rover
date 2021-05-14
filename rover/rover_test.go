package rover

import "testing"

func TestCreateRoverSuccess(t *testing.T) {

	r, err := NewRover("name", 1, 4, "n")
	if err != nil {
		t.Fatalf("Expected no error. Got: %v", err)
	}

	if r.Name != "name" {
		t.Error("Name not set correctly")
	}
}

func TestCreateRoverFail(t *testing.T) {

	_, err := NewRover("name", 1, 4, "x")
	if err == nil {
		t.Errorf("Expecting an error but got none")
	}
}

func TestDecodeDirection(t *testing.T) {

	type scenario struct {
		input       string
		expectError bool
		expect      direction
	}

	scenarios := []scenario{
		scenario{"n", false, DIR_N},
		scenario{"N", false, DIR_N},
		scenario{"s", false, DIR_S},
		scenario{"E", false, DIR_E},
		scenario{"W", false, DIR_W},
		scenario{"x", true, DIR_UNKNOWN},
	}

	for _, s := range scenarios {
		d, err := decodeDirection(s.input)
		if s.expectError && err == nil {
			t.Errorf("Input: %s, expected error and got none", s.input)
		}

		if !s.expectError && err != nil {
			t.Errorf("Input: %s, not expecting error. Got: %v", s.input, err)
		}

		if d != s.expect {
			t.Errorf("Input: %s, expected: %s, got: %s", s.input, s.expect, d)
		}
	}
}

func TestDirectionToString(t *testing.T) {

	type scenario struct {
		d        direction
		expected string
	}

	scenarios := []scenario{
		scenario{DIR_N, "N"},
		scenario{DIR_S, "S"},
		scenario{DIR_E, "E"},
		scenario{DIR_W, "W"},
		scenario{DIR_UNKNOWN, "UNKNOWN"},
	}

	for _, s := range scenarios {
		if s.d.String() != s.expected {
			t.Errorf("Wrong direction string. Expected: %s, got: %s", s.expected, s.d.String())
		}
	}
}
