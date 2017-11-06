package facade

import (
	"bytes"
	_ "crypto/sha1"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestCompare(t *testing.T) {
	result := "matched\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(bytes.NewBuffer([]byte(""))), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args2 := []string{"-a", "sha1", "-c", "da39a3ee5e6b4b0d3255bfef95601890afd80709"} //see https://en.wikipedia.org/wiki/SHA-1

	clearFlags()
	exit := Execute(ui, args2)
	if exit != ExitNormal {
		t.Errorf("Execute(compare) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outBuf.String()
	if str != "" {
		t.Errorf("Execute(compare) = \"%v\", want \"%v\".", str, "")
	}
	str = outErrBuf.String()
	if str != result {
		t.Errorf("Execute(compare) = \"%v\", want \"%v\".", str, result)
	}
}

func TestCompareErr(t *testing.T) {
	result := "unmatched\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(bytes.NewBuffer([]byte(""))), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args2 := []string{"-a", "sha1", "-c", "da39a3ee5e6b4b0d3255bfef95601890afd80709x"}

	clearFlags()
	exit := Execute(ui, args2)
	if exit != ExitAbnormal {
		t.Errorf("Execute(compare) = \"%v\", want \"%v\".", exit, ExitAbnormal)
	}
	str := outBuf.String()
	if str != "" {
		t.Errorf("Execute(compare) = \"%v\", want \"%v\".", str, "")
	}
	str = outErrBuf.String()
	if str != result {
		t.Errorf("Execute(compare) = \"%v\", want \"%v\".", str, result)
	}
}
