package mimir

// Errors
const (
	// ErrIBANTooshort is the error when an IBAN is too short for the validation process
	ErrIBANTooshort = mimirError("IBAN is too short")

	// ErrIBANIncorrectLength is the error when a IBAN does not have the required length
	ErrIBANIncorrectLength = mimirError("IBAN incorrect length")

	// ErrIBANInvalidChecksum is the error when a IBAN is invalid
	ErrIBANInvalidChecksum = mimirError("IBAN invalid checksum")

	// ErrCountryCodeDoesNotExist is the error when you lookup for a country code that does not exists
	ErrCountryCodeDoesNotExist = mimirError("Country Code does not exist")

	// ErrABARTNInvalidLength is the error when an ABA Routing Number is too short the the validation process
	ErrABARTNInvalidLength = mimirError("ABA is too short. 9 digits is expected")

	// ErrABAInvalidChecksum is the error when a ABA Routing Number is invalid
	ErrABAInvalidChecksum = mimirError("ABA Routing Number invalid checksum")
)

type mimirError string

func (m mimirError) Error() string {
	return string(m)
}

// IBAN - Structure digit keys
const (
	AccountNumberIBANDigitKey                = "a" // alphanumeric
	NationalBankCodeIBANDigitKey             = "b" // numeric
	CountryCodeIBANDigitKey                  = "c" // alphabetic
	CheckIBANDigitKey                        = "k" // checksum
	NationalIdentificationNumberIBANDigitKey = "i"
	CurrencyIBANDigitKey                     = "m" // alphanumeric
	AccountHolderIBANDigitKey                = "n"
	ReserveNumberIBANDigitKey                = "o" // always 0
	BranchCodeIBANDigitKey                   = "s" // counter code
	AccountTypeIBANDigitKey                  = "t"
	SWIFTBICCodeIBANDigitKey                 = "w" // alphanumeric
	NationalCheckIBANDigitKey                = "x" // numeric
)

// ABA Routing Number - Structure digit keys
const (
	FederalReserveRoutingSymbolABADigitKey = "f" // numeric
	ABAInstitutionIdentifierDigitKey       = "a" // numeric
	CheckABADigitKey                       = "k" // checksum // numeric
)
