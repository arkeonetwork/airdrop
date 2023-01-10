package types

type Token struct {
	Address      string `db:"address"`
	Name         string `db:"name"`
	Symbol       string `db:"symbol"`
	Chain        string `db:"chain"`
	GenesisBlock uint64 `db:"genesis_block"`
	Height       uint64 `db:"height"`
}

type Transfer struct {
	TxHash       string `db:"txHash"`
	LogIndex     uint64 `db:"log_index"`
	TokenAddress string `db:"token"`
	From         string `db:"transfer_from"`
	To           string `db:"transfer_to"`
	Value        string `db:"transfer_value"`
	BlockNumber  uint64 `db:"block_number"`
}
