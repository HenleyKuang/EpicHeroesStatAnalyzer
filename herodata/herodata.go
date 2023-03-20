package herodata

// HeroDmg is a struct that contains the dmg pcts for a hero.
type HeroDmg struct {
	BasicDmgPct   int
	PassiveDmgPct int
	SkillDmgPct   int
}

// nolint:all
const (
	SamuraiGirl = "toko"
	GodSlayer   = "gs"
)

// HeroDmgsMap is a mapping of all heroes to their dmg map.
var HeroDmgsMap = map[string]*HeroDmg{
	SamuraiGirl: heroDmgObj(
		/* BasicDmgPct */ 100,
		/* PassiveDmgPct */ 85,
		/* SkillDmgPct */ 240,
	),
	GodSlayer: heroDmgObj(
		/* BasicDmgPct */ 200,
		/* PassiveDmgPct */ 200,
		/* SkillDmgPct */ 0,
	),
}

// AllHeroes is a list of all the heroes that we have dmg data for.
var AllHeroes []string

func init() {
	for heroName, _ := range HeroDmgsMap {
		AllHeroes = append(AllHeroes, heroName)
	}
}

func heroDmgObj(basicDmgPct int, passiveDmgPct int, skilldmgPct int) *HeroDmg {
	return &HeroDmg{
		BasicDmgPct:   basicDmgPct,
		PassiveDmgPct: passiveDmgPct,
		SkillDmgPct:   skilldmgPct,
	}
}
