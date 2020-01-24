echo "Setting up the network.."

echo "Creating channel genesis block.."
# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=manfMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manf.vlm.com/users/Admin@manf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manf.vlm.com:7051" cli peer channel create -o orderer.vlm.com:7050 -c vlmchannel -f /etc/hyperledger/configtx/vlmchannel.tx
sleep 5
echo "Channel genesis block created."

echo "peer0.manf.vlm.com joining the channel..."
# Join peer0.manf.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=manfMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manf.vlm.com/users/Admin@manf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manf.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.manf.vlm.com joined the channel"

echo "peer0.rta.vlm.com joining the channel..."
# Join peer0.rta.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=rtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.rta.vlm.com joined the channel"

echo "peer0.dlr.vlm.com joining the channel..."
# Join peer0.dlr.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=dlrMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dlr.vlm.com/users/Admin@dlr.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.dlr.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.dlr.vlm.com joined the channel"

echo "peer0.scr.vlm.com joining the channel..."
# Join peer0.scr.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=scrMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/scr.vlm.com/users/Admin@scr.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.scr.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.scr.vlm.com joined the channel"

echo "peer0.bank.vlm.com joining the channel..."
# Join peer0.bank.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=bankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.vlm.com/users/Admin@bank.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.bank.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.bank.vlm.com joined the channel"

echo "peer0.insu.vlm.com joining the channel..."
# Join peer0.insu.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=insuMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insu.vlm.com/users/Admin@insu.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.insu.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.insu.vlm.com joined the channel"

echo "peer0.traf.vlm.com joining the channel..."
# Join peer0.traf.vlm.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=insuMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/traf.vlm.com/users/Admin@traf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.traf.vlm.com:7051" cli peer channel join -b vlmchannel.block
sleep 5
echo "peer0.traf.vlm.com joined the channel"


echo "Installing vlm chaincode to peer0.manf.vlm.com..."
# install chaincode
# Install code on manf peer
docker exec -e "CORE_PEER_LOCALMSPID=manfMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manf.vlm.com/users/Admin@manf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manf.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.manf.vlm.com "


echo "Installing vlm chaincode to peer0.rta.vlm.com...."
# Install code on rta peer
docker exec -e "CORE_PEER_LOCALMSPID=rtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.rta.vlm.com"

echo "Installing vlm chaincode to peer0.dlr.vlm.com..."
# Install code on dlr peer
docker exec -e "CORE_PEER_LOCALMSPID=dlrMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dlr.vlm.com/users/Admin@dlr.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.dlr.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.dlr.vlm.com"

echo "Installing vlm chaincode to peer0.scr.vlm.com..."
# Install code on scr peer
docker exec -e "CORE_PEER_LOCALMSPID=scrMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/scr.vlm.com/users/Admin@scr.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.scr.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.scr.vlm.com"

echo "Installing vlm chaincode to peer0.bank.vlm.com..."
# Install code on bank peer
docker exec -e "CORE_PEER_LOCALMSPID=bankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.vlm.com/users/Admin@bank.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.bank.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.bank.vlm.com"

echo "Installing vlm chaincode to peer0.insu.vlm.com..."
# Install code on insu peer
docker exec -e "CORE_PEER_LOCALMSPID=insuMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insu.vlm.com/users/Admin@insu.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.insu.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.insu.vlm.com"

echo "Installing vlm chaincode to peer0.tarf.vlm.com..."
# Install code on insu peer
docker exec -e "CORE_PEER_LOCALMSPID=trafMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/traf.vlm.com/users/Admin@traf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.traf.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.0 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.traf.vlm.com"


echo "Instantiating vlm chaincode.."
docker exec -e "CORE_PEER_LOCALMSPID=manfMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manf.vlm.com/users/Admin@manf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manf.vlm.com:7051" cli peer chaincode instantiate -o orderer.vlm.com:7050 -C vlmchannel -n vlmcc -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('manfMSP.member','rtaMSP.member','dlrMSP.member','scrMSP.member','bankMSP.member', 'insuMSP.member', 'trafMSP.member')"
echo "Instantiated vlm chaincode."
echo "Following is the docker network....."

docker ps

