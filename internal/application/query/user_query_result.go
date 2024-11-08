package query

import "github.com/Jollynjose/sistema-viatico-backend/internal/application/common"

// THIS FILE IS FOR QUERY RESULTS
type UserQueryResult struct {
	Result *common.UserResult `json:"result"`
}

type UsersQueryResult struct {
	Results []*common.UserResult `json:"results"`
}
