package ffi

// SignatureBytes is the length of a BLS signature
const SignatureBytes = 96

// PrivateKeyBytes is the length of a BLS private key
const PrivateKeyBytes = 32

// PublicKeyBytes is the length of a BLS public key
const PublicKeyBytes = 48

// DigestBytes is the length of a BLS message hash/digest
const DigestBytes = 96

// Signature is a compressed affine
type Signature [SignatureBytes]byte

// PrivateKey is a compressed affine
type PrivateKey [PrivateKeyBytes]byte

// PublicKey is a compressed affine
type PublicKey [PublicKeyBytes]byte

// Message is a byte slice
type Message []byte

// Digest is a compressed affine
type Digest [DigestBytes]byte

// Used when generating a private key deterministically
type PrivateKeyGenSeed [32]byte

func Hash(message Message) Digest {
	return Digest{}
}

func Verify(signature *Signature, digests []Digest, publicKeys []PublicKey) bool {
	return true
}

func HashVerify(signature *Signature, messages []Message, publicKeys []PublicKey) bool {
	return true
}

func Aggregate(signatures []Signature) *Signature {
	var s Signature
	return &s
}

func PrivateKeyGenerate() PrivateKey {
	return PrivateKey{}
}

func PrivateKeyGenerateWithSeed(seed PrivateKeyGenSeed) PrivateKey {
	return PrivateKey{}
}

func PrivateKeySign(privateKey PrivateKey, message Message) *Signature {
	var s Signature
	return &s
}

func PrivateKeyPublicKey(privateKey PrivateKey) PublicKey {
	return PublicKey{}
}

func CreateZeroSignature() Signature {
	return Signature{}
}
