package indexer

import (
	"fmt"
	"testing"

	"github.com/ArkeoNetwork/common/utils"
	"github.com/stretchr/testify/assert"
)

const envPath = "/Users/adamsamere/chaintech/oss/arkeo/airdrop/docker/dev/docker.env"

func TestIndexDelegations(t *testing.T) {
	c := utils.ReadDBConfig(envPath)
	if c == nil {
		fmt.Print("error: no config loaded")
		return
	}

	// height := int64(12940516)
	// height := int64(12940505)
	// height := int64(12940507)
	height := int64(13742000)
	chain := "GAIA"
	params := CosmosIndexerParams{Chain: chain, DB: *c}
	indxr, err := NewCosmosIndexer(params)
	assert.Nil(t, err)
	assert.NotNil(t, indxr)
	err = indxr.indexDelegations(height)
	assert.Nil(t, err)
}
