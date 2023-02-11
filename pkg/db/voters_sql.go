package db

var (
	sqlInsertVoters = `
		insert into snapshot_voters(address)
		values ($1)
		returning id, created, updated
	`
	sqlFindVoterByAddress = `select address from snapshot_voters where address = $1`
)
