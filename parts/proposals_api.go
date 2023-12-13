package parts

import (
	"context"
	"fmt"
	"math"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	v04516_types "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func GetProposalsApi(ctx context.Context, voterAddr string) (*v04516_types.QueryProposalsResponse, error) {

	var voter = types.AccAddress{}
	if voterAddr != "" {
		var err error
		voter, err = types.AccAddressFromBech32(voterAddr)

		if err != nil {
			return nil, fmt.Errorf("can't stringify address %v, %v", voterAddr, err.Error())
		}
	}

	grpcConn, err := GetConnection()
	if err != nil {
		return nil, fmt.Errorf("request failed for address %v, %v", voter, err.Error())
	}
	defer grpcConn.Close()

	queryClient := v04516_types.NewQueryClient(grpcConn)

	resp, err := queryClient.Proposals(
		ctx,
		&v04516_types.QueryProposalsRequest{
			ProposalStatus: v04516_types.StatusVotingPeriod,
			Voter:          voter.String(),
			Pagination: &query.PageRequest{
				Limit: math.MaxUint64 - 1,
			},
		},
	)

	if err != nil {
		return nil, fmt.Errorf("request failed for address %v, %v", voter, err.Error())
	}

	return resp, nil
}
