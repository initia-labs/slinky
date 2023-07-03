package coinbase

import (
	"context"

	"cosmossdk.io/log"
	"github.com/skip-mev/slinky/oracle/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
)

const (
	// Name is the name of the provider.
	Name = "coinbase"
)

var _ types.Provider = (*Provider)(nil)

// Provider implements the Provider interface for Coinbase. This provider
// is a very simple implementation that fetches spot prices from the Coinbase API.
type Provider struct {
	pairs  []oracletypes.CurrencyPair
	logger log.Logger
}

// NewProvider returns a new Coinbase provider.
//
// THIS PROVIDER SHOULD NOT BE USED IN PRODUCTION. IT IS ONLY MEANT FOR TESTING.
func NewProvider(logger log.Logger, pairs []oracletypes.CurrencyPair) *Provider {
	return &Provider{
		pairs:  pairs,
		logger: logger,
	}
}

// Name returns the name of the provider.
func (p *Provider) Name() string {
	return Name
}

// GetPrices returns the current set of prices for each of the currency pairs.
func (p *Provider) GetPrices(ctx context.Context) (map[oracletypes.CurrencyPair]types.QuotePrice, error) {
	resp := make(map[oracletypes.CurrencyPair]types.QuotePrice)

	for _, currencyPair := range p.pairs {
		spotPrice, err := getPriceForPair(ctx, currencyPair)
		if err != nil {
			p.logger.Error(
				p.Name(),
				"failed to get price for pair", currencyPair,
				"err", err,
			)
			continue
		}

		resp[currencyPair] = *spotPrice
	}

	return resp, nil
}

// SetPairs sets the currency pairs that the provider will fetch prices for.
func (p *Provider) SetPairs(pairs ...oracletypes.CurrencyPair) {
	p.pairs = pairs
}

// GetPairs returns the currency pairs that the provider is fetching prices for.
func (p *Provider) GetPairs() []oracletypes.CurrencyPair {
	return p.pairs
}