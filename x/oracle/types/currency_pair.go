package types

import (
	"fmt"
	"strings"
)

const (
	ethereum = "ETHEREUM"
)

func NewCurrencyPair(base, quote string) CurrencyPair {
	return CurrencyPair{
		Base:  base,
		Quote: quote,
	}
}

// ValidateBasic checks that the Base / Quote strings in the CurrencyPair are formatted correctly, i.e
// Base + Quote are non-empty, and are in upper-case.
func (cp CurrencyPair) ValidateBasic() error {
	// strings must be valid
	if cp.Base == "" || cp.Quote == "" {
		return fmt.Errorf("empty quote or base string")
	}
	// check formatting of base / quote
	if strings.ToUpper(cp.Base) != cp.Base {
		return fmt.Errorf("incorrectly formatted base string, expected: %s got: %s", strings.ToUpper(cp.Base), cp.Base)
	}
	if strings.ToUpper(cp.Quote) != cp.Quote {
		return fmt.Errorf("incorrectly formatted quote string, expected: %s got: %s", strings.ToUpper(cp.Quote), cp.Quote)
	}
	return nil
}

func (cp CurrencyPair) ToString() string {
	return fmt.Sprintf("%s/%s", cp.Base, cp.Quote)
}

func CurrencyPairFromString(s string) (CurrencyPair, error) {
	split := strings.Split(s, "/")
	if len(split) != 2 {
		return CurrencyPair{}, fmt.Errorf("incorrectly formatted CurrencyPair: %s", s)
	}
	return CurrencyPair{
		Base:  split[0],
		Quote: split[1],
	}, nil
}

// Decimals returns the number of decimals that the quote will be reported to. If the quote is Ethereum, then
// the number of decimals is 18. Otherwise, the decimals will be reorted to 8.
func (cp CurrencyPair) Decimals() int {
	if strings.ToUpper(cp.Quote) == ethereum {
		return 18
	}
	return 8
}