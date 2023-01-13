package types

type Token struct {
	Address      string `db:"address"`
	Name         string `db:"name"`
	Symbol       string `db:"symbol"`
	Decimals     uint64 `db:"decimals"`
	Chain        string `db:"chain"`
	GenesisBlock uint64 `db:"genesis_block"`
	Height       uint64 `db:"height"`
}

type Chain struct {
<<<<<<< HEAD
	Name               string `db:"name"`
	RpcUrl             string `db:"rpc_url"`
	SnapshotStartBlock uint64 `db:"snapshot_start_block"`
	SnapshotEndBlock   uint64 `db:"snapshot_end_block"`
=======
	Name   string `db:"name"`
	RpcUrl string `db:"rpc_url"`
>>>>>>> f98d274 (adds multichain functionality)
}

type Transfer struct {
	TxHash       string  `db:"txHash"`
	LogIndex     uint64  `db:"log_index"`
	TokenAddress string  `db:"token"`
	From         string  `db:"transfer_from"`
	To           string  `db:"transfer_to"`
	Value        float64 `db:"transfer_value"` // decimal version of transfer_value
	BlockNumber  uint64  `db:"block_number"`
}
