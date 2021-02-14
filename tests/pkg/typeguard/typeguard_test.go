package typeguard_test

import (
	"testing"

	"github.com/syncship/moby-dick/pkg/typeguard"
)

func TestToArrStringOnSuccess(t *testing.T) {
	output := typeguard.Output{Value: "a,b,c"}
	got, err := output.ToArrString()
	want := []string{"a", "b", "c"}

	if err != nil {
		t.Errorf("Error should be nil, got: %s", err)
	}

	if (got == nil) != (want == nil) {
		t.Errorf("Values should be equal, got: %s, %s", got, want)
	}

	if len(got) != len(want) {
		t.Errorf("Values should have same length, got: %v, %v", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Items in values should be equal, got: %s, %s", got[i], want[i])
		}
	}
}
