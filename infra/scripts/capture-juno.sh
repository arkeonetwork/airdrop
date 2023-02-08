#!/bin/bash

# curl "$JUNO_NODE/cosmos/staking/v1beta1/validators/junovaloper1jghdavdegcujzg96l08ghx086u0dwyjctw28d2/delegations"

outdir=$HOME/data/airdrop/juno/capture/blocks
page=1
pageKey=""
JUNO_NODE=https://lcd-juno.itastakers.com
ssValoper=junovaloper1jghdavdegcujzg96l08ghx086u0dwyjctw28d2
reqUrl="$JUNO_NODE/cosmos/staking/v1beta1/validators/${ssValoper}/delegations"

blockHeight=`curl -s "$JUNO_NODE/blocks/latest" | jq -r .block.header.height`
mkdir -p $outdir/$blockHeight

while true
do
  echo "requesting delegators to ${ssValoper} page $page with key $pageKey and writing $outdir/$blockHeight/page$page.json"
  curl -s "$reqUrl" --data-urlencode "pagination.key=$pageKey" -o $outdir/$blockHeight/page$page.json

  pageKey=`cat $outdir/$blockHeight/page$page.json | jq -r .pagination.next_key`
  if [ "$pageKey" == "null" ]; then
    echo "null no more pages"
    break
  fi
  if [ "$pageKey" == "" ]; then
    echo "empty no more pages"
    break
  fi
  page=$((page + 1))
done

echo DONE
