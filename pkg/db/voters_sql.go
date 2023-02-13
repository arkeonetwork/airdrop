package db

var (
	sqlInsertVoters = `
		insert into snapshot_voters(address)
		values ($1)
		on conflict on constraint snapshot_voters_address_unique
		do update set updated = now()
		returning id, created, updated
	`
	sqlFindVoterByAddress = `select address from snapshot_voters where address = $1`
)
