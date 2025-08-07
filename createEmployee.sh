
#!/bin/bash
# chmod +x createEmployee.sh diye neya lagbe
# Usage: ./createEmployee.sh <CHANNEL_NAME> <CHAINCODE_NAME> <EMP_ID> <NAME> <AGE> <EMPLOYER> <CONTRIBUTION_AMOUNT>

if [ "$#" -ne 7 ]; then
  echo "Usage: $0 <CHANNEL_NAME> <CHAINCODE_NAME> <EMP_ID> <NAME> <AGE> <EMPLOYER> <CONTRIBUTION_AMOUNT>"
  exit 1
fi

CHANNEL_NAME=$1
CC_NAME=$2
EMP_ID=$3
NAME=$4
AGE=$5
EMPLOYER=$6
CONTRIBUTION_AMOUNT=$7

export FABRIC_CFG_PATH=${PWD}/artifacts/channel/config/
export PEER0_ORG1_CA=${PWD}/artifacts/channel/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
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

peer chaincode invoke -o localhost:7050 \
  --ordererTLSHostnameOverride orderer.example.com \
  --tls $CORE_PEER_TLS_ENABLED \
  --cafile $ORDERER_CA \
  -C $CHANNEL_NAME -n $CC_NAME \
  --peerAddresses localhost:7051 --tlsRootCertFiles $PEER0_ORG1_CA \
  --peerAddresses localhost:9051 --tlsRootCertFiles $PEER0_ORG2_CA \
  -c '{"function": "CreateEmployee", "Args":["'"$EMP_ID"'", "'"$NAME"'", "'"$AGE"'", "'"$EMPLOYER"'", "'"$CONTRIBUTION_AMOUNT"'"]}'
