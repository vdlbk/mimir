package models

type Configuration struct {
	CountryName             string
	CountryCode             CountryCode
	IncludedCountryCode     []CountryCode
	IsSEPAAvailable         bool
	SEPAIncludedCountryCode []CountryCode
	AccountNumberExample    string
	IBANDefinition          Definition
	BBANDefinition          Definition
}

type Definition struct {
	Length      int
	Example     string
	PrintFormat string
}

type CountryCode string

func (c CountryCode) String() string {
	return string(c)
}
