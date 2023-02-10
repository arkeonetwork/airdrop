package snapshot

import (
	"fmt"

	"github.com/ArkeoNetwork/airdrop/pkg/db"
	"github.com/ArkeoNetwork/directory/pkg/logging"
)

type SnapshotIndexerAppParams struct {
	SnapshotStart uint64
	SnapshotEnd   uint64
	db.DBConfig
}

type SnapshotIndexerApp struct {
	params SnapshotIndexerAppParams
	db     *db.AirdropDB
}

type SnapshotProposalVoter struct {
	address string
}

var log = logging.WithoutFields()

func NewSnapshotIndexer(params SnapshotIndexerAppParams) *SnapshotIndexerApp {
	d, err := db.New(params.DBConfig)
	if err != nil {
		panic(fmt.Sprintf("error connecting to the db: %+v", err))
	}
	return &SnapshotIndexerApp{params: params, db: d}
}

func (app *SnapshotIndexerApp) Start() {
	log.Info("starting indexing snapshot data")
	err := app.IndexSnapshotData()
	if err != nil {
		panic(fmt.Sprintf("error indexing snapshot data: %+v", err))
	}
	log.Info("finished indexing snapshot data")
}

func (app *SnapshotIndexerApp) IndexSnapshotData() error {
	voters, err := app.GetSnapshotProposalVoters()
	if err != nil {
		panic(fmt.Sprintf("error getting snapshot proposal voters: %+v", err))
	}
	// app.db.saveVoters(voters)
	log.Info(voters)
	return nil
}

