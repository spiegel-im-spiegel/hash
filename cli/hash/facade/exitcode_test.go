package facade

import "testing"

func TestExitNormal(t *testing.T) {
	exit := ExitNormal
	if exit.Int() != 0 {
		t.Errorf("ExitNormal.Int() = \"%v\", want 0.", exit.Int())
	}
	if exit.String() != "normal end" {
		t.Errorf("ExitNormal.String() = \"%v\", want \"normal end\".", exit.String())
	}
}

func TestExitAbnormal(t *testing.T) {
	exit := ExitAbnormal
	if exit.Int() != 1 {
		t.Errorf("ExitNormal.Int() = \"%v\", want 1.", exit.Int())
	}
	if exit.String() != "abnormal end" {
		t.Errorf("ExitNormal.String() = \"%v\", want \"abnormal end\".", exit.String())
	}
}

func TestExitUnknown(t *testing.T) {
	exit := ExitCode(2)
	if exit.String() != "unknown" {
		t.Errorf("ExitNormal.String() = \"%v\", want \"unknown\".", exit.String())
	}
}
