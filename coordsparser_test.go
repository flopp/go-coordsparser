// Copyright 2016 Florian Pigorsch. All rights reserved.
//
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package coordsparser

import "testing"

func TestParseD(t *testing.T) {
	type testStruct struct {
		in       string
		lat, lng float64
		err      bool
	}

	var cases = []testStruct{
		{"bad_string", 0, 0, true},
		{"42.0, bad_string", 0, 0, true},
		{"bad_string, 23.123", 0, 0, true},
		{"142.0, 23.123", 0, 0, true},
		{"42.0, -223.123", 0, 0, true},

		{"42.0, 23.123", 42.0, 23.123, false},
		{"-42.0, 23.123", -42.0, 23.123, false},
		{"42.0, -23.123", 42.0, -23.123, false},
		{".0, -.123", 0.0, -0.123, false},

		// alternative separators
		{"42.0 23.123", 42.0, 23.123, false},
		{"42.0:23.123", 42.0, 23.123, false},
		{"42.0;23.123", 42.0, 23.123, false},
		{"42.0#23.123", 0, 0, true},

		// fancy spacing
		{"     42.0     23.123     ", 42.0, 23.123, false},
	}

	for _, c := range cases {
		lat, lng, err := ParseD(c.in)
		if c.err && (err == nil) {
			t.Errorf("ParseD(%q) did not raise expected error", c.in)
		} else if !c.err && (err != nil) {
			t.Errorf("ParseD(%q) raised unexpected error %q", c.in, err)
		} else if (lat != c.lat) || (lng != c.lng) {
			t.Errorf("ParseD(%q) == %f, %f; want %f, %f", c.in, lat, lng, c.lat, c.lng)
		}
	}
}

func TestParseHD(t *testing.T) {
	type testStruct struct {
		in       string
		lat, lng float64
		err      bool
	}

	var cases = []testStruct{
		{"N 40.76 W 73.984", 40.76, -73.984, false},
	}

	for _, c := range cases {
		lat, lng, err := ParseHD(c.in)
		if c.err && (err == nil) {
			t.Errorf("ParseHD(%q) did not raise expected error", c.in)
		} else if !c.err && (err != nil) {
			t.Errorf("ParseHD(%q) raised unexpected error %q", c.in, err)
		} else if (lat != c.lat) || (lng != c.lng) {
			t.Errorf("ParseHD(%q) == %f, %f; want %f, %f", c.in, lat, lng, c.lat, c.lng)
		}
	}
}

func TestParseHDM(t *testing.T) {
	type testStruct struct {
		in       string
		lat, lng float64
		err      bool
	}

	var cases = []testStruct{
		{"N 40 45.600 W 73 59.040", 40.76, -73.984, false},
	}

	for _, c := range cases {
		lat, lng, err := ParseHDM(c.in)
		if c.err && (err == nil) {
			t.Errorf("ParseHDM(%q) did not raise expected error", c.in)
		} else if !c.err && (err != nil) {
			t.Errorf("ParseHDM(%q) raised unexpected error %q", c.in, err)
		} else if (lat != c.lat) || (lng != c.lng) {
			t.Errorf("ParseHDM(%q) == %f, %f; want %f, %f", c.in, lat, lng, c.lat, c.lng)
		}
	}
}

func TestParseHDMS(t *testing.T) {
	type testStruct struct {
		in       string
		lat, lng float64
		err      bool
	}

	var cases = []testStruct{
		{"N 40 45 36.0 W 73 59 02.4", 40.76, -73.984, false},
	}

	for _, c := range cases {
		lat, lng, err := ParseHDMS(c.in)
		if c.err && (err == nil) {
			t.Errorf("ParseHDMS(%q) did not raise expected error", c.in)
		} else if !c.err && (err != nil) {
			t.Errorf("ParseHDMS(%q) raised unexpected error %q", c.in, err)
		} else if (lat != c.lat) || (lng != c.lng) {
			t.Errorf("ParseHDMS(%q) == %f, %f; want %f, %f", c.in, lat, lng, c.lat, c.lng)
		}
	}
}
