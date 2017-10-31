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
var algOrder = []string{
	algMap[crypto.MD4],
	algMap[crypto.MD5],
	algMap[crypto.MD5SHA1],
	algMap[crypto.SHA1],
	algMap[crypto.SHA224],
	algMap[crypto.SHA256],
	algMap[crypto.SHA384],
	algMap[crypto.SHA512],
	algMap[crypto.SHA512_224],
	algMap[crypto.SHA512_256],
	algMap[crypto.RIPEMD160],
	algMap[crypto.SHA3_224],
	algMap[crypto.SHA3_256],
	algMap[crypto.SHA3_384],
	algMap[crypto.SHA3_512],
	algMap[crypto.BLAKE2s_256],
	algMap[crypto.BLAKE2b_256],
	algMap[crypto.BLAKE2b_384],
	algMap[crypto.BLAKE2b_512],
}

//FuncList returns string of hash functions list
func FuncList() string {
	buf := new(bytes.Buffer)
	sep := ""
	for _, name := range algOrder {
		alg, err := Algorithm(name)
		if err == nil {
			if alg.Available() {
				io.WriteString(buf, sep+name)
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

func algoString(alg crypto.Hash) string {
	if str, ok := algMap[alg]; ok {
		return str
	}
	return "unknown hash algorithm"
}

//Value returns hash value string from io.Reader
func Value(r io.Reader, alg crypto.Hash) ([]byte, error) {
	if !alg.Available() {
		return nil, errors.Wrap(ErrNoImplement, "error "+algoString(alg))
	}
	h := alg.New()
	io.Copy(h, r)
	return h.Sum(nil), nil
}

//ValueFromBytes returns hash value string from []byte
func ValueFromBytes(b []byte, alg crypto.Hash) ([]byte, error) {
	if !alg.Available() {
		return nil, errors.Wrap(ErrNoImplement, "error "+algoString(alg))
	}
	return alg.New().Sum(b), nil
}
