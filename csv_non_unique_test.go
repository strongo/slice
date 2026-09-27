package slice

import (
	"testing"
)

func TestCommaSeparatedValuesList_Add(t *testing.T) {
	var csvl CommaSeparatedValuesList
	if csvl = csvl.Add("v1"); csvl != "v1" {
		t.Error("unexpected: " + csvl)
	}
	if csvl = csvl.Add("v3"); csvl != "v1,v3" {
		t.Error("unexpected")
	}
	if csvl = csvl.Add("v2"); csvl != "v1,v3,v2" {
		t.Error("unexpected")
	}
}
func TestCommaSeparatedValuesList_Remove(t *testing.T) {
	var csvl CommaSeparatedValuesList
	csvl = CommaSeparatedValuesList("")
	if csvl = csvl.Remove("v1"); csvl != "" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedValuesList("v1")
	if csvl = csvl.Remove("v1"); csvl != "" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedValuesList("v2")
	if csvl = csvl.Remove("v1"); csvl != "v2" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v1"); csvl != "v3,v2,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v2"); csvl != "v1,v3,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v3"); csvl != "v1,v2,v4" {
		t.Error("unexpected: " + csvl)
	}
	csvl = CommaSeparatedValuesList("v1,v3,v2,v4")
	if csvl = csvl.Remove("v4"); csvl != "v1,v3,v2" {
		t.Error("unexpected: " + csvl)
	}
}

func TestCommaSeparatedValuesList_Contains(t *testing.T) {
	if CommaSeparatedValuesList("").Contains("v1") {
		t.Error("unexpected true")
	}
	if !CommaSeparatedValuesList("v1").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedValuesList("v1,v2").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedValuesList("v1,v2").Contains("v2") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedValuesList("v1,v2,v3").Contains("v1") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedValuesList("v1,v2,v3").Contains("v2") {
		t.Error("unexpected false")
	}
	if !CommaSeparatedValuesList("v1,v2,v3").Contains("v3") {
		t.Error("unexpected false")
	}
}

func TestCommaSeparatedValuesList_Count(t *testing.T) {
	if count := CommaSeparatedValuesList("").Count(); count != 0 {
		t.Errorf("expected 0, got %d", count)
	}
	if count := CommaSeparatedValuesList("v1,v2,v3").Count(); count != 3 {
		t.Errorf("expected 3, got %d", count)
	}
}

func TestCommaSeparatedValuesList_Set(t *testing.T) {
	list := CommaSeparatedValuesList("v1,v2,v3")
	if updated := list.Set(1, "updated"); updated != "v1,updated,v3" {
		t.Errorf("expected v1,updated,v3, got %s", updated)
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic on out of range Set")
		}
	}()
	list.Set(5, "out_of_range")
}

func TestCommaSeparatedValuesList_Strings(t *testing.T) {
	if s := CommaSeparatedValuesList("").Strings(); len(s) != 0 {
		t.Errorf("expected empty slice, got %v", s)
	}
	if s := CommaSeparatedValuesList("v1,v2").Strings(); len(s) != 2 || s[0] != "v1" || s[1] != "v2" {
		t.Errorf("expected [v1 v2], got %v", s)
	}
}

func TestCommaSeparatedValuesList_String(t *testing.T) {
	if s := CommaSeparatedValuesList("v1,v2").String(); s != "v1,v2" {
		t.Errorf("expected v1,v2, got %s", s)
	}
}

func TestCommaSeparatedValuesList_Add_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic when adding value with comma")
		}
	}()
	CommaSeparatedValuesList("v1").Add("v2,v3")
}

