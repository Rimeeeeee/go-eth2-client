package spec

import (
	"testing"

	"github.com/ethpandaops/go-eth2-client/spec/heze"
)

func TestVersionedSignedBeaconBlockIsEmptyWithHeze(t *testing.T) {
	block := &VersionedSignedBeaconBlock{
		Version: DataVersionHeze,
		Heze:    &heze.SignedBeaconBlock{},
	}

	if block.IsEmpty() {
		t.Fatal("expected heze signed beacon block to be non-empty")
	}
}
