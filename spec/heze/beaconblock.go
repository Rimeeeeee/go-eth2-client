package heze

import (
	"fmt"

	"github.com/ethpandaops/go-eth2-client/spec/phase0"
	"github.com/goccy/go-yaml"
)

// BeaconBlock represents a Heze beacon block.
type BeaconBlock struct {
	Slot          phase0.Slot
	ProposerIndex phase0.ValidatorIndex
	ParentRoot    phase0.Root `ssz-size:"32"`
	StateRoot     phase0.Root `ssz-size:"32"`
	Body          *BeaconBlockBody
}

// String returns a string version of the structure.
func (b *BeaconBlock) String() string {
	data, err := yaml.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
