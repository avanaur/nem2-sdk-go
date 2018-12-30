// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestNewMosaicId(t *testing.T) {
	assert.Equal(t, big.NewInt(-3087871471161192663).Int64(), mosaicIdToBigInt(XemMosaicId).Int64())
}

func TestNewMosaicIdFromIdViaConstructor(t *testing.T) {
	mosaicId := bigIntToMosaicId(big.NewInt(-8884663987180930485))

	assert.Equal(t, big.NewInt(-8884663987180930485).Int64(), mosaicIdToBigInt(mosaicId).Int64())
}

//
func TestNewMosaic_ShouldCompareMosaicIdsForEquality(t *testing.T) {
	mosaicId := bigIntToMosaicId(big.NewInt(-8884663987180930485))
	mosaicId2 := bigIntToMosaicId(big.NewInt(-8884663987180930485))

	assert.Equal(t, mosaicId.String(), mosaicId2.String())
}
