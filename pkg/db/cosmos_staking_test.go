package db

import (
	"testing"

	"github.com/ArkeoNetwork/common/utils"
	"github.com/stretchr/testify/assert"
)

var envPath = "/Users/adamsamere/chaintech/oss/arkeo/airdrop/docker/dev/docker.env"

// #javanaming
func TestFindLatestIndexedCosmosStakingBlock(t *testing.T) {
	c := utils.ReadDBConfig(envPath)
	if !assert.NotNil(t, c) {
		t.FailNow()
	}
	d, err := New(*c)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	if !assert.NotNil(t, d) {
		t.FailNow()
	}

	latest, err := d.FindLatestIndexedCosmosStakingBlock()
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	if assert.NotNil(t, latest) {
		assert.Greater(t, latest, int64(0))
	}
}
