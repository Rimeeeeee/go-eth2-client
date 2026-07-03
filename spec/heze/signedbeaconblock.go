package heze

import (
	"fmt"

	"github.com/ethpandaops/go-eth2-client/spec/phase0"
	"github.com/goccy/go-yaml"
)

// SignedBeaconBlock is a signed Heze beacon block.
type SignedBeaconBlock struct {
	Message   *BeaconBlock
	Signature phase0.BLSSignature `ssz-size:"96"`
}

// String returns a string version of the structure.
func (s *SignedBeaconBlock) String() string {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
