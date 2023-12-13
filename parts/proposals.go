package parts

import (
	"context"
	"fmt"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func GetProposals(ctx context.Context) (*v1beta1.QueryProposalsResponse, error) {
	voterAddr := "cosmos13x77yexvf6qexfjg9czp6jhpv7vpjdwwnsrvej"

	resp, err := GetProposalsApi(ctx, voterAddr)

	GetDecodedProposals(resp)

	return resp, err
}

type ProposalExcerpt struct {
	ProposalId uint64                  `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty"`
	Content    *codectypes.Any         `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status     govtypes.ProposalStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.gov.v1beta1.ProposalStatus" json:"status,omitempty"`
	// final_tally_result is the final tally result of the proposal. When
	// querying a proposal via gRPC, this field is not populated until the
	// proposal's voting period has ended.
	FinalTallyResult govtypes.TallyResult                     `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result"`
	SubmitTime       time.Time                                `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3,stdtime" json:"submit_time"`
	DepositEndTime   time.Time                                `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3,stdtime" json:"deposit_end_time"`
	TotalDeposit     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=total_deposit,json=totalDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_deposit"`
	VotingStartTime  time.Time                                `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3,stdtime" json:"voting_start_time"`
	VotingEndTime    time.Time                                `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3,stdtime" json:"voting_end_time"`
}

func (p *ProposalExcerpt) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	print("ProposalExcerpt UnpackInterfaces")
	if p.Content != nil {
		var content govtypes.Content
		return unpacker.UnpackAny(p.Content, &content)
	}

	return nil
}

func GetDecodedProposals(proposalResponse *govtypes.QueryProposalsResponse) []ProposalExcerpt {

	proposals := make([]ProposalExcerpt, 0)

	for _, proposal := range proposalResponse.Proposals {
		fmt.Printf("orig proposal content: %#v", proposal.Content.GetCachedValue())

		proposalExcerpt := ProposalExcerpt{
			ProposalId:       proposal.ProposalId,
			Content:          proposal.Content,
			Status:           proposal.Status,
			FinalTallyResult: proposal.FinalTallyResult,
			SubmitTime:       proposal.SubmitTime,
			DepositEndTime:   proposal.DepositEndTime,
			TotalDeposit:     proposal.TotalDeposit,
			VotingStartTime:  proposal.VotingStartTime,
			VotingEndTime:    proposal.VotingEndTime,
		}

		if proposal.Content.TypeUrl == "/cosmos.gov.v1beta1.TextProposal" {

			fmt.Printf("Content plain: %#v", proposalExcerpt.Content)
			fmt.Printf("Content cached: %#v", proposalExcerpt.Content.GetCachedValue())
			fmt.Printf("proposal.GetContent(): %#v", proposal.GetContent())

			err := proposal.Content.Unmarshal(proposal.Content.Value)
			if err != nil {
				fmt.Printf("err: %#v", err)
			}
		}

		proposals = append(proposals, proposalExcerpt)
	}

	return proposals
}
