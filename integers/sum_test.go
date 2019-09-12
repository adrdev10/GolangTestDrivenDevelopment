package integers

import "testing"

func TestSum(t *testing.T) {
	got := Add(4, 3)
	want := 7

	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}
