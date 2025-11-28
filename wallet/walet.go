package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
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
	h2.Write(w.publicKey.X.Bytes())
	h2.Write(w.publicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// 3. perform ripemd160 hashing.

	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// 4. add version byte in front of ripemd160.

	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[1:])

	// 5. perform sha256 hash on extended ripemd160 result

	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// 6. perform sha256 hash on result of prev sha 256 hash

	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h5.Sum(nil)

	// 7. compute checksum = first 4 bytes of SHA256(SHA256(version + pubkeyhash)).

	chsum := digest6[:4]
	// 8. append the 4-byte checksum to (version + pubkeyhash).

	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[:21], chsum[:])

	// 9. Base58Check-encode that 25-byte payload → final Bitcoin address.

	address := base58.Encode(dc8)
	w.blockchainAddress = address

	return w
}

func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}
func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}
func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}
func (w *Wallet) BlockchainAddress() string {
	return w.blockchainAddress
}
