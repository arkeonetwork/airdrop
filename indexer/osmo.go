package indexer

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/ArkeoNetwork/airdrop/pkg/utils"
	"github.com/pkg/errors"
)

type BondedAccount struct {
	Account string `json:"account"`
	Bonded  []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"bonded"`
}

const batchSize = 1000

func (c *CosmosIndexer) IndexOsmoLP(height int64, dataDir string) error {
	jsonFile := fmt.Sprintf("%s/bonded_%d.json", dataDir, height)
	return c.indexOsmoLP(height, jsonFile)
}

func (c *CosmosIndexer) indexOsmoLP(height int64, jsonFile string) error {
	log := log.WithField("height", fmt.Sprintf("%d", height))
	_ = log
	raw, err := os.ReadFile(jsonFile)
	if err != nil {
		return errors.Wrapf(err, "error reading json file")
	}

	accounts := make([]*BondedAccount, 0, 8194)
	if err := json.Unmarshal(raw, &accounts); err != nil {
		return errors.Wrapf(err, "error unmarshaling json")
	}

	batch := make([]*types.OsmoLP, 0, batchSize)
	for i, acct := range accounts {
		for _, bond := range acct.Bonded {
			if bond.Denom == "uosmo" {
				iqty, err := strconv.ParseInt(bond.Amount, 10, 64)
				if err != nil {
					return errors.Wrapf(err, "error parsing qty %s", bond.Amount)
				}
				qty := utils.FromBaseUnits(iqty, 6)
				batch = append(batch, &types.OsmoLP{BlockNumber: height, Account: acct.Account, QtyOsmo: qty})
				break
			}
		}
		if i != 0 && i%batchSize == 0 {
			if err := c.db.InsertOsmoLP(batch); err != nil {
				return errors.Wrapf(err, "error inserting batch of %d", len(batch))
			}
			batch = make([]*types.OsmoLP, 0, batchSize)
		}
	}
	if len(batch) > 0 {
		log.Infof("inserting final batch of %d", len(batch))
		if err := c.db.InsertOsmoLP(batch); err != nil {
			return errors.Wrapf(err, "error inserting final batch of %d", len(batch))
		}
	}
	return nil
}
