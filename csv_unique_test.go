package slice

import (
	"testing"
)

func TestCommaSeparatedUniqueValuesList_Add(t *testing.T) {
	var csvl CommaSeparatedUniqueValuesList
	if csvl = csvl.Add("v1", 0); csvl != "v1" {
		t.Error("unexpected: " + csvl)
	}
	if csvl = csvl.Add("v3", 0); csvl != "v1,v3" {
		t.Error("unexpected")
	}
	if csvl = csvl.Add("v2", 0); csvl != "v1,v3,v2" {
		t.Error("unexpected")
	}
}
func TestCommaSeparatedUniqueValuesList_Remove(t *testing.T) {
	var csvl CommaSeparatedUniqueValuesList
	csvl = CommaSeparatedUniqueValuesList("")
	if csvl = csvl.Remove("v1"); csvl != "" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1")
	if csvl = csvl.Remove("v1"); csvl != "" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v2")
	if csvl = csvl.Remove("v1"); csvl != "v2" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v1"); csvl != "v3,v2,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v2"); csvl != "v1,v3,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v3"); csvl != "v1,v2,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedUniqueValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v4"); csvl != "v1,v3,v2" {
		t.Error("unexpected: " + csvl)
	}
}

func TestCommaSeparatedUniqueValuesList_Contains(t *testing.T) {
	if CommaSeparatedUniqueValuesList("").Contains("v1") {
		t.Error("unexpected true")
	}
	if !CommaSeparatedUniqueValuesList("v1").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2").Contains("v2") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2,v3").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2,v3").Contains("v2") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedUniqueValuesList("v1,v2,v3").Contains("v3") {
		t.Error("unexpected false")
	}
}

func TestCommaSeparatedUniqueValuesList_Add_cases(t *testing.T) {
	// Add existing duplicate
	list := CommaSeparatedUniqueValuesList("v1,v2")
	if updated := list.Add("v1", 0); updated != "v1,v2" {
		t.Errorf("expected unchanged v1,v2, got %s", updated)
	}

	// Add with limit exceeded
	list = CommaSeparatedUniqueValuesList("v1,v2,v3")
	if updated := list.Add("v4", 2); updated != "v1,v2,v4" {
		t.Errorf("expected v1,v2,v4, got %s", updated)
	}

	// Add panic on comma
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when adding value with comma")
		}
	}()
	list.Add("a,b", 0)
}

func TestCommaSeparatedUniqueValuesList_Strings(t *testing.T) {
	if s := CommaSeparatedUniqueValuesList("").Strings(); len(s) != 0 {
		t.Errorf("expected empty slice, got %v", s)
	}
	if s := CommaSeparatedUniqueValuesList("v1,v2").Strings(); len(s) != 2 || s[0] != "v1" || s[1] != "v2" {
		t.Errorf("expected [v1 v2], got %v", s)
	}
}

func TestCommaSeparatedUniqueValuesList_String(t *testing.T) {
	if s := CommaSeparatedUniqueValuesList("v1,v2").String(); s != "v1,v2" {
		t.Errorf("expected v1,v2, got %s", s)
	}
}

