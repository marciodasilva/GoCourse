package greeting

import "testing"

func Test(t *testing.T) {
	var prefixMap = map[string]string{
		"Marcio": "Mr. ",
		"Bob":    "Mr. ",
		"Joe":    "Dr. ",
		"Mary":   "Mrs. ",
		"John":   "Dr. ",
	}
	// ret := GetPrefix(para)
	// want := "Dr. "
	for k, _ := range prefixMap {
		ret := GetPrefix(k)
		if ret != "Mr. " && ret != "Mrs. " {
			t.Errorf("GetPrefix(%q) == %q", k, ret)
		}
	}
}
