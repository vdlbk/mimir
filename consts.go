package mimir

const (
	// ErrIBANTooshort is the error when a IBAN is too short for the validation process
	ErrIBANTooshort = mimirError("IBAN is too short")

	// ErrIBANIncorrectLength is the error when a IBAN does not have the required length
	ErrIBANIncorrectLength = mimirError("IBAN incorrect length")

	// ErrIBANInvalidChecksum is the error when a IBAN is invalid
	ErrIBANInvalidChecksum = mimirError("IBAN invalid checksum")

	// ErrCountryCodeDoesNotExist is the error when you lookup for a country code that does not exists
	ErrCountryCodeDoesNotExist = mimirError("Country Code does not exist")
)

type mimirError string

func (m mimirError) Error() string {
	return string(m)
}
