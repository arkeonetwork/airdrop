package db

var (
	sqlInsertTransferEvent = `
	insert into transfers(txhash,token,transfer_from,transfer_to,transfer_value,block_number)
	values ($1,$2,$3,$4,$5,$6,$7)
	returning id, created, updated
	`
)
