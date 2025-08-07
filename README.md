# Completed Commands
```bash
Input
./create-artifacts.sh 

Output
chmod: cannot access './crypto-config': No such file or directory
org1.example.com
org2.example.com
mychannel
2025-07-27 13:33:31.088 +06 0001 INFO [common.tools.configtxgen] main -> Loading configuration
2025-07-27 13:33:31.093 +06 0002 INFO [common.tools.configtxgen.localconfig] completeInitialization -> orderer type: etcdraft
2025-07-27 13:33:31.093 +06 0003 INFO [common.tools.configtxgen.localconfig] completeInitialization -> Orderer.EtcdRaft.Options unset, setting to tick_interval:"500ms" election_tick:10 heartbeat_tick:1 max_inflight_blocks:5 snapshot_interval_size:16777216 
2025-07-27 13:33:31.093 +06 0004 INFO [common.tools.configtxgen.localconfig] Load -> Loaded configuration: configtx.yaml
2025-07-27 13:33:31.097 +06 0005 INFO [common.tools.configtxgen] doOutputBlock -> Generating genesis block
2025-07-27 13:33:31.097 +06 0006 INFO [common.tools.configtxgen] doOutputBlock -> Creating system channel genesis block
2025-07-27 13:33:31.098 +06 0007 INFO [common.tools.configtxgen] doOutputBlock -> Writing genesis block
2025-07-27 13:33:31.127 +06 0001 INFO [common.tools.configtxgen] main -> Loading configuration
2025-07-27 13:33:31.134 +06 0002 INFO [common.tools.configtxgen.localconfig] Load -> Loaded configuration: configtx.yaml
2025-07-27 13:33:31.134 +06 0003 INFO [common.tools.configtxgen] doOutputChannelCreateTx -> Generating new channel configtx
2025-07-27 13:33:31.136 +06 0004 INFO [common.tools.configtxgen] doOutputChannelCreateTx -> Writing new channel tx
#######    Generating anchor peer update for Org1MSP  ##########
2025-07-27 13:33:31.162 +06 0001 INFO [common.tools.configtxgen] main -> Loading configuration
2025-07-27 13:33:31.167 +06 0002 INFO [common.tools.configtxgen.localconfig] Load -> Loaded configuration: configtx.yaml
2025-07-27 13:33:31.167 +06 0003 INFO [common.tools.configtxgen] doOutputAnchorPeersUpdate -> Generating anchor peer update
2025-07-27 13:33:31.168 +06 0004 INFO [common.tools.configtxgen] doOutputAnchorPeersUpdate -> Writing anchor peer update
#######    Generating anchor peer update for Org2MSP  ##########
2025-07-27 13:33:31.193 +06 0001 INFO [common.tools.configtxgen] main -> Loading configuration
2025-07-27 13:33:31.198 +06 0002 INFO [common.tools.configtxgen.localconfig] Load -> Loaded configuration: configtx.yaml
2025-07-27 13:33:31.198 +06 0003 INFO [common.tools.configtxgen] doOutputAnchorPeersUpdate -> Generating anchor peer update
2025-07-27 13:33:31.199 +06 0004 INFO [common.tools.configtxgen] doOutputAnchorPeersUpdate -> Writing anchor peer update
```
```bash
docker-compose -f ./artifacts/docker-compose.yaml up -d
# docker-compose -f ./artifacts/docker-compose.yaml down -v
# docker volume prune -f

[+] Running 14/14
 ✔ Network artifacts_test            Created                                                                      0.0s 
 ✔ Container couchdb2                Started                                                                      0.4s 
 ✔ Container ca.org2.example.com     Started                                                                      0.3s 
 ✔ Container orderer3.example.com    Started                                                                      0.4s 
 ✔ Container couchdb1                Started                                                                      0.5s 
 ✔ Container couchdb3                Started                                                                      0.5s 
 ✔ Container ca.org1.example.com     Started                                                                      0.3s 
 ✔ Container peer1.org1.example.com  Started                                                                      0.3s 
 ✔ Container orderer.example.com     Started                                                                      0.5s 
 ✔ Container orderer2.example.com    Started                                                                      0.4s 
 ✔ Container peer1.org2.example.com  Started                                                                      0.5s 
 ✔ Container couchdb0                Started                                                                      0.4s 
 ✔ Container peer0.org2.example.com  Started                                                                      0.5s 
 ✔ Container peer0.org1.example.com  Started  
```

```bash
sleep 5
./createChannel.sh

2025-07-27 19:58:01.060 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:01.068 +06 0002 INFO [cli.common] readBlock -> Expect block, but got status: &{NOT_FOUND}
2025-07-27 19:58:01.069 +06 0003 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:01.270 +06 0004 INFO [cli.common] readBlock -> Expect block, but got status: &{SERVICE_UNAVAILABLE}
2025-07-27 19:58:01.271 +06 0005 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:01.472 +06 0006 INFO [cli.common] readBlock -> Expect block, but got status: &{SERVICE_UNAVAILABLE}
2025-07-27 19:58:01.473 +06 0007 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:01.674 +06 0008 INFO [cli.common] readBlock -> Expect block, but got status: &{SERVICE_UNAVAILABLE}
2025-07-27 19:58:01.675 +06 0009 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:01.876 +06 000a INFO [cli.common] readBlock -> Expect block, but got status: &{SERVICE_UNAVAILABLE}
2025-07-27 19:58:01.877 +06 000b INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.078 +06 000c INFO [cli.common] readBlock -> Expect block, but got status: &{SERVICE_UNAVAILABLE}
2025-07-27 19:58:02.079 +06 000d INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.281 +06 000e INFO [cli.common] readBlock -> Received block: 0
2025-07-27 19:58:02.316 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.378 +06 0002 INFO [channelCmd] executeJoin -> Successfully submitted proposal to join channel
2025-07-27 19:58:02.410 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.471 +06 0002 INFO [channelCmd] executeJoin -> Successfully submitted proposal to join channel
2025-07-27 19:58:02.503 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.568 +06 0002 INFO [channelCmd] executeJoin -> Successfully submitted proposal to join channel
2025-07-27 19:58:02.600 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.654 +06 0002 INFO [channelCmd] executeJoin -> Successfully submitted proposal to join channel
2025-07-27 19:58:02.684 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.691 +06 0002 INFO [channelCmd] update -> Successfully submitted channel update
2025-07-27 19:58:02.721 +06 0001 INFO [channelCmd] InitCmdFactory -> Endorser and orderer connections initialized
2025-07-27 19:58:02.729 +06 0002 INFO [channelCmd] update -> Successfully submitted channel update

```
# Resting Command
```bash


sleep 2

./deployChaincode.sh
```