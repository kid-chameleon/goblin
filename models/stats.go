package models

type Stats struct {
	Assists         int     `json:"assists,omitempty"`
	Boosts          int     `json:"boosts,omitempty"`
	DamageDealt     float32 `json:"damageDealt,omitempty"`
	DBNOs           int     `json:"DBNOs,omitempty"`
	DeathType       string  `json:"deathType,omitempty"`
	HeadshotKills   int     `json:"headshotKills,omitempty"`
	Heals           int     `json:"heals,omitempty"`
	KillPlace       int     `json:"killPlace,omitempty"`
	KillPoints      int     `json:"killPoints,omitempty"`
	KillPointsDelta float32 `json:"killPointsDelta,omitempty"`
	KillStreaks     int     `json:"killStreaks,omitempty"`
	Kills           int     `json:"kills,omitempty"`
	LastKillPoints  int     `json:"lastKillPoints,omitempty"`
	LastWinPoints   int     `json:"lastWinPoints,omitempty"`
	LongestKill     float32 `json:"longestKill,omitempty"`
	MostDamage      int     `json:"mostDamage,omitempty"`
	Name            string  `json:"name,omitempty"`
	PlayerID        string  `json:"playerId,omitempty"`
	Rank            int     `json:"rank,omitempty"`
	Revives         int     `json:"revives,omitempty"`
	RideDistance    float32 `json:"rideDistance,omitempty"`
	RoadKills       int     `json:"roadKills,omitempty"`
	TeamID          int     `json:"teamId,omitempty"`
	TeamKills       int     `json:"teamKills,omitempty"`
	TimeSurvived    float32 `json:"timeSurvived,omitempty"`
	VehicleDestroys int     `json:"vehicleDestroys,omitempty"`
	WalkDistance    float32 `json:"walkDistance,omitempty"`
	WeaponsAcquired int     `json:"weaponsAcquired,omitempty"`
	WinPlace        int     `json:"winPlace,omitempty"`
	WinPoints       int     `json:"winPoints,omitempty"`
	WinPointsDelta  float32 `json:"winPointsDelta,omitempty"`
}

type GameModeStats struct {
	Duo      Stats `json:duo`
	DuoFPP   Stats `json:duo-fpp`
	Solo     Stats `json:solo`
	SoloFPP  Stats `json:solo-fpp`
	Squad    Stats `json:squad`
	SquadFPP Stats `json:squad-fpp`
}
