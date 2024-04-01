# start main chain (decentrio)
git checkout test/upgrade
./local_node.sh

# proposal upgrade
realio-networkd tx gov submit-legacy-proposal software-upgrade multistaking --upgrade-height 30 --description "testing, testing, 1, 2, 3" --upgrade-info "testing" --title "Test Proposal" --from dev0 --keyring-backend test --no-validate --chain-id realionetworklocal_7777-1 --deposit 10000000ario --yes --home "$HOME/.realio-network-tmp" --fees 1000000000000000000000000ario --gas 1000000

# vote
realio-networkd tx gov vote 1 yes --from dev0 --keyring-backend test --chain-id realionetworklocal_7777-1 --yes --home "$HOME/.realio-network-tmp" --fees 1000000000000000000000000ario

# wait until upgrade height and cancel node process

# copy state to config
# export new state
realio-networkd export --home "$HOME/.realio-network-tmp" >> ~/.realio-network-tmp/config/state.json

# checkout intergrate-multistaking
git checkout intergrate-multistaking
# build new binary
make install
# run node
realio-networkd start --log_level "info" --minimum-gas-prices=0.00001ario --json-rpc.api eth,txpool,personal,net,debug,web3 --api.enable --home "$HOME/.realio-network-tmp"
