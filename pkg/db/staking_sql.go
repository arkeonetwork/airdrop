package db

var (
	sqlFindAllStakingContracts = `select address,contract_name,chain,genesis_block,height from staking`
)
