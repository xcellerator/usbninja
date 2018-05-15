export GOOS=linux
export GOARCH=arm
export GOARM=6

./build_main.sh
./build_sendhid.sh
./build_stopall.sh
