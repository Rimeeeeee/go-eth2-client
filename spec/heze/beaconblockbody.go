package heze

import (
	"fmt"

	"github.com/ethpandaops/go-eth2-client/spec/altair"
	"github.com/ethpandaops/go-eth2-client/spec/capella"
	"github.com/ethpandaops/go-eth2-client/spec/deneb"
	"github.com/ethpandaops/go-eth2-client/spec/electra"
	"github.com/ethpandaops/go-eth2-client/spec/gloas"
	"github.com/ethpandaops/go-eth2-client/spec/phase0"
	"github.com/goccy/go-yaml"
)

// BeaconBlockBody represents a Heze beacon block body.
type BeaconBlockBody struct {
	RANDAOReveal          phase0.BLSSignature `ssz-size:"96"`
	ETH1Data              *phase0.ETH1Data
	Graffiti              [32]byte                      `ssz-size:"32"`
	ProposerSlashings     []*phase0.ProposerSlashing    `dynssz-max:"MAX_PROPOSER_SLASHINGS"         ssz-max:"16"`
	AttesterSlashings     []*electra.AttesterSlashing   `dynssz-max:"MAX_ATTESTER_SLASHINGS_ELECTRA" ssz-max:"1"`
	Attestations          []*electra.Attestation        `dynssz-max:"MAX_ATTESTATIONS_ELECTRA"       ssz-max:"8"`
	Deposits              []*phase0.Deposit             `dynssz-max:"MAX_DEPOSITS"                   ssz-max:"16"`
	VoluntaryExits        []*phase0.SignedVoluntaryExit `dynssz-max:"MAX_VOLUNTARY_EXITS"            ssz-max:"16"`
	SyncAggregate         *altair.SyncAggregate
	ExecutionPayload      *gloas.ExecutionPayload
	BLSToExecutionChanges []*capella.SignedBLSToExecutionChange `dynssz-max:"MAX_BLS_TO_EXECUTION_CHANGES"   ssz-max:"16"`
	BlobKZGCommitments    []deneb.KZGCommitment                 `dynssz-max:"MAX_BLOB_COMMITMENTS_PER_BLOCK" ssz-max:"4096" ssz-size:"?,48"`
	ExecutionRequests     *electra.ExecutionRequests
}

// String returns a string version of the structure.
func (b *BeaconBlockBody) String() string {
	data, err := yaml.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
