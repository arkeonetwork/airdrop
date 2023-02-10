package db

const (
	sqlFindLatestIndexedCosmosLPBlock = `
	select coalesce(max(block_number),0) from cosmos_lp_events where chain = $1
`
)
