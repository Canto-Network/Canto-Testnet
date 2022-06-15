### Updating the genesis file (Canto Hard Fork)

* Step 1) stop cantod `systemctl stop cantod`

* Step 2) remove the data directory `rm -rf ~/.cantod/data/`

* Step 3) create data directory `mkdir ~/.cantod/data/`

* Step 4) create priv_validator_state.json file in data directory `nano ~/.cantod/data/priv_validator_state.json`

* Step 5) copy empty state file to the `priv_validator_state.json` 

```bash
{
   "height": "0",
   "round": 0,
   "step": 0
 }
 ```
* Step 6) add new peers  `sed -i 's/persistent_peers = ""/persistent_peers = "4043448542e03beaa0c9ec3a6e4ab2928c51b86c@137.184.130.158:26656,b0a85e37973ba1e2a304c9c5e65c454c218eb2c0@canto.p2p.chandrastation.com:26656"/g' ~/.cantod/config/config.toml`

* Step 7) get new genesis - if you have access to the repo on your machine simply `cp Canto-Network/Canto-Testnet/Networks/Testnet/genesis/genesis.json ~/.cantod/config/` if you do not have access to the repo on your machine you can `rm ~/.cantod/config/genesis.json` and copy the contents of the [new genesis](https://github.com/Canto-Network/Canto-Testnet/blob/main/Networks/Testnet/genesis/genesis.json) into `nano ~/.cantod/config/genesis.json`

* Step 8) restart cantod `systemctl restart cantod && journalctl -u cantod -f`
