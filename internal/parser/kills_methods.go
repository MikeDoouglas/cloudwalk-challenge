package parser

type DeathType string

const (
	Unknown       DeathType = "MOD_UNKNOWN"
	Shotgun       DeathType = "MOD_SHOTGUN"
	Gauntlet      DeathType = "MOD_GAUNTLET"
	Machinegun    DeathType = "MOD_MACHINEGUN"
	Grenade       DeathType = "MOD_GRENADE"
	GrenadeSplash DeathType = "MOD_GRENADE_SPLASH"
	Rocket        DeathType = "MOD_ROCKET"
	RocketSplash  DeathType = "MOD_ROCKET_SPLASH"
	Plasma        DeathType = "MOD_PLASMA"
	PlasmaSplash  DeathType = "MOD_PLASMA_SPLASH"
	Railgun       DeathType = "MOD_RAILGUN"
	Lightning     DeathType = "MOD_LIGHTNING"
	BFG           DeathType = "MOD_BFG"
	BFGSplash     DeathType = "MOD_BFG_SPLASH"
	Water         DeathType = "MOD_WATER"
	Slime         DeathType = "MOD_SLIME"
	Lava          DeathType = "MOD_LAVA"
	Crush         DeathType = "MOD_CRUSH"
	Telefrag      DeathType = "MOD_TELEFRAG"
	Falling       DeathType = "MOD_FALLING"
	Suicide       DeathType = "MOD_SUICIDE"
	TargetLaser   DeathType = "MOD_TARGET_LASER"
	TriggerHurt   DeathType = "MOD_TRIGGER_HURT"
	Nail          DeathType = "MOD_NAIL"
	Chaingun      DeathType = "MOD_CHAINGUN"
	ProximityMine DeathType = "MOD_PROXIMITY_MINE"
	Kamikaze      DeathType = "MOD_KAMIKAZE"
	Juiced        DeathType = "MOD_JUICED"
	Grapple       DeathType = "MOD_GRAPPLE"
)

var weaponTypeMap = map[string]DeathType{
	"MOD_SHOTGUN":        Shotgun,
	"MOD_GAUNTLET":       Gauntlet,
	"MOD_MACHINEGUN":     Machinegun,
	"MOD_GRENADE":        Grenade,
	"MOD_GRENADE_SPLASH": GrenadeSplash,
	"MOD_ROCKET":         Rocket,
	"MOD_ROCKET_SPLASH":  RocketSplash,
	"MOD_PLASMA":         Plasma,
	"MOD_PLASMA_SPLASH":  PlasmaSplash,
	"MOD_RAILGUN":        Railgun,
	"MOD_LIGHTNING":      Lightning,
	"MOD_BFG":            BFG,
	"MOD_BFG_SPLASH":     BFGSplash,
	"MOD_WATER":          Water,
	"MOD_SLIME":          Slime,
	"MOD_LAVA":           Lava,
	"MOD_CRUSH":          Crush,
	"MOD_TELEFRAG":       Telefrag,
	"MOD_FALLING":        Falling,
	"MOD_SUICIDE":        Suicide,
	"MOD_TARGET_LASER":   TargetLaser,
	"MOD_TRIGGER_HURT":   TriggerHurt,
	"MOD_NAIL":           Nail,
	"MOD_CHAINGUN":       Chaingun,
	"MOD_PROXIMITY_MINE": ProximityMine,
	"MOD_KAMIKAZE":       Kamikaze,
	"MOD_JUICED":         Juiced,
	"MOD_GRAPPLE":        Grapple,
}

func GetDeathType(mod string) DeathType {
	if weaponType, exists := weaponTypeMap[mod]; exists {
		return weaponType
	}
	return Unknown
}
