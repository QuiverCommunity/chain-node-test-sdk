#!/bin/sh

# remove existing genesis configuration
rm -rf ~/.nsd

# init genesis json
nsd init quivernode --chain-id quiverchain

# set cli configuration
nscli config chain-id quiverchain
nscli config output json
nscli config indent true
nscli config trust-node true

# # ubuntu, docker set block time to 2s to reduce testing time
# sed -i 's/timeout_commit = "5s"/timeout_commit = "2s"/g' ~/.nsd/config/config.toml

# osx set block time to 2s to reduce testing time
sed -i '' 's/timeout_commit = "5s"/timeout_commit = "2s"/g' ~/.nsd/config/config.toml

# remove keys that were existing already
nscli keys delete alice --keyring-backend=test
nscli keys delete jack --keyring-backend=test

# add keys 
nscli keys add alice --keyring-backend=test --recover <<< "script awake speed sweet obscure act shuffle mammal laptop lab holiday nuclear toy slogan oak direct couch vanish shiver silly neglect nurse enemy indoor"
nscli keys add jack --keyring-backend=test --recover <<< "retreat priority grab fence large clog cricket nature caution undo toy hand island evil fish coil entry dizzy tattoo business elder tilt arch ball"

# add genesis coins
nsd add-genesis-account $(nscli keys show alice -a --keyring-backend=test) 10000000alicecoin,100000000stake
nsd add-genesis-account $(nscli keys show jack -a --keyring-backend=test) 10000000jackcoin,100000000stake

# # collect genesis transactions
nsd gentx --name jack --keyring-backend=test
nsd collect-gentxs

# # run node
nsd start