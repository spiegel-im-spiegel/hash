package main

import (
	"os"

	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"

	"github.com/spiegel-im-spiegel/gocli"
	"github.com/spiegel-im-spiegel/hash/cli/hash/facade"
)

func main() {
	os.Exit(facade.Execute(
		gocli.NewUI(
			gocli.Reader(os.Stdin),
			gocli.Writer(os.Stdout),
			gocli.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Int())
}
