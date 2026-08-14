package tracing

import (
	"math/big"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/params"
	"github.com/holiman/uint256"
)

// StateDB gives tracers access to the whole state.
type StateDB interface {
	GetBalance(common.Address) *uint256.Int
	GetNonce(common.Address) uint64
	GetCode(common.Address) []byte
	GetCodeHash(common.Address) common.Hash
	GetState(common.Address, common.Hash, ...any) common.Hash
	GetTransientState(common.Address, common.Hash) common.Hash
	Exist(common.Address) bool
	GetRefund() uint64
}

// VMContext provides the context for the EVM execution.
type VMContext struct {
	Coinbase    common.Address
	BlockNumber *big.Int
	Time        uint64
	Random      *common.Hash
	BaseFee     *big.Int
	StateDB     StateDB
}

type (
	TxStartHook         = func(vmContext *VMContext, tx *types.Transaction, from common.Address)
	TxEndHook           = func(receipt *types.Receipt, err error)
	BlockchainInitHook  = func(chainConfig *params.ChainConfig)
	CloseHook           = func()
	BlockStartHook      = func(block *types.Block)
	BlockEndHook        = func(err error)
	GenesisBlockHook    = func(genesis *types.Block, alloc types.GenesisAlloc)
	CommitHook          = func(originRoot common.Hash, root common.Hash, destructs map[common.Hash]struct{}, accounts map[common.Hash][]byte, accountsOrigin map[common.Address][]byte, storages map[common.Hash]map[common.Hash][]byte, storagesOrigin map[common.Address]map[common.Hash][]byte, codes map[common.Hash][]byte)
	LogHook             = func(log *types.Log)
	OnSystemCallStartHookV2 = func(vm *VMContext)
)

type Hooks struct {
	OnTxStart               TxStartHook
	OnTxEnd                 TxEndHook
	OnBlockchainInit        BlockchainInitHook
	OnClose                 CloseHook
	OnBlockStart            BlockStartHook
	OnSystemCallStartHookV2 OnSystemCallStartHookV2
	OnBlockEnd              BlockEndHook
	OnGenesisBlock          GenesisBlockHook
	OnLog                   LogHook
	OnCommit                CommitHook
}
