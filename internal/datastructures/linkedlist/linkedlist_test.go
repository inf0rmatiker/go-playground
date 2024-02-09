package linkedlist

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	testList := New()
	testList.Append(5)
	testList.Append(9)
	testList.Append(7)
	if testList.size != 3 {
		t.Fail()
	}

	// Eh, ok
	ok, item := testList.Get(1)
	if !ok || item != 9 {
		t.Fail()
	}

	// Nasty
	if ok, item = testList.Get(0); !ok || item != 5 {
		t.Fail()
	}
}
