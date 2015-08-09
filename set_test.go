// Copyright 2015 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set_test

import (
	"testing"

	"github.com/xtgo/set/internal/sliceset"
	"github.com/xtgo/set/internal/testdata"
)

func TestUniq(t *testing.T) {
	for _, tt := range testdata.UniqTests {
		s := sliceset.Set(tt.In).Copy()
		s = s.Uniq()

		if !testdata.IsEqual(s, tt.Out) {
			t.Errorf("Uniq(%v) = %v, want %v", tt.In, s, tt.Out)
		}
	}
}

func TestUnion(t *testing.T)   { runMut(t, "Union") }
func TestInter(t *testing.T)   { runMut(t, "Inter") }
func TestDiff(t *testing.T)    { runMut(t, "Diff") }
func TestSymDiff(t *testing.T) { runMut(t, "SymDiff") }
func TestIsSub(t *testing.T)   { runBool(t, "IsSub") }
func TestIsSuper(t *testing.T) { runBool(t, "IsSuper") }
func TestIsInter(t *testing.T) { runBool(t, "IsInter") }
func TestIsEqual(t *testing.T) { runBool(t, "IsEqual") }

const format = "%s(%v, %v) = %v, want %v"

type (
	mutOp  func(a, b sliceset.Set) sliceset.Set
	boolOp func(a, b sliceset.Set) bool
)

func runMut(t *testing.T, name string) {
	var op mutOp
	testdata.ConvMethod(&op, sliceset.Set(nil), name)

	for _, tt := range testdata.BinTests {
		a := sliceset.Set(tt.A).Copy()
		c := op(a, tt.B)
		want := tt.SelSlice(name)

		if !testdata.IsEqual(c, want) {
			t.Errorf(format, name, tt.A, tt.B, c, want)
		}
	}
}

func runBool(t *testing.T, name string) {
	var op boolOp
	testdata.ConvMethod(&op, sliceset.Set(nil), name)

	for _, tt := range testdata.BinTests {
		ok := op(tt.A, tt.B)
		want := tt.SelBool(name)

		if ok != want {
			t.Errorf(format, name, tt.A, tt.B, ok, want)
		}
	}
}
