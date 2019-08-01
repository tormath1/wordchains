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
		res := isWordInTheUniverse(e.w, &e.u)
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
		preProcessData(e.w, &e.list, &r)
		for i, el := range r {
			if el != e.res[i] {
				t.Errorf("should have %s, got %s", e.res[i], el)
			}
		}
	}
}

func TestIsNeighbor(t *testing.T) {
	d := []struct {
		w    string
		word string
		res  bool
	}{
		{
			w:    "dog",
			word: "dig",
			res:  true,
		},
		{
			w:    "dog",
			word: "cat",
			res:  false,
		},
		{
			w:    "dig",
			word: "dot",
			res:  false,
		},
	}
	for i, e := range d {
		res := isNeighbor(e.w, e.word)
		if res != e.res {
			t.Errorf("%d# should have %t, got %t", i, e.res, res)
		}
	}
}

func TestGetNeighbors(t *testing.T) {
	d := []struct {
		w   string
		u   []string
		exp []string
	}{
		{
			w:   "dog",
			u:   []string{"dig", "dot", "dat"},
			exp: []string{"dig", "dot"},
		},
		{
			w:   "dog",
			u:   []string{"zzz", "bbb", "ccc"},
			exp: []string{},
		},
	}
	for n, e := range d {
		var res []string
		getNeighbors(e.w, &e.u, &res)
		for i, el := range res {
			if el != e.exp[i] {
				t.Errorf("%d# should have %s, got %s", n, e.exp[i], el)
			}
		}
	}
}
