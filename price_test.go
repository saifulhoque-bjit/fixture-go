package price

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTotalCountsEveryItem(t *testing.T) {
	if diff := cmp.Diff(900, Total(250, 4, 10)); diff != "" {
		t.Fatalf("Total(250, 4, 10) mismatch (-want +got): %s", diff)
	}
}

func TestTotalWithNoDiscount(t *testing.T) {
	if got := Total(199, 3, 0); got != 597 {
		t.Fatalf("Total(199, 3, 0) = %d, want 597", got)
	}
}
