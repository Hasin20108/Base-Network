
# chmod -R 0755 ./crypto-config
# # Delete existing artifacts
# rm -rf ./crypto-config
# rm genesis.block mychannel.tx
# rm -rf ../../channel-artifacts/*

# #Generate Crypto artifactes for organizations
# cryptogen generate --config=./crypto-config.yaml --output=./crypto-config/



# # System channel
# SYS_CHANNEL="sys-channel"

# # channel name defaults to "mychannel"
# CHANNEL_NAME="mychannel"

# echo $CHANNEL_NAME

# # Generate System Genesis block
# configtxgen -profile OrdererGenesis -configPath . -channelID $SYS_CHANNEL  -outputBlock ./genesis.block


# # Generate channel configuration block
# configtxgen -profile BasicChannel -configPath . -outputCreateChannelTx ./mychannel.tx -channelID $CHANNEL_NAME

# echo "#######    Generating anchor peer update for Org1MSP  ##########"
# configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

# echo "#######    Generating anchor peer update for Org2MSP  ##########"
# configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP





#!/bin/bash

# Ensure binaries are available
export PATH=$PATH:$PWD/bin

# Fix permissions (optional, only if needed)
chmod -R 0755 ./crypto-config

# Delete existing artifacts safely
rm -rf ./crypto-config
[ -f genesis.block ] && rm genesis.block
[ -f mychannel.tx ] && rm mychannel.tx
rm -rf ../../channel-artifacts/*
rm -rf ./genesis.block
# Generate Crypto material
cryptogen generate --config=./crypto-config.yaml --output=./crypto-config/

# System channel
SYS_CHANNEL="sys-channel"

# Application channel name
CHANNEL_NAME="mychannel"

echo $CHANNEL_NAME

# Generate system genesis block
configtxgen -profile OrdererGenesis -configPath . -channelID $SYS_CHANNEL -outputBlock ./genesis.block
# configtxgen -profile OrdererGenesis -configPath . -channelID $SYS_CHANNEL  -outputBlock ./genesis.block

# Generate channel config tx
configtxgen -profile BasicChannel -configPath . -outputCreateChannelTx ./mychannel.tx -channelID $CHANNEL_NAME

# Generate anchor peer updates
echo "#######    Generating anchor peer update for Org1MSP  ##########"
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP

echo "#######    Generating anchor peer update for Org2MSP  ##########"
configtxgen -profile BasicChannel -configPath . -outputAnchorPeersUpdate ./Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP
