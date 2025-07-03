package blockchain

import (
	"math/big"

	"github.com/GoROSEN/rosen-apiserver/core/config"
)

type StellarChainAccess struct {
}

func NewStellarChainAccess(chainCfg *config.BlockchainConfig) (BlockChainAccess, error) {
	return nil, nil
}

func (bc *StellarChainAccess) NewWallet() (string, string) {
	return "", ""
}

func (bc *StellarChainAccess) FindTokenAccount(contractAddress string, walletAddress string) (string, error) {
	return "", nil
}

func (bc *StellarChainAccess) NewTokenAccount(contractAddress string, walletAddress string) (string, error) {
	return "", nil
}

func (bc *StellarChainAccess) QueryCoin(address string) (*big.Int, error) {
	return nil, nil
}

func (bc *StellarChainAccess) QueryToken(address string, contractAddress string) (*big.Int, error) {
	return nil, nil
}

func (bc *StellarChainAccess) TransferCoin(from, to string, value *big.Int) (string, error) {
	return "", nil
}

func (bc *StellarChainAccess) TransferToken(from, to string, value *big.Int, contractAddress string, decimal uint64) (string, error) {
	return "", nil
}

func (bc *StellarChainAccess) MintNFT(to string, contractAddress string, tokenId uint64, tokenUri string) (string, error) {
	return "", nil
}

func (bc *StellarChainAccess) TransferNFT(from, to string, contractAddress string, tokenId uint64) (string, error) {
	return "", nil
}

func (bc *StellarChainAccess) ConfirmTransaction(txhash string) (bool, error) {
	return false, nil
}
