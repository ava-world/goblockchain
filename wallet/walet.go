package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	// "fmt"
)

type Wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockchainAddress string
}

func NewWallet() *Wallet {
	// 1. creating ecdsa privatekey
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey

	// 2. compute the EC public key by doing a secp256k1 point multiply (g * priv).
	h2 := sha256.New()
	h2.Write(W.publicKey.X.Bytes())
	h2.Write(W.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// 3. optionally compress the public key to 33 bytes (02/03 + x).

	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// 4. SHA-256 the public key.
	// 5. RIPEMD-160 the SHA-256 output → pubkey hash (20 bytes).
	// 6. prepend the network version byte (0x00 for mainnet, 0x6f for testnet).
	// 7. compute checksum = first 4 bytes of SHA256(SHA256(version + pubkeyhash)).
	// 8. append the 4-byte checksum to (version + pubkeyhash).
	// 9. Base58Check-encode that 25-byte payload → final Bitcoin address.

	return w
}

// func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
// 	return w.privateKey
// }
// func (w *Wallet) PrivateKeyStr() string {
// 	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
// }

// func (w *Wallet) PublicKey() *ecdsa.PublicKey {
// 	return w.publicKey
// }
// func (w *Wallet) PublicKeyStr() string {
// 	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
// }
