package mainkata

import "testing"

func TestBalik(t *testing.T) {
	cases := []struct {
		awal, akhir string
	}{
		{"aHello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, coba := range cases {
		dapetnya := Balik(coba.awal)
		if dapetnya != coba.akhir {
			t.Errorf("Balik(%q) == %q, akhir %q", coba.awal, dapetnya, coba.akhir)
		}
	}
}
