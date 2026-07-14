package bounce_test

import (
	"testing"

	"github.com/jackli/frank/internal/store"
)

func TestBucketDEdgesExtendAcceptanceBouncesOnlyWithStaleChoiceSet(t *testing.T) {
	acceptanceBounces := []string{
		"form-validation",
		"lineage",
		"observe-predicate",
		"declared-vs-observed",
	}
	for _, edge := range acceptanceBounces {
		if !store.AcceptanceBounceEdge(edge) {
			t.Errorf("AcceptanceBounceEdge(%q) = false, want true", edge)
		}
		if !store.BucketDFailingEdge(edge) {
			t.Errorf("BucketDFailingEdge(%q) = false, want true", edge)
		}
	}

	if store.AcceptanceBounceEdge("stale_choice_set") {
		t.Fatal("stale_choice_set entered the FSM acceptance-bounce subset")
	}
	if !store.BucketDFailingEdge("stale_choice_set") {
		t.Fatal("stale_choice_set absent from Bucket D")
	}

	for _, edge := range []string{"", "egress", "stale_schema"} {
		if store.AcceptanceBounceEdge(edge) || store.BucketDFailingEdge(edge) {
			t.Errorf("unclassified edge %q entered an acceptance-bounce view", edge)
		}
	}
}
