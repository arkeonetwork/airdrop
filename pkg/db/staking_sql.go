package db

var (
	sqlFindAllStakingContracts     = `select address,contract_name,chain,genesis_block,height from staking_contracts`
	sqlFindStakingContractsByName  = `select address,contract_name,chain,genesis_block,height from staking_contracts where contract_name = $1`
	sqlUpdateStakingContractHeight = `update staking_contracts set height = $1 where address = $2 and chain = $3`
	sqlUpsertStakingEvent          = `
	insert into staking_events(txhash,log_index,token,staking_contract,staker,stake_value,block_number,chain)
	values ($1,$2,$3,$4,$5,$6,$7,$8)
	on conflict on constraint staking_events_txhash_log_index_unique
	do update set updated = now()
	where staking_events.txhash = $1
	  and staking_events.log_index = $2
	returning id, created, updated`
)
