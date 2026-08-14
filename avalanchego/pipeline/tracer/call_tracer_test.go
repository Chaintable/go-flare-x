package tracer

import (
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/avalanchego/graft/coreth/core/tracing"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/core/vm"
)

func TestCallTracerSkipsEmptyCallstackHooks(t *testing.T) {
	tracer := newCallTracerRaw()
	tx := types.NewTx(&types.LegacyTx{Gas: 21_000})

	tracer.OnTxStart(&tracing.VMContext{}, tx, common.Address{})
	tracer.CaptureState(0, vm.SSTORE, 0, 0, nil, nil, 0, nil)
	tracer.OnLog(&types.Log{})
	tracer.OnTxEnd(&types.Receipt{}, nil)
}
