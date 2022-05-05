package mock_test

import (
	"testing"

	tmproto "github.com/reapchain/reapchain-core/proto/reapchain-core/types"
	tmtypes "github.com/reapchain/reapchain-core/types"
	"github.com/stretchr/testify/require"

	"github.com/reapchain/ibc-go/testing/mock"
)

const chainID = "testChain"

func TestGetPubKey(t *testing.T) {
	pv := mock.NewPV()
	pk, err := pv.GetPubKey()
	require.NoError(t, err)
	require.Equal(t, "ed25519", pk.Type())
}

func TestSignVote(t *testing.T) {
	pv := mock.NewPV()
	pk, _ := pv.GetPubKey()

	vote := &tmproto.Vote{Height: 2}
	pv.SignVote(chainID, vote)

	msg := tmtypes.VoteSignBytes(chainID, vote)
	ok := pk.VerifySignature(msg, vote.Signature)
	require.True(t, ok)
}

func TestSignProposal(t *testing.T) {
	pv := mock.NewPV()
	pk, _ := pv.GetPubKey()

	proposal := &tmproto.Proposal{Round: 2}
	pv.SignProposal(chainID, proposal)

	msg := tmtypes.ProposalSignBytes(chainID, proposal)
	ok := pk.VerifySignature(msg, proposal.Signature)
	require.True(t, ok)
}
