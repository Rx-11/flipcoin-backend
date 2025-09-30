package config

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/Rx-11/flipcoin-backend/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AppContext struct {
	Client          *ethclient.Client
	Contract        *contracts.Contracts
	ChainID         *big.Int
	SignerKey       *ecdsa.PrivateKey
	SignerAddress   common.Address
	ContractAddress common.Address
}

var appCtx *AppContext

func InitContract() {
	log.Println("Connecting to Ethereum node...")
	dialCtx, dialCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer dialCancel()
	client, err := ethclient.DialContext(dialCtx, config.RPC_URL)
	if err != nil {
		log.Fatalf("FATAL: Failed to connect to Ethereum client at %s: %v", config.RPC_URL, err)
	}
	log.Println("Successfully connected to Ethereum node.")

	config.PrivateKeyHex = strings.TrimPrefix(config.PrivateKeyHex, "0x")

	privateKey, err := crypto.HexToECDSA(config.PrivateKeyHex)
	if err != nil {
		log.Fatalf("FATAL: Failed to parse private key: %v", err)
	}
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("FATAL: Error casting public key to ECDSA type.")
	}

	signerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Printf("Backend signer address loaded: %s", signerAddress.Hex())

	networkCtx, networkCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer networkCancel()
	chainID, err := client.NetworkID(networkCtx)
	if err != nil {
		log.Fatalf("FATAL: Failed to get chain ID from the node: %v", err)
	}
	log.Printf("Operating on Chain ID: %s", chainID.String())

	contractAddr := common.HexToAddress(config.ContractAddress)
	instance, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		log.Fatalf("FATAL: Failed to instantiate contract binding at %s: %v", contractAddr.Hex(), err)
	}
	log.Printf("Contract instance created for address: %s", contractAddr.Hex())

	appCtx = &AppContext{
		Client:          client,
		Contract:        instance,
		ChainID:         chainID,
		SignerKey:       privateKey,
		SignerAddress:   signerAddress,
		ContractAddress: contractAddr,
	}
}

func GetContract() *AppContext {
	if appCtx == nil {
		InitContract()
	}
	return appCtx
}
