package country

import (
	"fmt"
	"testing"
)

func Example_codeMethods() {
	fmt.Println(SE.String())
	fmt.Println(SE.Alpha2())
	fmt.Println(SE.Alpha3())
	fmt.Println(SE.Numeric())
	fmt.Println(SE.Country())
	// ouput:
	// SE SWE 752 Sweden
	// SE
	// SWE
	// 752
	// Sweden

}

func TestCode(t *testing.T) {
	if got := SE.Country(); got != "Sweden" {
		t.Error("wrong country", got)
	}

	if got := SE.Alpha2(); got != "SE" {
		t.Error("wrong alpha2", got)
	}

	if got := SE.Alpha3(); got != "SWE" {
		t.Error("wrong alpha3", got)
	}

	if got := SE.Numeric(); got != 752 {
		t.Error("wrong numeric", got)
	}
}
