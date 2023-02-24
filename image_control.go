package app_test
import (
	"fmt"
  	"github.com/cosmos/cosmos-sdk/testutil"
  }
  
  func image_contr(t *testing.T) {
	// Set MaxTxsInBlock  = 2
	numberimages := 2
    	require.NoError(t, ctx.Codec.UnmarshalJSON(res[numberTxsinBlocks].Bytes(), &resp))
	require.Equal(t, errCode, resp.Code)
}

//  bubble-like sorting position
  func sorting(net *network.Network) []string {
	return []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastAsync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdk.NewInt(10))).String()),
	}
}
