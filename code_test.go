package country

import "testing"

func TestCode(t *testing.T) {
	if got := SE.Country(); got != "Sweden" {
		t.Error("wrong country", got)
	}

	if got := SE.Alpha3(); got != "SWE" {
		t.Error("wrong alpha3", got)
	}

	if got := SE.Numeric(); got != 752 {
		t.Error("wrong numeric", got)
	}
}
