package facade

import (
	"bytes"
	_ "crypto/sha1"
	_ "crypto/sha256"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestList(t *testing.T) {
	result := "sha1 sha224 sha256\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"-l"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(list) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(list) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	if str != result {
		t.Errorf("Execute(list) = \"%v\", want \"%v\".", str, result)
	}
}
