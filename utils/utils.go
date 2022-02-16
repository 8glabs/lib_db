package utils

import (
	"crypto/ecdsa"
	"log"

	"github.com/8glabs/lib_db/models"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
)

func GenerateChainWallet() *models.ChainWallet {
	solanaPrivateKeyStr := generateSolanaKeypair()
	privateKey, publicKey, ethAddress := generateEthereumWallet()
	// TDOO(weiduan):Generate starkex public key and private key on the fly.
	chainWallet := models.ChainWallet{
		CustodialStarkexNextVaultId:    7263546,
		CustodialStarkexPublicKey:      "0x2f116d013fb6ecae90765a876a5bfcf66cd6a6be1f85c9841629cd0bd080ed3",
		CustodialStarkexPrivateKey:     "0x1be737e23fe35e5ecc0eccd5f1222b1bddbbe55a360849bab91ebbfcfd9925f",
		CustodialEthereumWalletAddress: ethAddress,
		CustodialEthereumPublicKey:     publicKey,
		CustodialEthereumPrivateKey:    privateKey,
		CustodialSolanaPrivateKeyStr:   solanaPrivateKeyStr,
	}
	return &chainWallet
}

func generateSolanaKeypair() string {
	account := solana.NewWallet()
	return account.PrivateKey.String()
}

func generateEthereumWallet() (string, string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// TODO(weiduan): Add mnemonic
	return hexutil.Encode(privateKeyBytes), hexutil.Encode(publicKeyBytes), crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}
