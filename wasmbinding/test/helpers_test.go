package wasmbinding

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/app"
	nibirusimapp "github.com/NibiruChain/nibiru/simapp"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func SetupCustomApp(t *testing.T, addr sdk.AccAddress) (*app.NibiruApp, sdk.Context) {
	nibiru, ctx := CreateTestInput()
	wasmKeeper := nibiru.WasmKeeper

	storePerpCode(t, ctx, nibiru, addr)

	cInfo := wasmKeeper.GetCodeInfo(ctx, 1)
	require.NotNil(t, cInfo)

	return nibiru, ctx
}

func storePerpCode(t *testing.T, ctx sdk.Context, nibiru *app.NibiruApp, addr sdk.AccAddress) {
	wasmCode, err := os.ReadFile("../testdata/perp.wasm")
	require.NoError(t, err)

	contractKeeper := keeper.NewDefaultPermissionKeeper(nibiru.WasmKeeper)
	_, _, err = contractKeeper.Create(ctx, addr, wasmCode, &wasmtypes.AccessConfig{
		Permission: wasmtypes.AccessTypeEverybody,
	})
	require.NoError(t, err)
}

func instantiatePerpContract(t *testing.T, ctx sdk.Context, nibiru *app.NibiruApp, funder sdk.AccAddress) sdk.AccAddress {
	initMsgBz := []byte("{}")
	contractKeeper := keeper.NewDefaultPermissionKeeper(nibiru.WasmKeeper)
	codeID := uint64(1)
	addr, _, err := contractKeeper.Instantiate(ctx, codeID, funder, funder, initMsgBz, "demo contract", nil)
	require.NoError(t, err)

	return addr
}

func CreateTestInput() (*app.NibiruApp, sdk.Context) {
	app := nibirusimapp.NewTestNibiruApp(true)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "kujira-1", Time: time.Now().UTC()})
	return app, ctx
}

func FundAccount(t *testing.T, ctx sdk.Context, app *app.NibiruApp, acct sdk.AccAddress) {
	err := simapp.FundAccount(app.BankKeeper, ctx, acct, sdk.NewCoins(
		sdk.NewCoin("uosmo", sdk.NewInt(10000000000)),
	))
	require.NoError(t, err)
}

// we need to make this deterministic (same every test run), as content might affect gas costs
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func RandomAccountAddress() sdk.AccAddress {
	_, _, addr := keyPubAddr()
	return addr
}

func RandomBech32AccountAddress() string {
	return RandomAccountAddress().String()
}
