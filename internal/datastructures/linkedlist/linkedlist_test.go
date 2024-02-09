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

func TestLinkedList_Insert(t *testing.T) {
	testList := New()
	testList.Append(5)
	testList.Append(9)
	testList.Append(7)

	var actual int

	// Insert at tail
	ok := testList.Insert(99, 3)
	if !ok {
		t.Fail()
	}
	if ok, actual = testList.Get(3); !ok || actual != 99 {
		t.Fail()
	}

	// Insert at head
	ok = testList.Insert(98, 0)
	if !ok {
		t.Fail()
	}
	if ok, actual = testList.Get(0); !ok || actual != 98 {
		t.Fail()
	}

	// Insert in middle
	ok = testList.Insert(97, 1)
	if !ok {
		t.Fail()
	}
	if ok, actual = testList.Get(1); !ok || actual != 97 {
		t.Fail()
	}
}

func TestLinkedList_RemoveHead(t *testing.T) {
	testList := New()
	testList.Append(5)
	testList.Append(9)
	testList.Append(7)

	var actual int

	ok := testList.Remove(0)
	if !ok {
		t.Fail()
	}
	if testList.size != 2 {
		t.Fail()
	}
	if ok, actual = testList.Get(0); !ok || actual != 9 {
		t.Fail()
	}
}

func TestLinkedList_RemoveTail(t *testing.T) {
	testList := New()
	testList.Append(5)
	testList.Append(9)
	testList.Append(7)

	var actual int

	ok := testList.Remove(testList.size - 1)
	if !ok {
		t.Fail()
	}
	if testList.size != 2 {
		t.Fail()
	}
	if ok, actual = testList.Get(testList.size - 1); !ok || actual != 9 {
		t.Fail()
	}
}

func TestLinkedList_RemoveMiddle(t *testing.T) {
	testList := New()
	testList.Append(5)
	testList.Append(9)
	testList.Append(7)

	var actual int

	ok := testList.Remove(1)
	if !ok {
		t.Fail()
	}
	if testList.size != 2 {
		t.Fail()
	}
	if ok, actual = testList.Get(1); !ok || actual != 7 {
		t.Fail()
	}
}
