package mimir

// configuration holds information about a country and his banking specification
type configuration struct {
	CountryName             string
	CountryCode             countryCode
	IncludedCountryCode     []countryCode
	IsSEPAAvailable         bool
	SEPAIncludedCountryCode []countryCode
	AccountNumberExample    string
	IBANDefinition          definition
	BBANDefinition          definition
}

// definition holds information about format
type definition struct {
	Length      int
	Example     string
	PrintFormat string
	Structure   string
}

type countryCode string

func (c countryCode) String() string {
	return string(c)
}
