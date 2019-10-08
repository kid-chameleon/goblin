package models

var (
	PlayersEndpoint = "/players"
)

// Player is the toplevel representation of a single player
type Player struct {
	APIResponse

	Data PlayerData `json:"data"`
}

// PlayerList is the toplevel representation of multiple players
type PlayerList struct {
	APIResponse

	Data []PlayerData `json:"data"`
}

// PlayerData is a single player's data
type PlayerData struct {
	Asset

	Attributes    PlayerAttributes    `json:"attributes"`
	Relationships PlayerRelationships `json:"relationships"`
	Links         LinksWithSchema     `json:"links"`
}

// PlayerAttributes is a single player's attributes
type PlayerAttributes struct {
	Name         string      `json:"name"`
	PatchVersion string      `json:"patchVersion"`
	ShardID      string      `json:"shardId"`
	Stats        interface{} `json:"stats,omitempty"`
	TitleID      string      `json:"titleId"`
}

// PlayerRelationships is a single player's relationships
type PlayerRelationships struct {
	Assests AssetList `json:"assets,omitempty"`
	Matches AssetList `json:"matches,omitempty"`
}
