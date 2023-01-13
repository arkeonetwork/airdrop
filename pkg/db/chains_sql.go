package db

var (
	sqlFindAllChains = `select name,rpc_url,snapshot_start_block,snapshot_end_block from chains`
)
