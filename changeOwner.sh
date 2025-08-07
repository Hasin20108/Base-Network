

#!/bin/bash

# chmod +x changeowner.sh diye neya lagbe

CC_NAME="fabcar"
export PEER0_ORG1_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export FABRIC_CFG_PATH=${PWD}/artifacts/channel/config/
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt


setGlobalsForPeer0Org1() {
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051

    # Set this based on whether your network uses TLS or not
    export CORE_PEER_TLS_ENABLED=true  
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
}

setGlobalsForPeer0Org1
export ORDERER_CA=${PWD}/artifacts/channel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export PEER0_ORG2_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

peer chaincode invoke \
  -o localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls \
  --cafile "$ORDERER_CA" \
  -C mychannel \
  -n fabcar \
  --peerAddresses localhost:7051 \
  --tlsRootCertFiles "$PEER0_ORG1_CA" \
  --peerAddresses localhost:9051 \
  --tlsRootCertFiles "$PEER0_ORG2_CA" \
  -c '{"function":"changeCarOwner","Args":["CAR0","Nasir"]}'