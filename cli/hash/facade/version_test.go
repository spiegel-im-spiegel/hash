package facade

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestVersionMin(t *testing.T) {
	result := "hash\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"-v"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outBuf.String()
	if str != "" {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, "")
	}
	str = outErrBuf.String()
	if str != result {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, result)
	}
}

func TestVersionNum(t *testing.T) {
	Version = "TestVersion"
	result := "hash vTestVersion\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"-v"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outBuf.String()
	if str != "" {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, "")
	}
	str = outErrBuf.String()
	if str != result {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, result)
	}
}

func TestVersionNumCommit(t *testing.T) {
	Version = "TestVersion"
	Commit = "<commit value>"
	result := "hash vTestVersion (commit: <commit value>)\n"

	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"-v"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outBuf.String()
	if str != "" {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, "")
	}
	str = outErrBuf.String()
	if str != result {
		t.Errorf("Execute(version) = \"%v\", want \"%v\".", str, result)
	}
}
