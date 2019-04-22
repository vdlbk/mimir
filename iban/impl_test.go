package iban

import (
	"github.com/vdlbk/mimir/utils/consts"
	"testing"
)

func Test_IsIBANValid(t *testing.T) {
	for _, configuration := range consts.CountriesConfiguration {

		if ok, _ := IsIBANValid(configuration.IBANDefinition.Example); !ok {
			t.Errorf("Output was incorrect, IsIBANValid(%s)= false, want: true.", configuration.IBANDefinition.Example)
		}
	}
}
