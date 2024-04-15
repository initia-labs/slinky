package sla

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/slinky/abci/strategies/currencypair"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	slatypes "github.com/skip-mev/slinky/x/sla/types"
)

// getStatuses returns the price feed status updates for each currency pair.
func getStatuses(ctx sdk.Context, currencyPairIDStrategy currencypair.CurrencyPairStrategy, currencyPairs []slinkytypes.CurrencyPair, prices map[string][]byte) map[slinkytypes.CurrencyPair]slatypes.UpdateStatus {
	validatorUpdates := make(map[slinkytypes.CurrencyPair]slatypes.UpdateStatus)

	for _, cp := range currencyPairs {
		currencyPairID := cp.String()

		if _, ok := prices[currencyPairID]; !ok {
			validatorUpdates[cp] = slatypes.VoteWithoutPrice
		} else {
			validatorUpdates[cp] = slatypes.VoteWithPrice
		}
	}

	return validatorUpdates
}
