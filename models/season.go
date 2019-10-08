package models

var (
	SeasonsEndpoint = "/seasons"
)

type Season struct {
	Asset

	Attributes SeasonAttributes `json:attributes`
}

type SeasonsList struct {
	APIResponse

	Data []Season `json:data`
}

type SeasonAttributes struct {
	IsCurrentSeason bool `json:isCurrentSeason`
	IsOffSeason     bool `json:isOffSeason`
}

type PlayerSeason struct {
	APIResponse

	Data PlayerSeasonData `json:data`
}

type PlayerSeasonData struct {
	Type          string                    `json:type`
	Attributes    PlayerSeasonAttributes    `json:attributes`
	Relationships PlayerSeasonRelationships `json:relationships`
}

type PlayerSeasonAttributes struct {
	GameModeStats GameModeStats `json:gameModeStats`
}

type PlayerSeasonRelationships struct {
	MatchesSquad    AssetList `json:matchesSquad,omitempty`
	MatchesSquadFPP AssetList `json:matchesSquadFPP,omitempty`
	MatchesSolo     AssetList `json:matchesSolo,omitempty`
	MatchesSoloFPP  AssetList `json:matchesSoloFPP,omitempty`
	MatchesDuo      AssetList `json:matchesDuo,omitempty`
	MatchesDuoFPP   AssetList `json:matchesDuoFPP,omitempty`
}
