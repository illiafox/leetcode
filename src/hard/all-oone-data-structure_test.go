package hard

import (
	"testing"
)

func oneOf(t *testing.T, got string, expected ...string) {
	t.Helper()
	for _, e := range expected {
		if got == e {
			return
		}
	}
	t.Fatalf("got %q; want one of %v", got, expected)
}

func TestAllOne_Empty(t *testing.T) {
	ao := NewAllOne()
	if maxKey := ao.GetMaxKey(); maxKey != "" {
		t.Fatalf("GetMaxKey on empty = %q; want \"\"", maxKey)
	}
	if minKey := ao.GetMinKey(); minKey != "" {
		t.Fatalf("GetMinKey on empty = %q; want \"\"", minKey)
	}
}

func TestAllOne_ExampleFromPrompt(t *testing.T) {
	allOne := NewAllOne()
	allOne.Inc("hello")
	allOne.Inc("hello")

	if got := allOne.GetMaxKey(); got != "hello" {
		t.Fatalf("GetMaxKey() = %q; want %q", got, "hello")
	}
	if got := allOne.GetMinKey(); got != "hello" {
		t.Fatalf("GetMinKey() = %q; want %q", got, "hello")
	}

	allOne.Inc("leet")

	if got := allOne.GetMaxKey(); got != "hello" {
		t.Fatalf("GetMaxKey() = %q; want %q", got, "hello")
	}
	if got := allOne.GetMinKey(); got != "leet" {
		t.Fatalf("GetMinKey() = %q; want %q", got, "leet")
	}
}

func TestAllOne_Basics(t *testing.T) {
	ao := NewAllOne()

	// inc(a), inc(b)
	ao.Inc("a")
	ao.Inc("b")
	oneOf(t, ao.GetMaxKey(), "a", "b")
	oneOf(t, ao.GetMinKey(), "a", "b")

	// inc(a) again => a:2, b:1
	ao.Inc("a")
	if got := ao.GetMaxKey(); got != "a" {
		t.Fatalf("max after Inc(a) = %q; want a", got)
	}
	if got := ao.GetMinKey(); got != "b" {
		t.Fatalf("min after Inc(a) = %q; want b", got)
	}

	// dec(a) => a:1, b:1 (tie)
	ao.Dec("a")
	oneOf(t, ao.GetMaxKey(), "a", "b")
	oneOf(t, ao.GetMinKey(), "a", "b")

	// dec(b) => remove b; only a:1 remains
	ao.Dec("b")
	if got := ao.GetMaxKey(); got != "a" {
		t.Fatalf("max after Dec(b) = %q; want a", got)
	}
	if got := ao.GetMinKey(); got != "a" {
		t.Fatalf("min after Dec(b) = %q; want a", got)
	}

	// dec(a) => empty
	ao.Dec("a")
	if got := ao.GetMaxKey(); got != "" {
		t.Fatalf("max after removing all = %q; want \"\"", got)
	}
	if got := ao.GetMinKey(); got != "" {
		t.Fatalf("min after removing all = %q; want \"\"", got)
	}
}

func TestAllOne_TiesAndRemovals(t *testing.T) {
	ao := NewAllOne()

	// a:2, b:2, c:1
	ao.Inc("a")
	ao.Inc("a")
	ao.Inc("b")
	ao.Inc("b")
	ao.Inc("c")

	oneOf(t, ao.GetMaxKey(), "a", "b")
	if got := ao.GetMinKey(); got != "c" {
		t.Fatalf("min = %q; want c", got)
	}

	// dec(a) => a:1, b:2, c:1  (two mins now)
	ao.Dec("a")
	if got := ao.GetMaxKey(); got != "b" {
		t.Fatalf("max after Dec(a) = %q; want b", got)
	}
	oneOf(t, ao.GetMinKey(), "a", "c")

	// dec(c) => c removed; mins collapse to a:1
	ao.Dec("c")
	if got := ao.GetMinKey(); got != "a" {
		t.Fatalf("min after Dec(c) = %q; want a", got)
	}

	// dec(b), dec(b) => b removed; only a:1 left
	ao.Dec("b")
	ao.Dec("b")
	if got := ao.GetMaxKey(); got != "a" {
		t.Fatalf("max after removing b = %q; want a", got)
	}
	if got := ao.GetMinKey(); got != "a" {
		t.Fatalf("min after removing b = %q; want a", got)
	}
}

func TestAllOne_LongSequence(t *testing.T) {
	type step struct {
		op     string // "inc" | "dec" | "max" | "min"
		arg    string // key for inc/dec
		expMax []string
		expMin []string
	}

	steps := []step{
		{"inc", "x", nil, nil},                      // x:1
		{"inc", "y", nil, nil},                      // x:1, y:1
		{"inc", "y", nil, nil},                      // x:1, y:2
		{"inc", "z", nil, nil},                      // x:1, y:2, z:1
		{"max", "", []string{"y"}, []string{}},      // max=y
		{"min", "", []string{}, []string{"x", "z"}}, // min in {x,z}
		{"inc", "z", nil, nil},                      // x:1, y:2, z:2
		{"max", "", []string{"y", "z"}, []string{}},
		{"dec", "y", nil, nil}, // x:1, y:1, z:2
		{"min", "", []string{}, []string{"x", "y"}},
		{"dec", "x", nil, nil}, // x removed, y:1, z:2
		{"min", "", []string{}, []string{"y"}},
		{"inc", "y", nil, nil}, // y:2, z:2
		{"inc", "y", nil, nil}, // y:3, z:2
		{"max", "", []string{"y"}, []string{}},
		{"dec", "y", nil, nil}, // y:2, z:2
		{"dec", "z", nil, nil}, // y:2, z:1
		{"min", "", []string{}, []string{"z"}},
		{"dec", "z", nil, nil}, // z removed, y:2
		{"max", "", []string{"y"}, []string{}},
		{"min", "", []string{}, []string{"y"}},
	}

	ao := NewAllOne()

	for i, s := range steps {
		switch s.op {
		case "inc":
			ao.Inc(s.arg)
		case "dec":
			ao.Dec(s.arg)
		case "max":
			got := ao.GetMaxKey()
			if len(s.expMax) == 0 {
				if got != "" {
					t.Fatalf("step %d: expected empty max, got %q", i, got)
				}
			} else {
				oneOf(t, got, s.expMax...)
			}
		case "min":
			got := ao.GetMinKey()
			if len(s.expMin) == 0 {
				if got != "" {
					t.Fatalf("step %d: expected empty min, got %q", i, got)
				}
			} else {
				oneOf(t, got, s.expMin...)
			}
		default:
			t.Fatalf("step %d: unknown op %q", i, s.op)
		}
	}
}

func TestAllOne_ManyOpsDeterministic(t *testing.T) {
	ao := NewAllOne()
	keys := []string{"a", "b", "c", "d", "e"}

	// Build counts: a:1, b:2, c:3, d:4, e:5
	for i, k := range keys {
		for j := 0; j <= i; j++ {
			ao.Inc(k)
		}
	}
	if got := ao.GetMaxKey(); got != "e" {
		t.Fatalf("max after setup = %q; want e", got)
	}
	if got := ao.GetMinKey(); got != "a" {
		t.Fatalf("min after setup = %q; want a", got)
	}

	// Bring all down to 1, removing as we go, checking edges.
	for i := len(keys) - 1; i >= 0; i-- {
		k := keys[i]
		for c := i; c > 0; c-- { // reduce k down to 1
			ao.Dec(k)
		}
	}

	// Now all keys should be at count 1
	oneOf(t, ao.GetMaxKey(), keys...)
	oneOf(t, ao.GetMinKey(), keys...)

	// Remove all
	for _, k := range keys {
		ao.Dec(k)
	}
	if ao.GetMaxKey() != "" || ao.GetMinKey() != "" {
		t.Fatalf("expected empty after removing all; got max=%q min=%q", ao.GetMaxKey(), ao.GetMinKey())
	}
}
