// test for chains.go
package db

import "testing"

// TestFindAllChains - tests FindAllChains
func TestFindAllChains(t *testing.T) {
	db, err := New(config)
	if err != nil {
		t.Errorf("error getting db: %+v", err)
	}

	// test
	chains, err := db.FindAllChains()
	if err != nil {
		t.Fatalf("unable to find chains for tokens: %+v", err)
	}
	// verify
	if len(chains) < 1 {
		t.Fatalf("expected > 1 chains, found %d", len(chains))
	}
}
