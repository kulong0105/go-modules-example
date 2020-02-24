package show

import "testing"

func TestInfo(t *testing.T) {
	name := "Allen"
	want := "Name: Allen"

	if got := Name(name); want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}