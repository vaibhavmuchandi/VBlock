rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile vlmOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile vlmOrgChannel -outputCreateChannelTx ./config/vlmchannel.tx -channelID vlmchannel
