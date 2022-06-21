package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/common"
)

// NewPairResponse returns a new PairResponse
func NewPairResponse(token1 string, token0 string, oracles []sdk.AccAddress, active bool) PairResponse {
	var strOracles []string
	for _, oracle := range oracles {
		strOracles = append(strOracles, oracle.String())
	}

	pairID := common.PairNameFromDenoms([]string{token0, token1})
	return PairResponse{
		PairID:  pairID,
		Token1:  token1,
		Token0:  token0,
		Oracles: strOracles,
		Active:  active,
	}
}

// PairResponses is a slice of PairResponse
type PairResponses []PairResponse

/*
NewCurrentPrice returns an instance of CurrentPrice

Args:
  token0 (string):
  token1 (string):
  price (sdk.Dec): Price in units of token1 / token0
Returns:
  (CurrentPrice): Price for the asset pair.
*/
func NewCurrentPrice(token0 string, token1 string, price sdk.Dec) CurrentPrice {
	assetPair := common.AssetPair{Token0: token0, Token1: token1}
	return CurrentPrice{PairID: assetPair.String(), Price: price}
}

/*
NewCurrentTWAP returns an instance of CurrentTWAP

Args:
  token0 (string):
  token1 (string):
  price (sdk.Dec): Price in units of token1 / token0
Returns:
  (CurrentTWAP): Current TWAP price for the asset pair.
*/
func NewCurrentTWAP(token0 string, token1 string, numerator sdk.Dec, denominator sdk.Int, price sdk.Dec) CurrentTWAP {
	assetPair := common.AssetPair{Token0: token0, Token1: token1}
	return CurrentTWAP{PairID: assetPair.String(), Numerator: numerator, Denominator: denominator, Price: price}
}

// CurrentPrices is a slice of CurrentPrice
type CurrentPrices []CurrentPrice

// NewCurrentPriceResponse returns an instance of CurrentPriceResponse
func NewCurrentPriceResponse(pairID string, price sdk.Dec) CurrentPriceResponse {
	return CurrentPriceResponse{PairID: pairID, Price: price}
}

// CurrentPriceResponses is a slice of CurrentPriceResponse
type CurrentPriceResponses []CurrentPriceResponse

// NewPostedPrice returns a new PostedPrice
func NewPostedPrice(pair common.AssetPair, oracle sdk.AccAddress, price sdk.Dec, expiry time.Time,
) PostedPrice {
	return PostedPrice{
		PairID:        pair.String(),
		OracleAddress: oracle,
		Price:         price,
		Expiry:        expiry,
	}
}

// Validate performs a basic check of a PostedPrice params.
func (pp PostedPrice) Validate() error {
	if strings.TrimSpace(pp.PairID) == "" {
		return errors.New("market id cannot be blank")
	}
	if _, err := common.NewAssetPair(pp.PairID); err != nil {
		return err
	}
	if len(pp.OracleAddress) == 0 {
		return errors.New("oracle address cannot be empty")
	}
	if pp.Price.IsNegative() {
		return fmt.Errorf("posted price cannot be negative %s", pp.Price)
	}
	if pp.Expiry.Unix() <= 0 {
		return errors.New("expiry time cannot be zero")
	}
	return nil
}

// PostedPrices is a slice of PostedPrice
type PostedPrices []PostedPrice

// Validate checks if all the posted prices are valid and there are no
// duplicated entries.
func (pps PostedPrices) Validate() error {
	seenPrices := make(map[string]bool)
	for _, pp := range pps {
		if !pp.OracleAddress.Empty() && seenPrices[pp.PairID+pp.OracleAddress.String()] {
			return fmt.Errorf("duplicated posted price for marked id %s and oracle address %s", pp.PairID, pp.OracleAddress)
		}

		if err := pp.Validate(); err != nil {
			return err
		}
		seenPrices[pp.PairID+pp.OracleAddress.String()] = true
	}

	return nil
}

// NewPostedPrice returns a new PostedPrice
func NewPostedPriceResponse(
	pairID string, oracle sdk.AccAddress, price sdk.Dec, expiry time.Time,
) PostedPriceResponse {
	return PostedPriceResponse{
		PairID:        pairID,
		OracleAddress: oracle.String(),
		Price:         price,
		Expiry:        expiry,
	}
}

// PostedPriceResponses is a slice of PostedPriceResponse
type PostedPriceResponses []PostedPriceResponse

// SortDecs provides the interface needed to sort sdk.Dec slices
type SortDecs []sdk.Dec

func (a SortDecs) Len() int           { return len(a) }
func (a SortDecs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortDecs) Less(i, j int) bool { return a[i].LT(a[j]) }
