echo "Updating chaincode.."

echo "Installing vlm chaincode to peer0.manf.vlm.com..."
# install chaincode
# Install code on manf peer
docker exec -e "CORE_PEER_LOCALMSPID=manfMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manf.vlm.com/users/Admin@manf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manf.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.manf.vlm.com "


echo "Installing vlm chaincode to peer0.rta.vlm.com...."
# Install code on rta peer
docker exec -e "CORE_PEER_LOCALMSPID=rtaMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rta.vlm.com/users/Admin@rta.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.rta.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.rta.vlm.com"

echo "Installing vlm chaincode to peer0.dlr.vlm.com..."
# Install code on dlr peer
docker exec -e "CORE_PEER_LOCALMSPID=dlrMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/dlr.vlm.com/users/Admin@dlr.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.dlr.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.dlr.vlm.com"

echo "Installing vlm chaincode to peer0.scr.vlm.com..."
# Install code on scr peer
docker exec -e "CORE_PEER_LOCALMSPID=scrMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/scr.vlm.com/users/Admin@scr.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.scr.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.scr.vlm.com"

echo "Installing vlm chaincode to peer0.bank.vlm.com..."
# Install code on bank peer
docker exec -e "CORE_PEER_LOCALMSPID=bankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.vlm.com/users/Admin@bank.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.bank.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.bank.vlm.com"

echo "Installing vlm chaincode to peer0.insu.vlm.com..."
# Install code on insu peer
docker exec -e "CORE_PEER_LOCALMSPID=insuMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insu.vlm.com/users/Admin@insu.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.insu.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.insu.vlm.com"

echo "Installing vlm chaincode to peer0.tarf.vlm.com..."
# Install code on insu peer
docker exec -e "CORE_PEER_LOCALMSPID=trafMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/traf.vlm.com/users/Admin@traf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.traf.vlm.com:7051" cli peer chaincode install -n vlmcc -v 1.3 -p github.com/vlm/go -l golang
sleep 5
echo "Installed vlm chaincode to peer0.traf.vlm.com"


echo "Instantiating vlm chaincode.."
docker exec -e "CORE_PEER_LOCALMSPID=manfMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manf.vlm.com/users/Admin@manf.vlm.com/msp" -e "CORE_PEER_ADDRESS=peer0.manf.vlm.com:7051" cli peer chaincode upgrade -o orderer.vlm.com:7050 -C vlmchannel -n vlmcc -l golang -v 1.3 -c '{"Args":[""]}' -P "OR ('manfMSP.member','rtaMSP.member','dlrMSP.member','scrMSP.member','bankMSP.member', 'insuMSP.member', 'trafMSP.member')"
echo "Instantiated vlm chaincode."
echo "Following is the docker network....."

docker ps

