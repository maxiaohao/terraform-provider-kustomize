package kustomize

import (
	"testing"
)

func TestLastAppliedConfig(t *testing.T) {
	srcJSON := "{\"apiVersion\": \"v1\", \"kind\": \"Namespace\", \"metadata\": {\"name\": \"test-unit\"}}"
	u, err := parseJSON(srcJSON)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	setLastAppliedConfig(u, srcJSON)

	annotations := u.GetAnnotations()
	count := len(annotations)
	if count != 1 {
		t.Errorf("TestLastAppliedConfig: incorect number of annotations, got: %d, want: %d.", count, 1)
	}

	lac := getLastAppliedConfig(u)
	if lac != srcJSON {
		t.Errorf("TestLastAppliedConfig: incorect annotation value, got: %s, want: %s.", srcJSON, lac)
	}
}

func TestGetPatch(t *testing.T) {
	srcJSON := "{\"apiVersion\": \"v1\", \"kind\": \"Namespace\", \"metadata\": {\"name\": \"test-unit\"}}"

	o, _ := parseJSON(srcJSON)
	m, _ := parseJSON(srcJSON)
	c, _ := parseJSON(srcJSON)

	original, _ := o.MarshalJSON()
	modified, _ := m.MarshalJSON()
	current, _ := c.MarshalJSON()

	_, err := getPatch(original, modified, current)
	if err != nil {
		t.Errorf("TestGetPatch: %s", err)
	}
}

func TestSimplifyCurrent(t *testing.T) {
	fullCurrent := []byte("{\"apiVersion\": \"v1\", \"kind\": \"Namespace\", \"metadata\": {\"name\": \"test-unit-old\", \"namespace\": \"test-unit\", \"foo\": \"bar\"}}")
	targeted := []byte("{\"apiVersion\": \"v1\", \"kind\": \"Namespace\", \"metadata\": {\"name\": \"test-unit-new\", \"baz\": \"qux\"}}")

	result, err := simplifyCurrent(fullCurrent, targeted)
	t.Logf("result: %s", result)
	if err != nil {
		t.Errorf("TestSimplifyCurrent: %s", err)
	}

	expected := "{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"name\":\"test-unit-old\"}}"
	if expected != string(result) {
		t.Errorf("expected: %s, got: %s", expected, result)
	}
}
