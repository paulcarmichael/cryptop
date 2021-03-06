package fincrypt

import "testing"

func Test_Luhn_ZeroLengthInput(t *testing.T) {
	operation := LuhnOperation{}
	operation.Mode = ModeValidate

	_, err := operation.Calculate()

	if err == nil {
		t.Errorf("Expected method to throw an error with zero length input")
	}
}

func Test_Luhn_NonNumericInput(t *testing.T) {
	operation := LuhnOperation{}
	operation.Input = "4929142A85824009"
	operation.Mode = ModeValidate

	_, err := operation.Calculate()

	if err == nil {
		t.Errorf("Expected method to throw an error with non numeric input")
	}
}

func Test_Luhn_Valid(t *testing.T) {
	operation := LuhnOperation{}
	operation.Input = "4929142285824009"
	operation.Mode = ModeValidate

	result, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	if result != "9" {
		t.Errorf("Expected a luhn digit of 9, method returned %s", result)
	}
}

func Test_Luhn_Invalid(t *testing.T) {
	operation := LuhnOperation{}
	operation.Input = "4929142285824003"
	operation.Mode = ModeValidate

	result, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	if result != "9" {
		t.Errorf("Expected a luhn digit of 9, method returned %s", result)
	}
}

func Test_Luhn_Generate(t *testing.T) {
	operation := LuhnOperation{}
	operation.Input = "7992739871"
	operation.Mode = ModeGenerate

	result, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	if result != "3" {
		t.Errorf("Expected a luhn digit of 3, method returned %s", result)
	}
}

func Test_Luhn_GenerateZero(t *testing.T) {
	operation := LuhnOperation{}
	operation.Input = "7992739871399"
	operation.Mode = ModeGenerate

	result, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	if result != "0" {
		t.Errorf("Expected a luhn digit of 0, method returned %s", result)
	}
}
