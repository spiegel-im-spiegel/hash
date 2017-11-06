# hash -- Calculating Hash Value

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/hash.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/hash)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/hash/master/LICENSE)

## Install

```
$ go get github.com/spiegel-im-spiegel/hash
```

Installing by [dep].

```
$ dep ensure -add github.com/spiegel-im-spiegel/hash
```

## Usage

```go
v, err := hash.Value(bytes.NewBuffer([]byte("")), crypto.SHA1)
if err != nil {
    return
}
fmt.Printf("%x\n", v)
// Output:
// da39a3ee5e6b4b0d3255bfef95601890afd80709
```

```go
v, err := hash.ValueFromBytes([]byte(""), crypto.SHA1)
if err != nil {
    return
}
fmt.Printf("%x\n", v)
// Output:
// da39a3ee5e6b4b0d3255bfef95601890afd80709
```

### Required hash function packages

| hash algorithm | import package |
|:---------------|:---------------|
| `crypto.MD4`         | `golang.org/x/crypto/md4` |
| `crypto.MD5`         | `crypto/md5` |
| `crypto.SHA1`        | `crypto/sha1` |
| `crypto.SHA224`      | `crypto/sha256` |
| `crypto.SHA256`      | `crypto/sha256` |
| `crypto.SHA384`      | `crypto/sha512` |
| `crypto.SHA512`      | `crypto/sha512` |
| `crypto.SHA512_224`  | `crypto/sha512` |
| `crypto.SHA512_256`  | `crypto/sha512` |
| `crypto.RIPEMD160`   | `golang.org/x/crypto/ripemd160` |
| `crypto.SHA3_224`    | `golang.org/x/crypto/sha3` |
| `crypto.SHA3_256`    | `golang.org/x/crypto/sha3` |
| `crypto.SHA3_384`    | `golang.org/x/crypto/sha3` |
| `crypto.SHA3_512`    | `golang.org/x/crypto/sha3` |
| `crypto.BLAKE2s_256` | `golang.org/x/crypto/blake2s` |
| `crypto.BLAKE2b_256` | `golang.org/x/crypto/blake2b` |
| `crypto.BLAKE2b_384` | `golang.org/x/crypto/blake2b` |
| `crypto.BLAKE2b_512` | `golang.org/x/crypto/blake2b` |

```go
package main

import (
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)
```


## Command Line Interface

See [cli/hash/README.md](cli/hash/README.md)

[hash]: https://github.com/spiegel-im-spiegel/hash "spiegel-im-spiegel/hash: Calculating Hash Value"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
