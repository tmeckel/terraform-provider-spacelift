package structs

import "github.com/shurcooL/graphql"

type StackDependencyReference struct {
	ID         string `graphql:"id"`
	OutputName string `graphql:"outputName"`
	InputName  string `graphql:"inputName"`
	Type       string `graphql:"type"`
}

type StackDependencyReferenceInput struct {
	OutputName graphql.String `json:"outputName"`
	InputName  graphql.String `json:"inputName"`
	Type       graphql.String `json:"type"`
}

type StackDependencyReferenceUpdateInput struct {
	ID         graphql.ID     `json:"id"`
	OutputName graphql.String `json:"outputName"`
	InputName  graphql.String `json:"inputName"`
	Type       graphql.String `json:"type"`
}
