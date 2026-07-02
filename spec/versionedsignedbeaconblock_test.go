package spec

import (
	"testing"

	"github.com/ethpandaops/go-eth2-client/spec/gloas"
)

func TestVersionedSignedBeaconBlockIsEmptyWithHeze(t *testing.T) {
	block := &VersionedSignedBeaconBlock{
		Version: DataVersionHeze,
		Heze:    &gloas.SignedBeaconBlock{},
	}

	if block.IsEmpty() {
		t.Fatal("expected heze signed beacon block to be non-empty")
	}
}
