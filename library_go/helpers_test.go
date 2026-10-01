package library

import (
	"reflect"
	"testing"
)

func fmtInt(p *int) any {
	if p == nil {
		return "nil"
	}
	return *p
}

func fmtStr(p *string) any {
	if p == nil {
		return "nil"
	}
	return *p
}

func eqInt(t *testing.T, got *int, want int) {
	t.Helper()
	if got == nil || *got != want {
		t.Fatalf("got %v, want %d", fmtInt(got), want)
	}
}

func nilInt(t *testing.T, got *int) {
	t.Helper()
	if got != nil {
		t.Fatalf("got %d, want nil", *got)
	}
}

func eqStr(t *testing.T, got *string, want string) {
	t.Helper()
	if got == nil || *got != want {
		t.Fatalf("got %v, want %q", fmtStr(got), want)
	}
}

func nilStr(t *testing.T, got *string) {
	t.Helper()
	if got != nil {
		t.Fatalf("got %q, want nil", *got)
	}
}

func eqBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func eqList(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}
