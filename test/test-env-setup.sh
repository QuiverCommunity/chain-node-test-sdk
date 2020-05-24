#!/bin/sh

# set block time to 2s to reduce testing time
sed -i 's/timeout_commit = "5s"/timeout_commit = "2s"/g' ~/.nsd/config/config.toml

# TODO should check alice and jack account is there already
# add keys 
nscli keys add alice --keyring-backend=test --recover <<< "script awake speed sweet obscure act shuffle mammal laptop lab holiday nuclear toy slogan oak direct couch vanish shiver silly neglect nurse enemy indoor"
nscli keys add jack --keyring-backend=test --recover <<< "retreat priority grab fence large clog cricket nature caution undo toy hand island evil fish coil entry dizzy tattoo business elder tilt arch ball"
nsd add-genesis-account $(nscli keys show alice -a --keyring-backend=test) 10000000alicecoin,10000000stake
nsd add-genesis-account $(nscli keys show jack -a --keyring-backend=test) 10000000jackcoin,10000000stake
