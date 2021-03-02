package lwpow

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

//LWPoW is a lightweight proof-of-work protocol that uses cryptography algorithms to achieve its results
type LWPoW struct {
	config LwpowConfig    // Consensus engine configuration parameters
	db     ethdb.Database // Database to store and retrieve snapshot checkpoints
}

// Author implements consensus.Engine, returning the header's coinbase as the
// proof-of-work verified author of the block.
func (lwpow *LWPoW) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}


// VerifyHeader checks whether a header conforms to the consensus rules of the
// consensus engine
func (lwpow *LWPoW) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {

	log.Info("will verfiyHeader")
	return nil
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications.
func (lwpow *LWPoW) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {

	log.Info("will verfiyHeaders")
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {

		for _, header := range headers {
			err := lwpow.VerifyHeader(chain, header, false)

			select {
				case <-abort:
					return

				case results <- err:
			}
		}
	}()

	return abort, results
}

func (lwpow *LWPoW) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {

	log.Info("will verfiy uncles")
	return nil
}

func (lwpow *LWPoW) VerifySeal(chain consensus.ChainReader, header *types.Header) error {

	log.Info("will verfiy VerifySeal")
	return nil
}

func (lwpow *LWPoW) Prepare(chain consensus.ChainReader, header *types.Header) error {

	log.Info("will prepare the block")

	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)

	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	header.Difficulty = LWPoW.CalcDifficulty(chain, header.Time.Uint64(), parent)

	return nil
}

func (lwpow *LWPoW) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {

	return calcDifficultyHomestead(time, parent)
}

func (lwpow *LWPoW) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, 
	uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {

	log.Info("will Finalize the block")

	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	b := types.NewBlock(header, txs, uncles, receipts)

	return b, nil
}

func (lwpow *LWPoW) Seal(chain consensus.ChainReader, block *types.Block, stop <-chan struct{}) (*types.Block, error) {

	log.Info("will Seal the block")

	//time.Sleep(15 * time.Second)

	header := block.Header()

	header.Nonce, header.MixDigest = getRequiredHeader()

	return block.WithSeal(header), nil
}
