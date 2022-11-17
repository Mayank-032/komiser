package models

import "time"

type Resource struct {
	Id         int64             `json:"id" bun:"id,pk,autoincrement"`
	ResourceId string            `json:"resourceId" bun:"resourceId,unique"`
	Provider   string            `json:"provider"`
	Account    string            `json:"account"`
	AccountId  string            `json:"accountId"`
	Service    string            `json:"service"`
	Region     string            `json:"region"`
	Name       string            `json:"name"`
	CreatedAt  time.Time         `json:"createdAt"`
	FetchedAt  time.Time         `json:"fetchedAt"`
	Cost       float64           `json:"cost"`
	Metadata   map[string]string `json:"metadata"`
	Tags       []Tag             `json:"tags" bun:"tags"`
	Link       string            `json:"link" bson:"link"`
	Value      string            //to be deprecated
}

type Tag struct {
	Key   string `json:"key" bun:"key"`
	Value string `json:"value" bun:"value"`
}

type BulkUpdateTag struct {
	Tags      []Tag `json:"tags"`
	Resources []int `json:"resources"`
}
