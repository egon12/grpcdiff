package bodygenerator

import (
	"reflect"
	"testing"
)

func TestSeperatedValueSource_Parse(t *testing.T) {
	s := &SeperatedValueSource{}
	s.Parse("1,2,3,4,5")

	want := []string{"1", "2", "3", "4", "5"}
	got := s.value
	if !reflect.DeepEqual(want, s.value) {
		t.Errorf("\nwant: %s\n got: %s", want, got)
	}
}
