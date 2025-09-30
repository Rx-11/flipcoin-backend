package public

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/Rx-11/flipcoin-backend/config"
	"github.com/Rx-11/flipcoin-backend/ws"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/core/types"
)

func RoundManager(wsBroadcaster *ws.WsBroadcaster) {
	appCtx := config.GetContract()
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	var currentEpoch *big.Int
	var prevEpoch *big.Int
	var intervalSecs *big.Int
	var err error

	prevEpoch = big.NewInt(-1)

	intervalSecs, err = appCtx.Contract.IntervalSeconds(&bind.CallOpts{Context: ctx, Pending: false})
	if err != nil {
		log.Printf("WARN [RoundManager]: Could not determine interval secs: %v.", err)
		intervalSecs = big.NewInt(300)
	}

	currentEpoch, err = appCtx.Contract.CurrentEpoch(&bind.CallOpts{Context: ctx, Pending: false})
	if err != nil {
		log.Printf("WARN [RoundManager]: Could not determine initial round state: %v. Will attempt to start epoch 0 or next logical one.", err)
		currentEpoch = big.NewInt(-1)
		prevEpoch = big.NewInt(-1)
	}

	log.Println("Current Epoch:" + currentEpoch.String())

	if currentEpoch.Cmp(big.NewInt(-1)) == 0 {
		_, currentEpoch, err = startRound()
		if err != nil {
			log.Printf("ERROR [RoundManager]: Failed to start round: %v", err)
			currentEpoch = big.NewInt(-1)
			prevEpoch = big.NewInt(-1)
		}
	}

	prevEpoch = prevEpoch.Sub(currentEpoch, big.NewInt(1))

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			round, err := fetchRound(currentEpoch.Int64())
			if err != nil {
				log.Printf("ERROR [RoundManager]: Failed to fetch round: %v", err)
			}

			var prevround struct {
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
			}
			if prevEpoch.Cmp(big.NewInt(-1)) != 0 {
				prevround, err = fetchRound(prevEpoch.Int64())
				if err != nil {
					log.Printf("ERROR [RoundManager]: Failed to fetch round: %v", err)
				}
			}

			var tx *types.Transaction

			if round.StartTimestamp.Int64() == 0 {
				tx, currentEpoch, err = startRound()
				if err != nil {
					log.Printf("ERROR [RoundManager]: Failed to start round: %v", err)
				} else {
					log.Printf("SUCCESS [RoundManager]: startRound TX submitted for epoch %s: %s", currentEpoch.String(), tx.Hash().Hex())
					wsBroadcaster.Broadcast("round_start", map[string]interface{}{
						"epoch":           currentEpoch.String(),
						"transactionHash": tx.Hash().Hex(),
					})
					// time.Sleep(5 * time.Second)
				}
			} else if round.StartTimestamp.Int64() > 0 && round.LockTimestamp.Int64() < time.Now().Unix() && round.LockPrice.Int64() == 0 {
				timeSinceStart := time.Since(time.Unix(round.StartTimestamp.Int64(), 0))
				if timeSinceStart.Seconds() >= float64(intervalSecs.Int64()) {
					tx, prevEpoch, err = lockRound(currentEpoch.Int64())
					if err != nil {
						log.Printf("ERROR [RoundManager]: Failed to lock round: %v", err)
					} else {
						log.Printf("SUCCESS [RoundManager]: lockRound TX submitted for epoch %s: %s", prevEpoch.String(), tx.Hash().Hex())
						wsBroadcaster.Broadcast("round_locked", map[string]interface{}{
							"epoch":           prevEpoch.String(),
							"transactionHash": tx.Hash().Hex(),
						})
						time.Sleep(1 * time.Second)
					}

					tx, currentEpoch, err = startRound()
					if err != nil {
						log.Printf("ERROR [RoundManager]: Failed to start round: %v", err)
					} else {
						log.Printf("SUCCESS [RoundManager]: startRound TX submitted for epoch %s: %s", currentEpoch.String(), tx.Hash().Hex())
						wsBroadcaster.Broadcast("round_start", map[string]interface{}{
							"epoch":           currentEpoch.String(),
							"transactionHash": tx.Hash().Hex(),
						})
						// time.Sleep(5 * time.Second)
					}
				}
			} else if round.StartTimestamp.Int64() > 0 && round.LockTimestamp.Int64() > 0 && round.CloseTimestamp.Int64() < time.Now().Unix() && round.ClosePrice.Int64() == 0 {
				timeSinceLock := time.Since(time.Unix(round.LockTimestamp.Int64(), 0))
				if timeSinceLock.Seconds() >= float64(intervalSecs.Int64()) {
					tx, currentEpoch, err = closeRound(currentEpoch.Int64())
					if err != nil {
						log.Printf("ERROR [RoundManager]: Failed to close round: %v", err)
					} else {
						log.Printf("SUCCESS [RoundManager]: closeRound TX submitted for epoch %s: %s", currentEpoch.String(), tx.Hash().Hex())
						wsBroadcaster.Broadcast("round_close", map[string]interface{}{
							"epoch":           currentEpoch.String(),
							"transactionHash": tx.Hash().Hex(),
						})
						time.Sleep(1 * time.Second)
					}

				}
			} else if round.ClosePrice.Int64() > 0 {
				tx, currentEpoch, err = startRound()
				if err != nil {
					log.Printf("ERROR [RoundManager]: Failed to start round: %v", err)
				} else {
					log.Printf("SUCCESS [RoundManager]: startRound TX submitted for epoch %s: %s", currentEpoch.String(), tx.Hash().Hex())
					wsBroadcaster.Broadcast("round_start", map[string]interface{}{
						"epoch":           currentEpoch.String(),
						"transactionHash": tx.Hash().Hex(),
					})
					time.Sleep(5 * time.Second)
				}
			} else if prevround.StartTimestamp.Int64() > 0 && prevround.LockTimestamp.Int64() > 0 && prevround.CloseTimestamp.Int64() < time.Now().Unix() && prevround.ClosePrice.Int64() == 0 {
				timeSinceLock := time.Since(time.Unix(prevround.LockTimestamp.Int64(), 0))
				if timeSinceLock.Seconds() >= float64(intervalSecs.Int64()) {
					tx, prevEpoch, err = closeRound(prevEpoch.Int64())
					if err != nil {
						log.Printf("ERROR [RoundManager]: Failed to close round: %v", err)
					} else {
						log.Printf("SUCCESS [RoundManager]: closeRound TX submitted for epoch %s: %s", prevEpoch.String(), tx.Hash().Hex())
						wsBroadcaster.Broadcast("round_close", map[string]interface{}{
							"epoch":           prevEpoch.String(),
							"transactionHash": tx.Hash().Hex(),
						})
						// time.Sleep(1 * time.Second)
					}

				}
			} else if prevround.ClosePrice.Int64() > 0 {
				prevEpoch = currentEpoch
			}
		}
	}

}
