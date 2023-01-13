package db

var (
<<<<<<< HEAD
	sqlFindAllChains = `select name,rpc_url,snapshot_start_block,snapshot_end_block from chains`
=======
	sqlFindAllChains = `select name,rpc_url from chains`
>>>>>>> f98d274 (adds multichain functionality)
)
