package parts

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func GetBank() error {

	myAddress, err := sdk.AccAddressFromBech32("cosmos13x77yexvf6qexfjg9czp6jhpv7vpjdwwnsrvej")

	if err != nil {
		return err
	}

	// Create a connection to the gRPC server.
	grpcConn, err := GetConnection()
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	// This creates a gRPC client to query the x/bank service.
	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: "uatom"},
	)
	if err != nil {
		return err
	}

	fmt.Println(bankRes.GetBalance()) // Prints the account balance

	return nil

}
