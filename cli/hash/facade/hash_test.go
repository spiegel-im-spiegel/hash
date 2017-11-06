package facade

import (
	"bytes"
	_ "crypto/sha256"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestHash(t *testing.T) {
	result := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855\n" //see https://en.wikipedia.org/wiki/SHA-2

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(bytes.NewBuffer([]byte(""))), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args2 := []string{}

	clearFlags()
	exit := Execute(ui, args2)
	if exit != ExitNormal {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	if str != result {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", str, result)
	}
}

func TestHashFile(t *testing.T) {
	result := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855\n" //see https://en.wikipedia.org/wiki/SHA-2

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(bytes.NewBuffer([]byte(""))), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args2 := []string{"testdata/empty"}

	clearFlags()
	exit := Execute(ui, args2)
	if exit != ExitNormal {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	if str != result {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", str, result)
	}
}

func TestHashFileErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(bytes.NewBuffer([]byte(""))), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args2 := []string{"testdata/noexist"}

	clearFlags()
	exit := Execute(ui, args2)
	if exit != ExitAbnormal {
		t.Errorf("Execute(hash) = \"%v\", want \"%v\".", exit, ExitAbnormal)
	}
}

func clearFlags() {
	rootCmd.Flag("algo").Value.Set(defaultAlg)
	rootCmd.Flag("compare").Value.Set("")
	rootCmd.Flag("list").Value.Set("false")
	rootCmd.Flag("version").Value.Set("false")
}
