package flexdds

import "testing"

func TestFlexDDS(t *testing.T) {
	flexdds, err := Open(Config{
		Host:     "10.163.100.3",
		Slot:     2,
		SysClock: 1e9,
	})
	if err != nil {
		t.Fatalf("failed to open flexdds connection: %s", err)
	}

	if err := flexdds.Singletone(0, 0.0, 60e6); err != nil {
		t.Errorf("failed to configure channel 0 as singletone: %s", err)
	}

	if err = flexdds.Close(); err != nil {
		t.Errorf("failed to close flexdds connection: %s", err)
	}
}
