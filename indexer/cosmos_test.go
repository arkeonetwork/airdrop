package indexer

import (
	"fmt"
	"testing"

	"github.com/ArkeoNetwork/airdrop/pkg/utils"
	arkutils "github.com/ArkeoNetwork/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestIndexDelegations(t *testing.T) {
	c := arkutils.ReadDBConfig(utils.GetEnvPath())
	if c == nil {
		fmt.Print("error: no config loaded")
		return
	}

	// height := int64(12940516)
	// height := int64(12940505)
	// height := int64(12940507)
	// height := int64(13742000)
	height := int64(12940754)
	chain := "GAIA"
	params := CosmosIndexerParams{Chain: chain, DB: *c}
	indxr, err := NewCosmosIndexer(params)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	assert.NotNil(t, indxr)
	err = indxr.indexDelegations(height)
	assert.Nil(t, err)
}

func TestIndexLP(t *testing.T) {
	c := arkutils.ReadDBConfig(utils.GetEnvPath())
	if c == nil {
		fmt.Print("error: no config loaded")
		return
	}

	height := int64(9286118)
	chain := "THOR"
	params := CosmosIndexerParams{Chain: chain, DB: *c}
	indxr, err := NewCosmosIndexer(params)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	assert.NotNil(t, indxr)
	err = indxr.indexLP(height, "ETH.FOX-0XC770EEFAD204B5180DF6A14EE197D99D808EE52D")
	assert.Nil(t, err)
}
