package main

import (
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	//Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress() "B", 1.0)

	//Blockchain
	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAdress(), walletB.BlockchainAddress(), 1.0, 
		walletA.PublicKey(), t.GenerateSignature())
	fmt.Println("Added? ", isAdded)	

	blockchain.Mining.()
	blockchain.Printf()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
}
