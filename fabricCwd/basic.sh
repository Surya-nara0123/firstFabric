./fabricShit.sh down

echo ""
echo ""
echo ""
echo ""
echo ""

./fabricShit.sh up createChannel

echo ""
echo ""
echo ""
echo ""
echo ""

./fabricShit.sh deployCC -c mychannel -ccn chaincode1 -ccl go -ccs auto -ccp ./chaincode

echo ""
echo ""
echo ""
echo ""
echo ""

docker ps -a 

echo ""
echo ""
echo ""
echo ""
echo ""

./fabricShit.sh cc invoke -c mychannel -ccn chaincode1 -ccic '{"Args":["InitLedger"]}'