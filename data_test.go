package rrdb_test

import (
	"testing"

	"github.com/snowmerak/rrdb"
)

func TestGetColumnsOf(t *testing.T) {
	v := struct {
		A *int
		B *int
		C *int
		D *string
	}{}

	rs := rrdb.GetColumnsOf(v)
	if len(rs) != 4 {
		t.Error("GetColumnsOf failed")
	}

	a, b, c, d := 0, 0, 0, 0
	for _, v := range rs {
		if v != "A" && v != "B" && v != "C" && v != "D" {
			t.Error("GetColumnsOf failed")
		}
		switch v {
		case "A":
			a++
		case "B":
			b++
		case "C":
			c++
		case "D":
			d++
		}
	}

	if a != 1 || b != 1 || c != 1 || d != 1 {
		t.Error("GetColumnsOf failed")
	}
}

func TestGetColumnsOfDataByCache(t *testing.T) {
	v := struct {
		A *int
		B *int
		C *int
		D *string
	}{}

	rs := rrdb.GetColumnsOf(v)
	if len(rs) != 4 {
		t.Error("GetColumnsOf failed")
	}

	rs = rrdb.GetColumnsOf(v)
	if len(rs) != 4 {
		t.Error("GetColumnsOf failed")
	}
}

func TestGetNotNilColumnsOf(t *testing.T) {
	v := struct {
		A *int
		B *int
		C *int
		D *string
	}{}

	rs := rrdb.GetNotNilColumnsOf(v)
	t.Log(len(rs))
	if len(rs) != 0 {
		t.Error("GetNotNilColumnsOf failed")
	}

	a, b, c, d := 0, 0, 0, 0
	for _, v := range rs {
		if v != "A" && v != "B" && v != "C" && v != "D" {
			t.Error("GetNotNilColumnsOf failed")
		}
		switch v {
		case "A":
			a++
		case "B":
			b++
		case "C":
			c++
		case "D":
			d++
		}
	}

	if a != 0 || b != 0 || c != 0 || d != 0 {
		t.Error("GetNotNilColumnsOf failed")
	}
}
