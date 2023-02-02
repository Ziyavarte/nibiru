package asset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NibiruChain/nibiru/x/common/denoms"
)

func TestTryNewPair(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		tokenStr string
		err      error
	}{
		{
			"only one token",
			denoms.NIBI,
			ErrInvalidTokenPair,
		},
		{
			"more than 2 tokens",
			fmt.Sprintf("%s:%s:%s", denoms.NIBI, denoms.NUSD, denoms.USDC),
			ErrInvalidTokenPair,
		},
		{
			"different separator",
			fmt.Sprintf("%s,%s", denoms.NIBI, denoms.NUSD),
			ErrInvalidTokenPair,
		},
		{
			"correct pair",
			fmt.Sprintf("%s:%s", denoms.NIBI, denoms.NUSD),
			nil,
		},
		{
			"empty token identifier",
			fmt.Sprintf(":%s", denoms.ETH),
			fmt.Errorf("empty token identifiers are not allowed"),
		},
		{
			"invalid denom 1",
			fmt.Sprint("-invalid1:valid"),
			fmt.Errorf("invalid denom"),
		},
		{
			"invalid denom 2",
			fmt.Sprint("valid:-invalid2"),
			fmt.Errorf("invalid denom"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := TryNew(tc.tokenStr)
			if tc.err != nil {
				require.ErrorContains(t, err, tc.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetDenoms(t *testing.T) {
	pair := MustNew("uatom:unibi")

	require.Equal(t, "uatom", pair.BaseDenom())
	require.Equal(t, "unibi", pair.QuoteDenom())
}

func TestEquals(t *testing.T) {
	pair := MustNew("abc:xyz")
	matchingOther := MustNew("abc:xyz")
	mismatchToken1 := MustNew("abc:abc")
	inversePair := MustNew("xyz:abc")

	require.True(t, pair.Equal(matchingOther))
	require.False(t, pair.Equal(inversePair))
	require.False(t, pair.Equal(mismatchToken1))
}

func TestMustNewAssetPair(t *testing.T) {
	require.Panics(t, func() {
		MustNew("aaa:bbb:ccc")
	})

	require.NotPanics(t, func() {
		MustNew("aaa:bbb")
	})
}

func TestInverse(t *testing.T) {
	pair := MustNew("abc:xyz")
	inverse := pair.Inverse()
	require.Equal(t, "xyz", inverse.BaseDenom())
	require.Equal(t, "abc", inverse.QuoteDenom())
}