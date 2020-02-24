package show

import "testing"

func TestInfo(t *testing.T) {
	name := "Allen"
	want := "My Name: Allen."

	if got := NameV3(name); want != got {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
