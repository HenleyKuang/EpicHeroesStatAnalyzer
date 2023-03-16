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
func PassiveAtkDmg(basicAtk int, atkBuff int, brokenArmor int, passiveDmgMultiplier int) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	passiveDmgMultiplier = passiveDmgMultiplier / 100
	return basicAtk * (1 + atkBuff) * (1 + brokenArmor) * passiveDmgMultiplier
}

// PassiveAtkCritDmg computes the passive atk damage with crit. a(1+A)(1+x+xy)(1+P)B
func PassiveAtkCritDmg(basicAtk int, atkBuff int, brokenArmor int, passiveDmgMultiplier int, crit int, critDmg int) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return PassiveAtkDmg(basicAtk, atkBuff, brokenArmor, passiveDmgMultiplier) * (1 + crit + crit*critDmg)
}

// BasicAtkDmg computes the basic atk damage. a(1+A)(1+P)
func BasicAtkDmg(basicAtk int, atkBuff int, brokenArmor int) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	return basicAtk * (1 + atkBuff) * (1 + brokenArmor)
}

// BasicAtkCritDmg computes the basic atk damage with crit. a(1+A)(1+x+xy)(1+P)
func BasicAtkCritDmg(basicAtk int, atkBuff int, brokenArmor int, crit int, critDmg int) int {
	crit = crit / 100
	critDmg = critDmg / 100
	return BasicAtkDmg(basicAtk, atkBuff, brokenArmor) * (1 + crit + crit*critDmg)
}

// SkillAtkDmg computes the skill atk damage. a(1+A)(1+J+T)(1+P)T
func SkillAtkDmg(basicAtk int, atkBuff int, brokenArmor int, skillDmg int, skillDmgMultiplier int) int {
	atkBuff = atkBuff / 100
	brokenArmor = brokenArmor / 100
	skillDmg = skillDmg / 100
	skillDmgMultiplier = skillDmgMultiplier / 100
	return basicAtk * (1 + atkBuff) * (1 + brokenArmor) * (1 + skillDmg + skillDmgMultiplier) * skillDmgMultiplier
}

// SkillAtkCritDmg computes the skill atk damage with crit. a(1+A)(1+x+xy)(1+J+T)(1+P)T
func SkillAtkCritDmg(basicAtk int, atkBuff int, brokenArmor int, skillDmg int, skillDmgMultiplier int, crit int, critDmg int) int {
	return SkillAtkDmg(basicAtk, atkBuff, brokenArmor, skillDmg, skillDmgMultiplier) * (1 + crit + crit*critDmg)
}
