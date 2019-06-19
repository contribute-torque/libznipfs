// package main implements libznipfs as an executable instead of a library
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/contribute-torque/libznipfs/src/ipfs"
	"github.com/contribute-torque/libznipfs/src/zeronet"
)

// main runs libznipfs and retrieves test data from ZeroNet and IPFS
func main() {

	fmt.Println("Test libznipfs as executable")

	baseDataPath := "/tmp"

	zn, err := zeronet.New(filepath.Join(baseDataPath, "zn-test"))
	if err != nil {
		fmt.Printf("Unable to create ZeroNet instance: %s\n", err)
		os.Exit(1)
	}

	// ZeroNet address 1KvWEyqhyHsU9y6UT8xYCFDC8Y1vKaNueX - Live JSON
	// ZeroNet address 1CH9ApTd83RM8ggz35pZApnKoZqDf7wXyh - Testnet JSON
	znAddress := "1KvWEyqhyHsU9y6UT8xYCFDC8Y1vKaNueX"

	fmt.Println("Fetching file from ZeroNet:", znAddress)
	// This is a well-known ZeroNet address. We store the IPFS hash in ipfs.hash
	content, err := zn.GetFile(znAddress, "ipfs.hash")
	if err != nil {
		fmt.Printf("Unable fetch from ZeroNet: %s\n", err)
		os.Exit(1)
	}

	// The content retrieved from ZeroNet is the IPFS hash containing the nodelist
	ipfsHash := strings.TrimSpace(string(content))
	fmt.Printf("Get IPFS hash from ZeroNet: %s\n", ipfsHash)

	ipfsNode, err := ipfs.New(baseDataPath)
	if err != nil {
		panic(err)
	}

	err = ipfsNode.Start()
	if err != nil {
		panic(err)
	}

	result, err := ipfsNode.Get(ipfsHash)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	err = ipfsNode.Stop()
	if err != nil {
		panic(err)
	}

}
