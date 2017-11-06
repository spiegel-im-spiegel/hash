package hash

import (
	"bytes"
	"crypto"
	"io"
	"strings"

	"github.com/pkg/errors"
)

var (
	//ErrNoImplement is error "no implementation"
	ErrNoImplement = errors.New("no implementation")
)

var algMap = map[crypto.Hash]string{
	crypto.MD4:         "md4",         //require "golang.org/x/crypto/md4" package
	crypto.MD5:         "md5",         //require "crypto/md5" package
	crypto.MD5SHA1:     "md5sha1",     //no implementation
	crypto.SHA1:        "sha1",        //require "crypto/sha1" package
	crypto.SHA224:      "sha224",      //require "crypto/sha256" package
	crypto.SHA256:      "sha256",      //require "crypto/sha256" package
	crypto.SHA384:      "sha384",      //require "crypto/sha512" package
	crypto.SHA512:      "sha512",      //require "crypto/sha512" package
	crypto.SHA512_224:  "sha512/224",  //require "crypto/sha512" package
	crypto.SHA512_256:  "sha512/256",  //require "crypto/sha512" package
	crypto.RIPEMD160:   "ripemd160",   //require "golang.org/x/crypto/ripemd160" package
	crypto.SHA3_224:    "sha3-224",    //require "golang.org/x/crypto/sha3" package
	crypto.SHA3_256:    "sha3-256",    //require "golang.org/x/crypto/sha3" package
	crypto.SHA3_384:    "sha3-384",    //require "golang.org/x/crypto/sha3" package
	crypto.SHA3_512:    "sha3-512",    //require "golang.org/x/crypto/sha3" package
	crypto.BLAKE2s_256: "blake2s",     //require "golang.org/x/crypto/blake2s" package
	crypto.BLAKE2b_256: "blake2b/256", //require "golang.org/x/crypto/blake2b" package
	crypto.BLAKE2b_384: "blake2b/384", //require "golang.org/x/crypto/blake2b" package
	crypto.BLAKE2b_512: "blake2b/512", //require "golang.org/x/crypto/blake2b" package
}
var algOrder = []crypto.Hash{
	crypto.MD4,
	crypto.MD5,
	crypto.MD5SHA1,
	crypto.SHA1,
	crypto.SHA224,
	crypto.SHA256,
	crypto.SHA384,
	crypto.SHA512,
	crypto.SHA512_224,
	crypto.SHA512_256,
	crypto.RIPEMD160,
	crypto.SHA3_224,
	crypto.SHA3_256,
	crypto.SHA3_384,
	crypto.SHA3_512,
	crypto.BLAKE2s_256,
	crypto.BLAKE2b_256,
	crypto.BLAKE2b_384,
	crypto.BLAKE2b_512,
}

//FuncList returns string of hash functions list
func FuncList() string {
	buf := new(bytes.Buffer)
	sep := ""
	for _, alg := range algOrder {
		if alg.Available() {
			if name, ok := algMap[alg]; ok {
				io.WriteString(buf, sep)
				io.WriteString(buf, name)
				sep = " "
			}
		}
	}
	return string(buf.Bytes())
}

//Algorithm returns crypto.Hash drom string
func Algorithm(s string) (crypto.Hash, error) {
	s = strings.ToLower(s)
	for alg, v := range algMap {
		if v == s {
			return alg, nil
		}
	}
	return crypto.Hash(0), errors.Wrap(ErrNoImplement, "error "+s)
}

//AlgoString returns string of hash algorithm
func AlgoString(alg crypto.Hash) string {
	if str, ok := algMap[alg]; ok {
		return str
	}
	return "unknown hash algorithm"
}

//Value returns hash value string from io.Reader
func Value(r io.Reader, alg crypto.Hash) ([]byte, error) {
	if !alg.Available() {
		return nil, errors.Wrap(ErrNoImplement, "error "+AlgoString(alg))
	}
	h := alg.New()
	io.Copy(h, r)
	return h.Sum(nil), nil
}

//ValueFromBytes returns hash value string from []byte
func ValueFromBytes(b []byte, alg crypto.Hash) ([]byte, error) {
	if !alg.Available() {
		return nil, errors.Wrap(ErrNoImplement, "error "+AlgoString(alg))
	}
	return alg.New().Sum(b), nil
}
