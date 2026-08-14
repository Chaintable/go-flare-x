// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import (
	"math/big"

	"github.com/ava-labs/libevm/core/types"
)

// SetEffectiveGasPrice sets the effective gas price on a receipt
func (r *types.Receipt) SetEffectiveGasPrice(tx *types.Transaction, baseFee *big.Int) {
	r.EffectiveGasPrice = tx.EffectiveGasTip(baseFee)
	if baseFee != nil {
		r.EffectiveGasPrice = new(big.Int).Add(r.EffectiveGasPrice, baseFee)
	}
}
