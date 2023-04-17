package dmgformula

/*
P = broken armor
J = skill damage stat
A = Attack gain from buffs or passives etc during battle
a = the attack power as shown in the stats page of the hero
x = crit
y = crit dmg
T = skill damage multiplier, for example generals main skill level 3 is 340% the damage, 340% being the multiplier. (340% = 3.4 multiplier)
B = passive damage multiplier
*/

// PassiveAtkDmg computes the passive atk damage. a(1+A)(1+P)B
func PassiveAtkDmg(baseAtk int, atkBuff float32, brokenArmor float32, passiveAtkPct int, dmgImmune float32) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	passiveDmgMultiplier := float32(passiveAtkPct) / 100.0
	return int(float32(baseAtk) * (1.0 + atkBuff) * (1.0 + brokenArmor) * passiveDmgMultiplier)
}

// PassiveAtkCritDmg computes the passive atk damage with crit. a(1+A)(1+x+xy)(1+P)B
func PassiveAtkCritDmg(baseAtk int, atkBuff float32, brokenArmor float32, passiveAtkPct int, crit float32, critDmg float32, dmgImmune float32) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return int(float32(PassiveAtkDmg(baseAtk, atkBuff, brokenArmor, passiveAtkPct, dmgImmune)) * (1 + crit + crit*critDmg))
}

// BasicAtkDmg computes the basic atk damage. a(1+A)(1+P)
func BasicAtkDmg(baseAtk int, atkBuff float32, brokenArmor float32, basicAtkPct int, dmgImmune float32) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	basicDmgMultiplier := float32(basicAtkPct) / 100.0
	return int(float32(baseAtk) * (1.0 + atkBuff) * (1.0 + brokenArmor) * basicDmgMultiplier)
}

// BasicAtkCritDmg computes the basic atk damage with crit. a(1+A)(1+x+xy)(1+P)
func BasicAtkCritDmg(baseAtk int, atkBuff float32, brokenArmor float32, basicAtkPct int, crit float32, critDmg float32, dmgImmune float32) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return int(float32(BasicAtkDmg(baseAtk, atkBuff, brokenArmor, basicAtkPct, dmgImmune)) * (1.0 + crit + crit*critDmg))
}

// SkillAtkDmg computes the skill atk damage. a(1+A)(1+J+T)(1+P)T
func SkillAtkDmg(baseAtk int, atkBuff float32, brokenArmor float32, skillDmg float32, skillAtkPct int, dmgImmune float32) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	skillDmg = skillDmg / 100
	skillDmgMultiplier := float32(skillAtkPct) / 100.0
	return int(float32(baseAtk) * (1.0 + atkBuff) * (1.0 + brokenArmor) * (skillDmg + skillDmgMultiplier))
}

// SkillAtkCritDmg computes the skill atk damage with crit. a(1+A)(1+x+xy)(1+J+T)(1+P)T
func SkillAtkCritDmg(baseAtk int, atkBuff float32, brokenArmor float32, skillDmg float32, skillAtkPct int, crit float32, critDmg float32, dmgImmune float32) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return int(float32(SkillAtkDmg(baseAtk, atkBuff, brokenArmor, skillDmg, skillAtkPct, dmgImmune)) * (1.0 + crit + crit*critDmg))
}
