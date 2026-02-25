package input

import (
	"strings"
	"testing"
)

func TestParseJSONBasic(t *testing.T) {
	json := `{"project": {}}`
	r := strings.NewReader(json)
	root, err := ParseJSON(r)

	if err != nil {
		t.Errorf("ParseJSON should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	if root.Name != "project" {
		t.Errorf("Root name should be 'project', got %s", root.Name)
	}
}

func TestParseJSONEmptyObject(t *testing.T) {
	json := `{"empty_project": {}}`
	r := strings.NewReader(json)
	root, err := ParseJSON(r)

	if err != nil {
		t.Errorf("ParseJSON should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}

	if root.Name != "empty_project" {
		t.Errorf("Root name should be 'empty_project', got %s", root.Name)
	}
}

func TestParseJSONNestedObject(t *testing.T) {
	json := `{"backend": {"src": {}, "tests": {}}}`
	r := strings.NewReader(json)
	root, err := ParseJSON(r)

	if err != nil {
		t.Errorf("ParseJSON should not error: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}
}

func TestParseJSONInvalidJSON(t *testing.T) {
	json := `{"incomplete": {`
	r := strings.NewReader(json)
	_, err := ParseJSON(r)

	if err == nil {
		t.Error("ParseJSON should error on invalid JSON")
	}
}

func TestParseJSONWithNull(t *testing.T) {
	json := `{"project": {"file.txt": null}}`
	r := strings.NewReader(json)
	root, err := ParseJSON(r)

	if err != nil {
		t.Errorf("ParseJSON should handle null values: %v", err)
	}

	if root == nil {
		t.Fatal("Root should not be nil")
	}
}
