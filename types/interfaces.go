package types

import (
	"context"
	"math/big"

	"github.com/0xPolygon/agglayer/tx"
	"github.com/0xPolygonHermez/zkevm-node/ethtxmanager"
	"github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

type IDB interface {
	BeginStateTransaction(ctx context.Context) (pgx.Tx, error)
}

type IEtherman interface {
	GetSequencerAddr(rollupId uint32) (common.Address, error)
	BuildTrustedVerifyBatchesTxData(lastVerifiedBatch, newVerifiedBatch uint64, proof tx.ZKP, rollupId uint32) (data []byte, err error)
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

type IEthTxManager interface {
	Add(ctx context.Context, owner, id string, from common.Address, to *common.Address, value *big.Int, data []byte, gasOffset uint64, dbTx pgx.Tx) error
	Result(ctx context.Context, owner, id string, dbTx pgx.Tx) (ethtxmanager.MonitoredTxResult, error)
	ResultsByStatus(ctx context.Context, owner string, statuses []ethtxmanager.MonitoredTxStatus, dbTx pgx.Tx) ([]ethtxmanager.MonitoredTxResult, error)
	ProcessPendingMonitoredTxs(ctx context.Context, owner string, failedResultHandler ethtxmanager.ResultHandler, dbTx pgx.Tx)
}

type IZkEVMClient interface {
	BatchByNumber(ctx context.Context, number *big.Int) (*types.Batch, error)
}

type IZkEVMClientClientCreator interface {
	NewClient(rpc string) IZkEVMClient
}
