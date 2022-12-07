BINARY="./nibid"
TXFLAG="--gas-prices 0.1$DENOM --gas auto --gas-adjustment 1.3 -y -b block --chain-id $CHAIN_ID --node $RPC"
DEFAULT_DEV_ADDRESS="juno16g2rahf5846rxzp3fwlswy08fz8ccuwk03k57y"

# validator addr
VALIDATOR_ADDR=$($BINARY keys show validator --address)
echo "Validator address:"
echo "$VALIDATOR_ADDR"

BALANCE_1=$($BINARY q bank balances $VALIDATOR_ADDR)
echo "Pre-store balance:"
echo "$BALANCE_1"

# you ideally want to run locally, get a user and then
# pass that addr in here
echo "Address to deploy contracts: $DEFAULT_DEV_ADDRESS"
echo "TX Flags: $TXFLAG"

CONTRACT_CODE=$($BINARY tx wasm store "./scripts/e2e/contracts/whoami.wasm" --from validator $TXFLAG --output json | jq -r '.logs[0].events[-1].attributes[-1].value')
echo "Stored: $CONTRACT_CODE"

BALANCE_2=$($BINARY q bank balances $VALIDATOR_ADDR)
echo "Post-store balance:"
echo "$BALANCE_2"
