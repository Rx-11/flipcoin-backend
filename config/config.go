package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RPC_URL         string
	ContractAddress string
	OracleAddress   string
	OwnerAddress    string
	OperatorAddress string
	PrivateKeyHex   string
}

var config Config

func Init() {
	_ = godotenv.Load()
	config.RPC_URL = os.Getenv("RPC_URL")
	config.ContractAddress = os.Getenv("CONTRACT_ADDRESS")
	config.OracleAddress = os.Getenv("ORACLE_ADDRESS")
	config.OwnerAddress = os.Getenv("OWNER_ADDRESS")
	config.OperatorAddress = os.Getenv("OPERATOR_ADDRESS")
	config.PrivateKeyHex = os.Getenv("PRIVATE_KEY_HEX")
}

func GetConfig() Config {
	return config
}
