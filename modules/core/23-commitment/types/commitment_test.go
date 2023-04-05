package types_test

import (
	"testing"

	"github.com/reapchain/cosmos-sdk/store/iavl"
	"github.com/reapchain/cosmos-sdk/store/rootmulti"
	storetypes "github.com/reapchain/cosmos-sdk/store/types"
	"github.com/reapchain/reapchain-core/libs/log"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type MerkleTestSuite struct {
	suite.Suite

	store     *rootmulti.Store
	storeKey  *storetypes.KVStoreKey
	iavlStore *iavl.Store
}

func (suite *MerkleTestSuite) SetupTest() {
	db := dbm.NewMemDB()
	suite.store = rootmulti.NewStore(db, log.NewNopLogger())

	suite.storeKey = storetypes.NewKVStoreKey("iavlStoreKey")

	suite.store.MountStoreWithDB(suite.storeKey, storetypes.StoreTypeIAVL, nil)
	suite.store.LoadVersion(0)

	suite.iavlStore = suite.store.GetCommitStore(suite.storeKey).(*iavl.Store)
}

func TestMerkleTestSuite(t *testing.T) {
	suite.Run(t, new(MerkleTestSuite))
}
