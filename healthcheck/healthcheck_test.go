package healthcheck

import (
	"testing"
)

func TestReturnHealthCheck(t *testing.T) {
	hc := &HealthCheck{
		Status:  "ok",
		Version: "1.0.1",
	}

	status := hc.GetHealthCheck().Status
	if status != "ok" {
		t.Errorf("The Status value is incorrect, got: %s, want: %s.", status, "ok")
	}

	version := hc.GetHealthCheck().Version
	if version != "1.0.1" {
		t.Errorf("The Version value is incorrect, got: %s, want: %s.", version, "1.0.1")
	}
}
