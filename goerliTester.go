package main

import (
	"fmt"
	"log"

	"./contracts"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func initViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func getClient(host string) (theClient *ethclient.Client, err error) {
	theClient, err = ethclient.Dial(host)
	return
}

func runTest(host string) bool {
	client, err := getClient(host)
	if err != nil {
		log.Println(host, err.Error())
		return false
	}
	defer client.Close()
	address := common.HexToAddress("0x4E10A95f0bb2fEc6ec1c4296a16420a018a5F9Fe")
	light, err := contracts.NewLighthouse(address, client)
	if err != nil {
		log.Println(host, err.Error())
		return false
	}
	_, err = light.PeekData(nil)
	if err != nil {
		log.Println(host, err.Error())
		return false
	}

	return true
}

func main() {
	initViper()
	logName := "rpc.log"
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/" + logName,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	})

	rpcHostArray := []string{"GOERLI_1", "GOERLI_2", "GOERLI_3", "GOERLI_4"}
	rpcArray := []string{}
	for _, rpc := range rpcHostArray {
		rpcArray = append(rpcArray, viper.GetString(rpc))
	}
	scores := []int{0, 0, 0, 0}
	for pass := 0; pass < 100; pass++ {
		fmt.Print("run # ", pass)
		for index, rpc := range rpcArray {
			if runTest(rpc) {
				scores[index]++
			}
		}
		fmt.Println(scores)
	}

}
