

#!/bin/bash
# chmod +x querySingleCar.sh

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

peer chaincode query -C $1 -n $2 -c '{"Args":["queryCar","CAR0"]}'