package data

import "testing"

func TestCheckValidation(t *testing.T){
	p := &Product{
		Name: "Felex",
		Price: 1.00,
		SKU: "abc-ase-rew",
	}

	err := p.Validator()

	if err != nil {
		t.Fatal(err)
	}
}