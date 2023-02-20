package db

const (
	sqlInsertOsmoLP = `
		insert into osmo_lp(block_number, account, qty_osmo)
		values ($1,$2,$3)
	`
)
