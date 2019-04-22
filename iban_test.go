package mimir

import (
	"testing"
)

func Test_IsIBANValid(t *testing.T) {
	for _, configuration := range countriesConfiguration {

		if ok, _ := IsIBANValid(configuration.IBANDefinition.Example); !ok {
			t.Errorf("Output was incorrect, IsIBANValid(%s)= false, want: true.", configuration.IBANDefinition.Example)
		}
	}
}
