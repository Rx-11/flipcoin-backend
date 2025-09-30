package public

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/Rx-11/flipcoin-backend/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func fetchRound(epoch int64) (struct {
	Epoch          *big.Int
	StartTimestamp *big.Int
	LockTimestamp  *big.Int
	CloseTimestamp *big.Int
	LockPrice      *big.Int
	ClosePrice     *big.Int
	TotalAmount    *big.Int
	BullAmount     *big.Int
	BearAmount     *big.Int
	OracleCalled   bool
}, error) {
	log.Println("Attempting to fetch round...")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	tx, err := config.GetContract().Contract.Rounds(&bind.CallOpts{Context: ctx, Pending: false}, big.NewInt(epoch))
	if err != nil {
		log.Printf("ERROR [fetchRoundInternal]: Contract call fetchRound failed: %v", err)
		return tx, err
	}
	log.Println(tx)
	return tx, nil

}

func startRound() (*types.Transaction, *big.Int, error) {
	log.Println("Attempting to start new round...")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	txOpts, err := createTxOpts(200000, 0, ctx)
	if err != nil {
		log.Fatalf("Failed to create txOpts: %v", err)
		return nil, nil, err
	}

	currentEpoch, err := config.GetContract().Contract.CurrentEpoch(&bind.CallOpts{Context: ctx})
	var expectedNextEpoch *big.Int = big.NewInt(-1) // Initialize to indicate potential issue
	if err != nil {
		log.Printf("WARN [startRoundInternal]: Failed to get current epoch before starting round: %v", err)
	} else {
		expectedNextEpoch = new(big.Int).Add(currentEpoch, big.NewInt(1))
		log.Printf("INFO [startRoundInternal]: Current epoch is %s, attempting to start epoch %s", currentEpoch.String(), expectedNextEpoch.String())
	}

	tx, err := config.GetContract().Contract.StartRound(txOpts)
	if err != nil {
		log.Fatalf("ERROR [startRoundInternal]: Contract call StartRound failed: %v", err)
		return nil, nil, err
	}
	receipt, err := bind.WaitMined(ctx, config.GetContract().Client, tx)
	if err != nil {
		log.Printf("ERROR [startRoundInternal]: Failed waiting for start tx %s receipt: %v", tx.Hash().Hex(), err)
		return nil, nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatalf("ERROR [handleStartRound]: Start transaction %s failed (reverted)", tx.Hash().Hex())
	}
	return tx, expectedNextEpoch, nil

}

func lockRound(epoch int64) (*types.Transaction, *big.Int, error) {
	log.Println("Attempting to lock round...")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	txOpts, err := createTxOpts(200000, 0, ctx)
	if err != nil {
		log.Fatalf("Failed to create txOpts: %v", err)
		return nil, nil, err
	}

	tx, err := config.GetContract().Contract.LockRound(txOpts, big.NewInt(epoch))
	if err != nil {
		log.Fatalf("ERROR [lockRoundInternal]: Contract call lockRound failed: %v", err)
		return nil, nil, err
	}
	receipt, err := bind.WaitMined(ctx, config.GetContract().Client, tx)
	if err != nil {
		log.Printf("ERROR [lockRoundInternal]: Failed waiting for lock tx %s receipt: %v", tx.Hash().Hex(), err)
		return nil, nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatalf("ERROR [lockRoundInternal]: lock transaction %s failed (reverted)", tx.Hash().Hex())
	}
	return tx, big.NewInt(epoch), nil

}

func closeRound(epoch int64) (*types.Transaction, *big.Int, error) {
	log.Println("Attempting to close round...")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	txOpts, err := createTxOpts(200000, 0, ctx)
	if err != nil {
		log.Fatalf("Failed to create txOpts: %v", err)
		return nil, nil, err
	}

	tx, err := config.GetContract().Contract.CloseRound(txOpts, big.NewInt(epoch))
	if err != nil {
		log.Fatalf("ERROR [closeRoundInternal]: Contract call closeRound failed: %v", err)
		return nil, nil, err
	}
	receipt, err := bind.WaitMined(ctx, config.GetContract().Client, tx)
	if err != nil {
		log.Printf("ERROR [closeRoundInternal]: Failed waiting for close tx %s receipt: %v", tx.Hash().Hex(), err)
		return nil, nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatalf("ERROR [closeRoundInternal]: close transaction %s failed (reverted)", tx.Hash().Hex())
	}
	return tx, big.NewInt(epoch), nil

}

func createTxOpts(gasLimit uint64, value int64, ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(config.GetContract().SignerKey, (config.GetContract().ChainID))
	if err != nil {
		log.Printf("Failed to create transactor: %v", err)
		return nil, err
	}

	nonce, err := config.GetContract().Client.PendingNonceAt(ctx, config.GetContract().SignerAddress)
	if err != nil {
		log.Printf("Failed to get nonce: %v", err)
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))

	gasPrice, err := config.GetContract().Client.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("Failed to suggest gas price: %v", err)
		return nil, err
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = gasLimit
	auth.Value = big.NewInt(value)
	auth.Context = ctx

	return auth, nil
}

func fetchPrice() (*big.Int, error) {
	log.Println("Attempting to fetch price...")
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	price, err := config.GetContract().Contract.GetLatestPrice(&bind.CallOpts{Context: ctx, Pending: false})
	if err != nil {
		log.Printf("ERROR [fetchPriceInternal]: Contract call fetchPrice failed: %v", err)
		return nil, err
	}
	return price, nil

}
