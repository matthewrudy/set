package set

import "testing"

func Test_Add(t *testing.T) {
	set := New()
	set.Add("abc")
	set.Add("abc")
	set.Add("def")

	if set.Size() != 2 {
		t.Error("size should be 2")
	}

}

func Test_Remove(t *testing.T) {
	set := New()
	set.Add("abc")
	set.Add("abc")
	set.Add("def")
	set.Remove("abc")

	if set.Size() != 1 {
		t.Error("size should be 1")
	}

	if set.Has("abc") == true {
		t.Error("abc should have been removed")
	}
}

func Test_Clear(t *testing.T) {
	set := New()
	set.Add("abc")
	set.Add("def")

	if set.Size() != 2 {
		t.Error("size should be 2")
	}

	set.Clear()

	if set.Size() != 0 {
		t.Error("size should be 0")
	}
}

func Test_Has(t *testing.T) {
	set := New()
	set.Add("abc")
	set.Add("abc")
	set.Add("def")

	if set.Has("abc") != true {
		t.Error("should have abc")
	}

	if set.Has("ghi") != false {
		t.Error("should not have ghi")
	}
}
