#!/bin/sh

outdir=.
# page 6. have 1-5 from manual dl
#page=6
#pageKey="FAoeaQZfhHwHAGDafwDGWqAmgHoSFDO6L50rtCqJveSho4r7vAcxpUIc"
page=22
pageKey="FCzUdP0EVy4hE2xskAx0pGWyJbELFDKf8iyJ0K2BsplobJtM8BU8cCKQ"
# cat ./page1.json | jq .pagination

reqUrl="http://10.0.166.52:1317/cosmos/staking/v1beta1/validators/osmovaloper1xf9zpq5kpxks49cg606tzd8qstaykxgt2vs0d5/delegations"

while true
do
   echo "requesting page $page with key $pageKey"
   curl -s "$reqUrl" --data-urlencode "pagination.key=$pageKey" -H 'x-cosmos-block-height: 7117606' -o $outdir/page$page.json

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