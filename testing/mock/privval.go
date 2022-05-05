package mock

import (
	cryptocodec "github.com/reapchain/cosmos-sdk/crypto/codec"
	"github.com/reapchain/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/reapchain/cosmos-sdk/crypto/types"
	"github.com/reapchain/reapchain-core/crypto"
	tmproto "github.com/reapchain/reapchain-core/proto/reapchain-core/types"
	tmtypes "github.com/reapchain/reapchain-core/types"
	"github.com/reapchain/reapchain-core/vrfunc"
)

var _ tmtypes.PrivValidator = PV{}

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type PV struct {
	PrivKey cryptotypes.PrivKey
	Type              string
}

func (pv PV) GetType() (string, error) {
	return pv.Type, nil
}

func NewPV(validatorType string) PV {
	return PV{ed25519.GenPrivKey(), validatorType}
}

// GetPubKey implements PrivValidator interface
func (pv PV) GetPubKey() (crypto.PubKey, error) {
	return cryptocodec.ToTmPubKeyInterface(pv.PrivKey.PubKey())
}

// SignVote implements PrivValidator interface
func (pv PV) SignVote(chainID string, vote *tmproto.Vote) error {
	signBytes := tmtypes.VoteSignBytes(chainID, vote)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	vote.Signature = sig
	return nil
}

// SignProposal implements PrivValidator interface
func (pv PV) SignProposal(chainID string, proposal *tmproto.Proposal) error {
	signBytes := tmtypes.ProposalSignBytes(chainID, proposal)
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	proposal.Signature = sig
	return nil
}

func (pv PV) SignQrn(qrn *tmtypes.Qrn) error {
	signBytes := qrn.GetQrnBytesForSign()
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	qrn.Signature = sig
	return nil
}

func (pv PV) SignSettingSteeringMember(settingSteeringMember *tmtypes.SettingSteeringMember) error {
	signBytes := settingSteeringMember.GetSettingSteeringMemberBytesForSign()
	sig, err := pv.PrivKey.Sign(signBytes)
	if err != nil {
		return err
	}
	settingSteeringMember.Signature = sig
	return nil
}

func (pv PV) ProveVrf(vrf *tmtypes.Vrf) error {
	privateKey := vrfunc.PrivateKey(pv.PrivKey.Bytes())
	value, proof := privateKey.Prove(vrf.Seed)
	vrf.Value = value
	vrf.Proof = proof

	return nil
}
