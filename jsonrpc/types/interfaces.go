package types

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/pool"
	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/0xPolygonHermez/zkevm-node/state/runtime"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v4"
)

// PoolInterface contains the methods required to interact with the tx pool.
type PoolInterface interface {
	AddTx(ctx context.Context, tx types.Transaction, ip string) error
	GetGasPrices(ctx context.Context) (pool.GasPrices, error)
	GetNonce(ctx context.Context, address common.Address) (uint64, error)
	GetPendingTxHashesSince(ctx context.Context, since time.Time) ([]common.Hash, error)
	GetPendingTxs(ctx context.Context, limit uint64) ([]pool.Transaction, error)
	CountPendingTransactions(ctx context.Context) (uint64, error)
	GetTransactionByHash(ctx context.Context, hash common.Hash) (*pool.Transaction, error)
	GetTransactionByL2Hash(ctx context.Context, hash common.Hash) (*pool.Transaction, error)
	CheckPolicy(ctx context.Context, policy pool.PolicyName, address common.Address) (bool, error)
	CalculateEffectiveGasPrice(rawTx []byte, txGasPrice *big.Int, txGasUsed uint64, l1GasPrice uint64, l2GasPrice uint64) (*big.Int, error)
	EffectiveGasPriceEnabled() bool
}

// StateInterface gathers the methods required to interact with the state.
type StateInterface interface {
	StartToMonitorNewL2Blocks()
	BeginStateTransaction(ctx context.Context) (pgx.Tx, error)
	DebugTransaction(ctx context.Context, transactionHash common.Hash, traceConfig state.TraceConfig, dbTx pgx.Tx) (*runtime.ExecutionResult, error)
	EstimateGas(transaction *types.Transaction, senderAddress common.Address, l2BlockNumber *uint64, dbTx pgx.Tx) (uint64, []byte, error)
	GetBalance(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error)
	GetCode(ctx context.Context, address common.Address, root common.Hash) ([]byte, error)
	GetL2BlockByHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (*state.L2Block, error)
	GetL2BlockByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*state.L2Block, error)
	BatchNumberByL2BlockNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (uint64, error)
	GetL2BlockHashesSince(ctx context.Context, since time.Time, dbTx pgx.Tx) ([]common.Hash, error)
	GetL2BlockHeaderByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*state.L2Header, error)
	GetL2BlockTransactionCountByHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (uint64, error)
	GetL2BlockTransactionCountByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (uint64, error)
	GetLastVirtualizedL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetLastConsolidatedL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetLastL2Block(ctx context.Context, dbTx pgx.Tx) (*state.L2Block, error)
	GetLastL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetLogs(ctx context.Context, fromBlock uint64, toBlock uint64, addresses []common.Address, topics [][]common.Hash, blockHash *common.Hash, since *time.Time, dbTx pgx.Tx) ([]*types.Log, error)
	GetNonce(ctx context.Context, address common.Address, root common.Hash) (uint64, error)
	GetStorageAt(ctx context.Context, address common.Address, position *big.Int, root common.Hash) (*big.Int, error)
	GetSyncingInfo(ctx context.Context, dbTx pgx.Tx) (state.SyncingInfo, error)
	GetTransactionByHash(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*types.Transaction, error)
	GetTransactionByL2Hash(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*types.Transaction, error)
	GetTransactionByL2BlockHashAndIndex(ctx context.Context, blockHash common.Hash, index uint64, dbTx pgx.Tx) (*types.Transaction, error)
	GetTransactionByL2BlockNumberAndIndex(ctx context.Context, blockNumber uint64, index uint64, dbTx pgx.Tx) (*types.Transaction, error)
	GetTransactionReceipt(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*types.Receipt, error)
	IsL2BlockConsolidated(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (bool, error)
	IsL2BlockVirtualized(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (bool, error)
	ProcessUnsignedTransaction(ctx context.Context, tx *types.Transaction, senderAddress common.Address, l2BlockNumber *uint64, noZKEVMCounters bool, dbTx pgx.Tx) (*runtime.ExecutionResult, error)
	RegisterNewL2BlockEventHandler(h state.NewL2BlockEventHandler)
	GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetLastVerifiedBatch(ctx context.Context, dbTx pgx.Tx) (*state.VerifiedBatch, error)
	GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error)
	GetBatchL2DataByNumbers(ctx context.Context, batchNumbers []uint64, dbTx pgx.Tx) (map[uint64][]byte, error)
	GetForcedBatchDataByNumbers(ctx context.Context, batchNumbers []uint64, dbTx pgx.Tx) (map[uint64][]byte, error)
	GetTransactionsByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (txs []types.Transaction, effectivePercentages []uint8, err error)
	GetVirtualBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.VirtualBatch, error)
	GetVerifiedBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.VerifiedBatch, error)
	GetExitRootByGlobalExitRoot(ctx context.Context, ger common.Hash, dbTx pgx.Tx) (*state.GlobalExitRoot, error)
	GetL2BlocksByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]state.L2Block, error)
	GetNativeBlockHashesInRange(ctx context.Context, fromBlockNumber uint64, toBlockNumber uint64, dbTx pgx.Tx) ([]common.Hash, error)
	GetLastClosedBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error)
	GetLastVerifiedL2BlockNumberUntilL1Block(ctx context.Context, l1FinalizedBlockNumber uint64, dbTx pgx.Tx) (uint64, error)
	GetLastVerifiedBatchNumberUntilL1Block(ctx context.Context, l1BlockNumber uint64, dbTx pgx.Tx) (uint64, error)
	GetBatchTimestamp(ctx context.Context, batchNumber uint64, forcedForkId *uint64, dbTx pgx.Tx) (*time.Time, error)
	GetLatestBatchGlobalExitRoot(ctx context.Context, dbTx pgx.Tx) (common.Hash, error)
	GetL2TxHashByTxHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (*common.Hash, error)
	PreProcessUnsignedTransaction(ctx context.Context, tx *types.Transaction, sender common.Address, l2BlockNumber *uint64, dbTx pgx.Tx) (*state.ProcessBatchResponse, error)
}

// EthermanInterface provides integration with L1
type EthermanInterface interface {
	GetSafeBlockNumber(ctx context.Context) (uint64, error)
	GetFinalizedBlockNumber(ctx context.Context) (uint64, error)
}
