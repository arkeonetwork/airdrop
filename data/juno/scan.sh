#!/bin/bash

#JUNO_NODE=https://lcd.juno.chaintools.tech
JUNO_NODE=https://lcd-juno.itastakers.com
#height=5899400
for (( height=5899456; height<5900001; height++))
do
  echo "requesting height ${height}"
  curl -s $JUNO_NODE/cosmos/staking/v1beta1/validators/junovaloper1jghdavdegcujzg96l08ghx086u0dwyjctw28d2/delegations -H "x-cosmos-block-height: $height" -o ./juno_${height}.json
  sleep 1
done

