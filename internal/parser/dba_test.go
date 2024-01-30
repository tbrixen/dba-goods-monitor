package parser

import "testing"

func TestIsAct(t *testing.T) {
	dba := NewDba()
	dba.IsActive("foo")
}
