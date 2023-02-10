package snapshot

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

var snapshotGraphqlclient = graphql.NewClient("https://hub.snapshot.org/graphql", nil)

func (app *SnapshotIndexerApp) GetSnapshotProposalVoters() ([]*SnapshotProposalVoter, error) {
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
	result := make([]*SnapshotProposalVoter, 0)
	for _, proposal := range proposalsQuery.Proposals {
		if proposal.Start >= app.params.SnapshotStart && proposal.End <= app.params.SnapshotEnd {
			// TODO: get proposal voters and append unexisting ones
			log.Info(proposal.Title)
		}
	}
	return result, nil
}
