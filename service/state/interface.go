package state

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/celestiaorg/nmt/namespace"
)

// Accessor represents the behaviors necessary for a user to
// query for state-related information and submit transactions/
// messages to the Celestia network.
type Accessor interface {
	// Start starts the state Accessor.
	Start(context.Context) error
	// Stop stops the state Accessor.
	Stop(context.Context) error

	AccountAddress(ctx context.Context) (Address, error)
	// Balance retrieves the Celestia coin balance for the node's account/signer
	// and verifies it against the corresponding block's AppHash.
	Balance(ctx context.Context) (*Balance, error)
	// BalanceForAddress retrieves the Celestia coin balance for the given address and verifies
	// the returned balance against the corresponding block's AppHash.
	//
	// NOTE: the balance returned is the balance reported by the block right before
	// the node's current head (head-1). This is due to the fact that for block N, the block's
	// `AppHash` is the result of applying the previous block's transaction list.
	BalanceForAddress(ctx context.Context, addr Address) (*Balance, error)

	// Transfer sends the given amount of coins from default wallet of the node to the given account address.
	Transfer(ctx context.Context, to AccAddress, amount math.Int, gasLimit uint64) (*TxResponse, error)
	// SubmitTx submits the given transaction/message to the
	// Celestia network and blocks until the tx is included in
	// a block.
	SubmitTx(ctx context.Context, tx Tx) (*TxResponse, error)
	// SubmitPayForData builds, signs and submits a PayForData transaction.
	SubmitPayForData(ctx context.Context, nID namespace.ID, data []byte, gasLim uint64) (*TxResponse, error)

	// CancelUnbondingDelegation cancels a user's pending undelegation from a validator.
	CancelUnbondingDelegation(
		ctx context.Context,
		valAddr ValAddress,
		amount,
		height Int,
		gasLim uint64,
	) (*TxResponse, error)
	// BeginRedelegate sends a user's delegated tokens to a new validator for redelegation.
	BeginRedelegate(ctx context.Context, srcValAddr, dstValAddr ValAddress, amount Int, gasLim uint64) (*TxResponse, error)
	// Undelegate undelegates a user's delegated tokens, unbonding them from the current validator.
	Undelegate(ctx context.Context, delAddr ValAddress, amount Int, gasLim uint64) (*TxResponse, error)
	// Delegate sends a user's liquid tokens to a validator for delegation.
	Delegate(ctx context.Context, delAddr ValAddress, amount Int, gasLim uint64) (*TxResponse, error)

	// QueryDelegation retrieves the delegation information between a delegator and a validator.
	QueryDelegation(ctx context.Context, valAddr ValAddress) (*types.QueryDelegationResponse, error)
	// QueryUnbonding retrieves the unbonding status between a delegator and a validator.
	QueryUnbonding(ctx context.Context, valAddr ValAddress) (*types.QueryUnbondingDelegationResponse, error)
	// QueryRedelegations retrieves the status of the redelegations between a delegator and a validator.
	QueryRedelegations(ctx context.Context, srcValAddr, dstValAddr ValAddress) (*types.QueryRedelegationsResponse, error)
}
