package fincrypt

import (
	"math/big"
	"testing"
)

func Test_RSA_Encrypt(t *testing.T) {
	operation := RSAOperation{}
	operation.Input = "736E64207265696E666F7263656D656E74732C20776527726520676F696E6720746F20616476616E6365"
	operation.PublicExponent = "010001"
	operation.PrivateExponent = "526BFCEE31E859D15EA8C48F0364E1B0BD8A0C2FB8AE7BB22792E705AAC863562D6850C6B1BAF20B24148F62BCA521413E623299CD7D8C3DC9A7C80A0A97AE2AB20379F749A62BA9E2FEB89C1BE9C5FAE2C728A4CDE4BD5C1CF52A6485CEC52970BB11913801C3B5E351BAC870A96A09F3DA205374033F8B43A1E9AE337495E1"
	operation.Modulus = "AEFB4488970CE282229D8EB7CD58B1507BE28EE23C63E1D1C0C8BDCFAF2DAD8EA8DCE7FE07CF120344B645AFECB1730A986275F5165E630B4F82187660937C77F1F91007EA81C56DD08F7CB907BBDEBB57A3DC54F5D30C7C0E4AC39DF39B8D3D5A87A8A126CDE771DD9046AA48FD6EC4D7C4813BC9D74BD4DD853445F56BCB87"
	operation.HashMode = HashModeSHA256
	operation.Direction = DirectionEncrypt

	_, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	// no expected check here as the result is different each time
}

func Test_RSA_Decrypt(t *testing.T) {
	operation := RSAOperation{}
	operation.Input = "5E9021E2C2342D8B75D3D31D788764A7CA6C1895AACA126DFA286921D95D70715BE4BDD29546CC388F8C871C1B5ABAFD000045174E6C02D9581993C05575BEC5C344A2F7020C0C3CB62F4ACAFAD0B0053689A4D5A88F17FE2DBC22B13F00C002D7DD142ED1801FBE84EF42B5210B1F25BDE4095372C7B092E8F830E296F26917"
	operation.PublicExponent = "010001"
	operation.PrivateExponent = "526BFCEE31E859D15EA8C48F0364E1B0BD8A0C2FB8AE7BB22792E705AAC863562D6850C6B1BAF20B24148F62BCA521413E623299CD7D8C3DC9A7C80A0A97AE2AB20379F749A62BA9E2FEB89C1BE9C5FAE2C728A4CDE4BD5C1CF52A6485CEC52970BB11913801C3B5E351BAC870A96A09F3DA205374033F8B43A1E9AE337495E1"
	operation.Modulus = "AEFB4488970CE282229D8EB7CD58B1507BE28EE23C63E1D1C0C8BDCFAF2DAD8EA8DCE7FE07CF120344B645AFECB1730A986275F5165E630B4F82187660937C77F1F91007EA81C56DD08F7CB907BBDEBB57A3DC54F5D30C7C0E4AC39DF39B8D3D5A87A8A126CDE771DD9046AA48FD6EC4D7C4813BC9D74BD4DD853445F56BCB87"
	operation.HashMode = HashModeSHA256
	operation.Direction = DirectionDecrypt

	result, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	expected := "73656E64207265696E666F7263656D656E74732C20776527726520676F696E6720746F20616476616E6365"

	if result != expected {
		t.Errorf("Expected [%s], Calculate returned [%s]", expected, result)
	}
}

func Test_RSA_Decrypt_WithLabel(t *testing.T) {
	operation := RSAOperation{}
	operation.Input = "4d1ee10e8f286390258c51a5e80802844c3e6358ad6690b7285218a7c7ed7fc3a4c7b950fbd04d4b0239cc060dcc7065ca6f84c1756deb71ca5685cadbb82be025e16449b905c568a19c088a1abfad54bf7ecc67a7df39943ec511091a34c0f2348d04e058fcff4d55644de3cd1d580791d4524b92f3e91695582e6e340a1c50b6c6d78e80b4e42c5b4d45e479b492de42bbd39cc642ebb80226bb5200020d501b24a37bcc2ec7f34e596b4fd6b063de4858dbf5a4e3dd18e262eda0ec2d19dbd8e890d672b63d368768360b20c0b6b8592a438fa275e5fa7f60bef0dd39673fd3989cc54d2cb80c08fcd19dacbc265ee1c6014616b0e04ea0328c2a04e73460"
	operation.PublicExponent = "03"

	i, ok := new(big.Int).SetString("14314132931241006650998084889274020608918049032671858325988396851334124245188214251956198731333464217832226406088020736932173064754214329009979944037640912127943488972644697423190955557435910767690712778463524983667852819010259499695177313115447116110358524558307947613422897787329221478860907963827160223559690523660574329011927531289655711860504630573766609239332569210831325633840174683944553667352219670930408593321661375473885147973879086994006440025257225431977751512374815915392249179976902953721486040787792801849818254465486633791826766873076617116727073077821584676715609985777563958286637185868165868520557", 10)

	if ok == false {
		t.Errorf("Invalid big number for RSA modulus")
	}

	operation.Modulus = i.Text(16)

	j, ok := new(big.Int).SetString("9542755287494004433998723259516013739278699355114572217325597900889416163458809501304132487555642811888150937392013824621448709836142886006653296025093941418628992648429798282127303704957273845127141852309016655778568546006839666463451542076964744073572349705538631742281931858219480985907271975884773482372966847639853897890615456605598071088189838676728836833012254065983259638538107719766738032720239892094196108713378822882383694456030043492571063441943847195939549773271694647657549658603365629458610273821292232646334717612674519997533901052790334279661754176490593041941863932308687197618671528035670452762731", 10)

	if ok == false {
		t.Errorf("Invalid big number for RSA private exponent")
	}

	operation.PrivateExponent = j.Text(16)
	operation.Label = "orders"
	operation.HashMode = HashModeSHA256
	operation.Direction = DirectionDecrypt

	result, err := operation.Calculate()

	if err != nil {
		t.Errorf(err.Error())
	}

	expected := "73656E64207265696E666F7263656D656E74732C20776527726520676F696E6720746F20616476616E6365"

	if result != expected {
		t.Errorf("Expected [%s], Calculate returned [%s]", expected, result)
	}
}
