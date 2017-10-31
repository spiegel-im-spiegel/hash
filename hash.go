package hash

import (
	"crypto"
	"fmt"
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
	crypto.MD5SHA1:     "md5sha1",     //no implementation
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

func algoAtring(alg crypto.Hash) string {
	if str, ok := algMap[alg]; ok {
		return str
	}
	return "unknown hash algorithm"
}

//Value returns hash value string from io.Reader
func Value(r io.Reader, alg crypto.Hash) (string, error) {
	if !alg.Available() {
		return "", errors.Wrap(ErrNoImplement, "error "+algoAtring(alg))
	}
	h := alg.New()
	io.Copy(h, r)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

//ValueFromBytes returns hash value string from []byte
func ValueFromBytes(b []byte, alg crypto.Hash) (string, error) {
	if !alg.Available() {
		return "", errors.Wrap(ErrNoImplement, "error "+algoAtring(alg))
	}
	return fmt.Sprintf("%x", alg.New().Sum(b)), nil
}
