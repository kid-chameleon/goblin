package models

import "time"

var (
	MatchesEndpoint = "/matches"
)

// Match is a single match response
type Match struct {
	APIResponse

	Data     MatchData  `json:"data"`
	Included []Included `json:"included"`
}

// MatchAttributes is a single Match's attributes
type MatchAttributes struct {
	CreatedAt     time.Time   `json:"createdAt"`
	Duration      int         `json:"duration"`
	GameMode      string      `json:"gameMode"`
	MapName       string      `json:"mapName"`
	IsCustomMatch bool        `json:"isCustomMatch"`
	PatchVersion  string      `json:"patchVersion"`
	SeasonState   string      `json:"seasonState"`
	ShardID       string      `json:"shardId"`
	Stats         interface{} `json:"stats,omitempty"`
	Tags          Tags        `json:"tags,omitempty"`
	TitleID       string      `json:"titleId"`
}

type MatchRelationships struct {
	Assets     AssetList `json:"assets,omitempty"`
	Rosters    AssetList `json:"rosters,omitempty"`
	Rounds     Asset     `json:"rounds,omitempty"`
	Spectators Asset     `json:"rounds,omitempty"`
}

type MatchData struct {
	Asset

	Attributes    MatchAttributes    `json:"attributes"`
	Relationships MatchRelationships `json:"relationships"`
	Links         LinksWithSchema    `json:"links"`
}

type Included struct {
	Asset

	Attributes    IncludedAttributes    `json:"attributes"`
	Relationships IncludedRelationships `json:"relationships"`
}

type IncludedAttributes struct {
	Stats   Stats  `json:"stats"`
	Actor   string `json:"actor,omitempty"`
	ShardID string `json:"shardId,omitempty"`
	Won     string `json:"won,omitempty"`
}

type IncludedRelationships struct {
	Participants AssetList `json:"participants,omitempty"`
	Team         Asset     `json:"team,omitempty"`
}
