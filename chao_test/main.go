package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/harmony-one/harmony/core/rawdb"
)

func main() {
	folder := "/home/chao/go/src/github.com/harmony-one/harmony/test/db/"
	subfolder := "harmony_127.0.0.1_9001"
	dbPath := folder + subfolder
	db, err := ethdb.NewLDBDatabase(dbPath, 128, 128)
	defer db.Close()
	if err != nil {
		fmt.Println("ops, cannot open database")
		return
	}
	hash := rawdb.ReadHeadBlockHash(db)
	number := rawdb.ReadHeaderNumber(db, hash)
	fmt.Println("number:", *number)
	//shash := common.HexToHash("0x401948297e660dbba2fef8cd1fcd7e2167684be835f3181943e8dc58281790ec")
	shash := common.HexToHash("ad2a217b15f784348dde8d5b780ce578e9a9a45dd6608c7918f5cb0a5a70acc3")
	for i := 0; i < 230; i++ {
		num := uint64(i)
		hash := rawdb.ReadCanonicalHash(db, num)
		block := rawdb.ReadBlock(db, hash, num)
		if block.Root() == shash {
			fmt.Println("fount it: ", i)
		}
	}
	//fmt.Printf("db=%v, blockNum=%v, blockRoot=%v\n", subfolder, block.Number(), block.Root().Hex())
	fmt.Println(shash.Hex())

	stateCache := state.NewDatabase(db)
	sdb := trie.NewDatabase(db)
	tr, err := trie.NewSecure(shash, sdb, 1000)
	fmt.Println("tr is: ", tr, "err is: ", err)

	ss, err := state.New(shash, stateCache)
	fmt.Println("state is ", ss)
}
