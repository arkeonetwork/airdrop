package snapshot

import (
	"context"
	"fmt"

	"github.com/ArkeoNetwork/airdrop/pkg/types"
	"github.com/hasura/go-graphql-client"
)

var snapshotGraphqlclient = graphql.NewClient("https://hub.snapshot.org/graphql", nil)

func getSingleProposalVoters(proposalId string) ([]string, error) {
	var proposalVotesQuery struct {
		Votes []struct {
			Id string `json:"id"`
			Voter string `json:"voter"`
		} `graphql:"votes (first: 1000, where: { proposal: $proposalId }, orderBy: \"created\", orderDirection: desc)"`
		// NOTE: the above line should be a one-liner, and there should be no whitespaces 
		// between `graphql:` and `"`.
	}
	variables := map[string]interface{}{
		"proposalId":  proposalId,
	}
	err := snapshotGraphqlclient.Query(context.Background(), &proposalVotesQuery, variables)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, vote := range proposalVotesQuery.Votes {
		result = append(result, vote.Voter)
	}
	return result, nil
}

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func (app *SnapshotIndexerApp) GetSnapshotProposalVoters() ([]*types.SnapshotVoter, error) {
	var proposalsQuery struct {
		Proposals []struct {
			Id string `json:"id"`
			Start uint64 `json:"start"`
			End uint64 `json:"end"`
			Title string `json:"title"`
		} `graphql:"proposals (first: 1000, where: { space: \"shapeshiftdao.eth\" }, orderBy: \"created\", orderDirection: desc)"`
		// NOTE: the above line should be a one-liner, and there should be no whitespaces 
		// between `graphql:` and `"`.
	}
	err := snapshotGraphqlclient.Query(context.Background(), &proposalsQuery, nil)
	if err != nil {
		return nil, err
	}
	allVoters := make([]string, 0)
	for _, proposal := range proposalsQuery.Proposals {
		if proposal.Start >= app.params.SnapshotStart && proposal.End <= app.params.SnapshotEnd {
			// getting voters for a single proposal
			log.Info("getting voters for proposal ", proposal.Title)
			proposalVoters, err := getSingleProposalVoters(proposal.Id)
			if err != nil {
				panic(fmt.Sprintf("error getting proposal voters: %+v", err))
			}
			// appending this proposal voters and removing duplicated addresses
			allVoters = removeDuplicateStr(append(allVoters, proposalVoters...))
		}
	}
	result := make([]*types.SnapshotVoter, 0)
	for _, voter := range allVoters {
		result = append(result, &types.SnapshotVoter{Address: voter})
	}
	return result, nil
}
