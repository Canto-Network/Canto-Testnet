# Becoming A Validator
**How to validate on the Canto Testnet**
*This is the Canto Testnet-1 (canto_7777-1)*

​> Genesis (Published)[https://github.com/Canto-Network/Canto-Testnet/Networks/Testnet/raw/main/genesis.json]

​> Peers (Published)[https://hack.md]

## Hardware Requirements
**Minimum**
* 16 GB RAM
* 100 GB NVME SSD
* 3.2 GHz x4 CPU

**Recommended**
* 32 GB RAM
* 500 GB NVME SSD
* 4.2 GHz x6 CPU 

## Operating System 

> Linux (x86_64) or Linux (amd64) Reccomended Arch Linux

**Dependencies**
> Prerequisite: go1.18+ required.   ​
* Arch Linux: `pacman -S go`
* Ubuntu: `sudo snap install go --classic`

> Prerequisite: git. 
* Arch Linux: `pacman -S git`
* Ubuntu: `sudo apt-get install git`

> Optional requirement: GNU make. 
* Arch Linux: `pacman -S make`
* Ubuntu: `sudo apt-get install make`

## Cantod Installation Steps

**Clone git repository**

```bash
git clone https://github.com/Canto-Network/Canto-Testnet
cd Canto-Testnet
cd cmd/cantod
go install -tags ledger ./...
mv $HOME/go/bin/cantod /usr/bin/
```
**Generate keys**

* `cantod keys add [key_name]`

* `craftd keys add [key_name] --recover` to regenerate keys with your mnemonic

* `craftd keys add [key_name] --ledger` to generate keys with ledger device

## Validator setup instructions

* Install cantod binary

* Initialize node: `cantod init <moniker> --chain-id canto_7777-1`

* Download the Genesis file: `wget https://github.com/Canto-Network/Canto-Testnet/raw/main/genesis.json $HOME/.cantod/config/`
 
* Edit the minimum-gas-prices in ${HOME}/.cantod/config/app.toml `sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.01acanto"/g' $HOME/.cantod/config/app.toml`

* Start craftd by creating a systemd service to run the node in the background
`nano /etc/systemd/system/cantod.service`
> Copy and paste the following file into your service file. Be sure to edit as you see fit.

```bash
[Unit]
Description=Craft Node
After=network.target
​
[Service]
Type=simple
User=root
WorkingDirectory=/root/
ExecStart=/root/go/bin/craftd start
Restart=on-failure
StartLimitInterval=0
RestartSec=3
LimitNOFILE=65535
LimitMEMLOCK=209715200
​
[Install]
WantedBy=multi-user.target
```
Reload the service files `sudo systemctl daemon-reload` Create the symlinlk `sudo systemctl enable cantod.service` Start the node sudo `systemctl start cantod && journalctl -u cantod -f`

### Create Validator Transaction

cantod tx staking create-validator \
--from {{KEY_NAME}} \
--chain-id canto_7777-1 \
--moniker="<VALIDATOR_NAME>" \
--commission-max-change-rate=0.01 \
--commission-max-rate=1.0 \
--commission-rate=0.05 \
--details="<description>" \
--security-contact="<contact_information>" \
--website="<your_website>" \
--pubkey $(cantod tendermint show-validator) \
--min-self-delegation="1" \
--amount <token delegation>acanto \
