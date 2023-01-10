package db

var (
	sqlUpsertTransferEvent = `
	insert into transfers(txhash,log_index,token,transfer_from,transfer_to,transfer_value,block_number)
	values ($1,$2,$3,$4,$5,$6,$7)
	on conflict on constraint transfers_txhash_log_index_unique
	do update set updated = now()
	where transfers.txhash = $1
	  and transfers.log_index = $2
	returning id, created, updated
	`
)
