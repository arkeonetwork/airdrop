package db

var (
	sqlFindTokensByChain = `select * from tokens where chain = $1`
	sqlFindAllChains     = `select distinct(chain) from tokens`
)
