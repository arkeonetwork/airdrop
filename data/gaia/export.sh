#!/bin/sh

outdir=.
# page 1
page=1
pageKey=""
nodeIp="10.0.191.10"
ssValoper=cosmosvaloper199mlc7fr6ll5t54w7tts7f4s0cvnqgc59nmuxf
reqUrl="http://${nodeIp}:1317/cosmos/staking/v1beta1/validators/${ssValoper}/delegations"
blockHeight="13050670"

while true
do
   echo "requesting height ${blockHeight} delegators to ${ssValoper} page $page with key $pageKey and writing $outdir/page$page.json"
   curl -s "$reqUrl" --data-urlencode "pagination.key=$pageKey" -H "x-cosmos-block-height: ${blockHeight}" -o $outdir/page$page.json

   pageKey=`cat $outdir/page$page.json | jq -r .pagination.next_key`
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

