# Cantod Hard Fork Steps

**Pull updates from git repository - or clone if you haven't already**

1. Stop the cantod process
2. Remove .cantod `rm -rf ~/.cantod`
3. cd into Canto-Testnet and pull latest changes `git pull`
4. Init cantod again `cantod init <MONIKER> --chain-id canto_9624-1`
5. Update minimum gas price in .cantod/config/app.toml to 0.0001acanto
6. cp updated genesis.json into config folder `cp ~/Canto-Testnet/Networks/Testnet/genesis/genesis.json ~/.cantod/config/`
7. Add the following persistent peer to .cantod/config.toml:
    - f7f2a6263b05c36fcce16183a5c206cada503738@137.184.130.158:26656
8. Start cantod process
9. You should immediately see blocks syncing


## Becoming a Validator:

1. Add a key `cantod keys add [key_name]`
2. Request tokens
3. Send a validator transaction 

```bash
cantod tx staking create-validator \
--from {{KEY_NAME}} \
--chain-id canto_9624-1 \
--moniker="<VALIDATOR_NAME>" \
--commission-max-change-rate=0.01 \
--commission-max-rate=1.0 \
--commission-rate=0.05 \
--details="<description>" \
--security-contact="<contact_information>" \
--website="<your_website>" \
--pubkey $(cantod tendermint show-validator) \
--min-self-delegation="1" \
--amount <token amount>acanto \
--node http://137.184.130.158:26657
```

4. Query the validator set to check your entry `cantod query staking validators`