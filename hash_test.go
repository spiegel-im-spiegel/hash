package hash

import (
	"bytes"
	"crypto"
	_ "crypto/sha1"
	"fmt"
	"io"
	"testing"
)

func TestFuncList(t *testing.T) {
	res := "sha1"
	buf := new(bytes.Buffer)
	io.Copy(buf, FuncList())
	str := string(buf.Bytes())
	if str != res {
		t.Errorf("TestFuncList() = \"%v\", want \"%v\".", str, res)
	}
}

func TestValueFromBytesFoo(t *testing.T) {
	_, err := Algorithm("foo")
	if err == nil {
		t.Error("Algorithm() is nil error, not want nil error.")
	}
}

func TestValueFromBytesZero(t *testing.T) {
	_, err := ValueFromBytes([]byte(""), crypto.Hash(0))
	if err == nil {
		t.Error("ValueFromBytes() is nil error, not want nil error.")
	}
}

func TestValueMS5SHA1(t *testing.T) {
	alg, err := Algorithm("md5sha1")
	if err != nil {
		t.Errorf("Algorithm() is \"%v\", want nil error.", err)
		return
	}
	_, err = Value(bytes.NewBuffer([]byte("")), alg)
	if err == nil {
		t.Error("Value() is nil error, not want nil error.")
	}
}
func TestValueFromBytesMS5SHA1(t *testing.T) {
	alg, err := Algorithm("md5sha1")
	if err != nil {
		t.Errorf("Algorithm() is \"%v\", want nil error.", err)
		return
	}
	_, err = ValueFromBytes([]byte(""), alg)
	if err == nil {
		t.Error("ValueFromBytes() is nil error, not want nil error.")
	}
}

func TestValueSHA1(t *testing.T) {
	hv := "da39a3ee5e6b4b0d3255bfef95601890afd80709" //see https://en.wikipedia.org/wiki/SHA-1
	alg, err := Algorithm("sha1")
	if err != nil {
		t.Errorf("Algorithm() is \"%v\", want nil error.", err)
		return
	}
	v, err := Value(bytes.NewBuffer([]byte("")), alg)
	if err != nil {
		t.Errorf("ValueFromBytes() is \"%v\", want nil error.", err)
		return
	}
	str := fmt.Sprintf("%x", v)

	if str != hv {
		t.Errorf("ValueFromBytes() = \"%v\", want \"%v\".", str, hv)
	}
}

func TestValueFromBytesSHA1(t *testing.T) {
	hv := "da39a3ee5e6b4b0d3255bfef95601890afd80709" //see https://en.wikipedia.org/wiki/SHA-1
	alg, err := Algorithm("sha1")
	if err != nil {
		t.Errorf("Algorithm() is \"%v\", want nil error.", err)
		return
	}
	v, err := ValueFromBytes([]byte(""), alg)
	if err != nil {
		t.Errorf("Value() is \"%v\", want nil error.", err)
		return
	}
	str := fmt.Sprintf("%x", v)
	if str != hv {
		t.Errorf("Value() = \"%v\", want \"%v\".", str, hv)
	}
}

func ExampleValue() {
	v, err := Value(bytes.NewBuffer([]byte("")), crypto.SHA1)
	if err != nil {
		return
	}
	fmt.Printf("%x\n", v)
	// Output:
	// da39a3ee5e6b4b0d3255bfef95601890afd80709
}

func ExampleValueFromBytes() {
	v, err := ValueFromBytes([]byte(""), crypto.SHA1)
	if err != nil {
		return
	}
	fmt.Printf("%x\n", v)
	// Output:
	// da39a3ee5e6b4b0d3255bfef95601890afd80709
}
