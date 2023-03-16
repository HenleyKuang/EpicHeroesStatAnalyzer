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
func PassiveAtkDmg(basicAtk int, atkBuff float32, brokenArmor float32, passiveAtk int) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	passiveDmgMultiplier := float32(passiveAtk / 100)
	return int(float32(basicAtk) * (1.0 + atkBuff) * (1.0 + brokenArmor) * passiveDmgMultiplier)
}

// PassiveAtkCritDmg computes the passive atk damage with crit. a(1+A)(1+x+xy)(1+P)B
func PassiveAtkCritDmg(basicAtk int, atkBuff float32, brokenArmor float32, passiveAtk int, crit float32, critDmg float32) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return int(float32(PassiveAtkDmg(basicAtk, atkBuff, brokenArmor, passiveAtk)) * (1 + crit + crit*critDmg))
}

// BasicAtkDmg computes the basic atk damage. a(1+A)(1+P)
func BasicAtkDmg(basicAtk int, atkBuff float32, brokenArmor float32) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	return int(float32(basicAtk) * (1.0 + atkBuff) * (1.0 + brokenArmor))
}

// BasicAtkCritDmg computes the basic atk damage with crit. a(1+A)(1+x+xy)(1+P)
func BasicAtkCritDmg(basicAtk int, atkBuff float32, brokenArmor float32, crit float32, critDmg float32) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return int(float32(BasicAtkDmg(basicAtk, atkBuff, brokenArmor)) * (1.0 + crit + crit*critDmg))
}

// SkillAtkDmg computes the skill atk damage. a(1+A)(1+J+T)(1+P)T
func SkillAtkDmg(basicAtk int, atkBuff float32, brokenArmor float32, skillDmg float32, skillAtk int) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	skillDmg = skillDmg / 100
	skillDmgMultiplier := float32(skillAtk / 100)
	return int(float32(basicAtk) * (1.0 + atkBuff) * (1.0 + brokenArmor) * (1.0 + skillDmg + skillDmgMultiplier) * skillDmgMultiplier)
}

// SkillAtkCritDmg computes the skill atk damage with crit. a(1+A)(1+x+xy)(1+J+T)(1+P)T
func SkillAtkCritDmg(basicAtk int, atkBuff float32, brokenArmor float32, skillDmg float32, skillAtk int, crit float32, critDmg float32) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return int(float32(SkillAtkDmg(basicAtk, atkBuff, brokenArmor, skillDmg, skillAtk)) * (1.0 + crit + crit*critDmg))
}
