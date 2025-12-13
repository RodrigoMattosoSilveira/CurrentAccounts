package utilities

import (
	"sort"
	"testing"
)

func TestNewTemplateData(t *testing.T) {
	templateData := NewTemplateData()
	if len(templateData.Data) != 2 {
		t.Errorf("Expected 1 file, got %d", len(templateData.Data))
	}
	keys := []string{}
	for key := range templateData.Data {
		keys = append(keys, key)
	}

	if !in("Tenant", keys) {
		t.Errorf("Key 'Tenant' not found in template data.")
	}
	if !in("Host", keys) {
		t.Errorf("Key 'Host' not found in template data.")
	}
}
func TestTemplateDataAddInt(t *testing.T) {
	templateData := NewTemplateData()
	templateData.AddData("UserID", 12345)
	if len(templateData.Data) != 3 {
		t.Errorf("Expected 3 file, got %d", len(templateData.Data))
	}
	keys := []string{}
	for key := range templateData.Data {
		keys = append(keys, key)
	}

	if !in("Tenant", keys) {
		t.Errorf("Key 'Tenant' not found in template data.")
	}
	if !in("Host", keys) {
		t.Errorf("Key 'Host' not found in template data.")
	}
	if !in("UserID", keys) {
		t.Errorf("Key 'UserID' not found in template data.")
	}

	if templateData.Data["UserID"] != 12345 {
		t.Errorf("Value for 'UserID' incorrect, got: %v, want: %v.", templateData.Data["UserID"], 12345)
	}
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	 // Index value: [0, LEN (STR_ARRAY)]
	 if (index <len (str_array) && str_array [index] == target) {// Need to pay attention to the judgment of this, first judge the conditions of && on the left side, if you do not satisfy, the end is ended here, no further judgment again
		return true
	}
	return false
}