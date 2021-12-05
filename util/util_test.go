package util

import (
	"reflect"
	"testing"
)

func TestSplitOnBlankSimple(t *testing.T) {
	items := SplitOnWhiteSpace(" 42 13")
	if !reflect.DeepEqual(items, []string{"42", "13"}) {
		t.Error("wrong items")
	}
}
