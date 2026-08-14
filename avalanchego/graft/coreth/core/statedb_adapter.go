package core

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/avalanchego/graft/coreth/core/tracing"
	"github.com/holiman/uint256"
)

// stateDBAdapter wraps libevm's StateDB to implement tracing.StateDB interface
type stateDBAdapter struct {
	*state.StateDB
}

// newStateDBAdapter creates a new adapter for libevm's StateDB
func newStateDBAdapter(db *state.StateDB) tracing.StateDB {
	return &stateDBAdapter{StateDB: db}
}

// GetState implements tracing.StateDB by discarding variadic args
func (s *stateDBAdapter) GetState(addr common.Address, key common.Hash, _ ...any) common.Hash {
	return s.StateDB.GetState(addr, key)
}

// Ensure we implement all other required methods by embedding
func (s *stateDBAdapter) GetBalance(addr common.Address) *uint256.Int {
	return s.StateDB.GetBalance(addr)
}

func (s *stateDBAdapter) GetNonce(addr common.Address) uint64 {
	return s.StateDB.GetNonce(addr)
}

func (s *stateDBAdapter) GetCode(addr common.Address) []byte {
	return s.StateDB.GetCode(addr)
}

func (s *stateDBAdapter) GetCodeHash(addr common.Address) common.Hash {
	return s.StateDB.GetCodeHash(addr)
}

func (s *stateDBAdapter) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	return s.StateDB.GetTransientState(addr, key)
}

func (s *stateDBAdapter) Exist(addr common.Address) bool {
	return s.StateDB.Exist(addr)
}

func (s *stateDBAdapter) GetRefund() uint64 {
	return s.StateDB.GetRefund()
}
