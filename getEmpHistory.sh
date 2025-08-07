#!/bin/bash
# chmod +x getEmpHistory.sh diye neya lagbe
# Usage: ./getEmpHistory.sh <CHANNEL_NAME> <CHAINCODE_NAME> <EMPLOYEE_ID>

if [ "$#" -ne 3 ]; then
  echo "Usage: $0 <CHANNEL_NAME> <CHAINCODE_NAME> <EMPLOYEE_ID>"
  exit 1
fi

CHANNEL_NAME=$1
CC_NAME=$2
EMPLOYEE_ID=$3

export FABRIC_CFG_PATH=${PWD}/artifacts/channel/config/
export PEER0_ORG1_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
export ORDERER_CA=${PWD}/artifacts/channel/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export PEER0_ORG2_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt

setGlobalsForPeer0Org1() {
    export CORE_PEER_LOCALMSPID="Org1MSP"
    export CORE_PEER_MSPCONFIGPATH=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
}

setGlobalsForPeer0Org1

peer chaincode query -C "$CHANNEL_NAME" -n "$CC_NAME" \
  -c '{"function":"GetEmployeeHistory","Args":["'"$EMPLOYEE_ID"'"]}'
