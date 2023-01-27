package db

var (
	sqlFindAllChains = `select name,rpc_url,lcd_url,snapshot_start_block,snapshot_end_block,decimals from chains`
	sqlFindChain     = `select name,rpc_url,lcd_url,snapshot_start_block,snapshot_end_block,decimals from chains where name = $1`
)
