package main

import "testing"

func TestIsWordInTheUniverse(t *testing.T) {
	d := []struct {
		u []string
		w string
		r bool
	}{
		{
			u: []string{"a", "b", "c"},
			w: "d",
			r: false,
		},
		{
			u: []string{"a", "b", "c"},
			w: "a",
			r: true,
		},
	}

	for _, e := range d {
		res := IsWordInTheUniverse(e.w, &e.u)
		if res != e.r {
			t.Errorf("should have %t, got %t", e.r, res)
		}
	}
}

func TestPreProcessData(t *testing.T) {
	d := []struct {
		w    string
		list []string
		res  []string
	}{
		{
			w:    "a",
			list: []string{"a", "ab", "ac", "b"},
			res:  []string{"a", "b"},
		},
		{
			w:    "ar",
			list: []string{"a", "ab", "ac", "b"},
			res:  []string{"ab", "ac"},
		},
	}
	for _, e := range d {
		var r []string
		PreProcessData(e.w, &e.list, &r)
		for i, el := range r {
			if el != e.res[i] {
				t.Errorf("should have %s, got %s", e.res[i], el)
			}
		}
	}
}
