# pylons

Pylons daemon is the project responsible for providing rest api and amino endpoints necessary for interacting with pylons eco system


## How to

- Install

```

go install ./cmd/pylonsd
go install ./cmd/pylonscli

```

- create a genesis block and some users

```

# Initialize configuration files and genesis file
pylonsd init --chain-id namechain

# Copy the `Address` output here and save it for later use 
# [optional] add "--ledger" at the end to use a Ledger Nano S 
pylonscli keys add jack

# Copy the `Address` output here and save it for later use
pylonscli keys add alice

# Add both accounts, with coins to the genesis file
nsd add-genesis-account $(pylonscli keys show jack -a) 100pylons,1000jackcoin
nsd add-genesis-account $(pylonscli keys show alice -a) 100pylons,1000alicecoin

# Configure your CLI to eliminate need for chain-id flag
pylonscli config chain-id namechain
pylonscli config output json
pylonscli config indent true
pylonscli config trust-node true


```

- start the `pylonsd` node

```

pylonsd start

```

- play with the api

```
pylonscli tx pylons get-pylons --from alice
```
