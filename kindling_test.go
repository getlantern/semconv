package semconv

import (
	"testing"

	"go.opentelemetry.io/otel/attribute"
)

func TestKindlingTransportAttr(t *testing.T) {
	for key, want := range map[string]string{
		string(KindlingOutcomeKey):   "kindling.smart.outcome",
		string(KindlingStatusKey):    "kindling.smart.status",
		string(KindlingElapsedMsKey): "kindling.smart.elapsed_ms",
	} {
		if got := KindlingTransportAttr("smart", attribute.Key(key)); string(got) != want {
			t.Errorf("KindlingTransportAttr(smart, %s) = %s, want %s", key, got, want)
		}
	}
}
