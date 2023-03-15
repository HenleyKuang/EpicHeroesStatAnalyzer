package dmgformula

/*
P= broken armor
J= skill damage
A= attack Small
a= basic attack
X= crit
Y= crit dmg
T= skill damage multiplier (eg 250% =2.5)
B= passive damage multiplier
*/

// PassiveAtkDmg computes the passive atk damage. a(1+A)(1+P)B
func PassiveAtkDmg() int {
	return 0
}

// PassiveAtkCritDmg computes the passive atk damage with crit. a(1+A)(1+x+xy)(1+P)B
func PassiveAtkCritDmg() int {
	return 0
}

// BasicAtkDmg computes the basic atk damage. a(1+A)(1+P)
func BasicAtkDmg() int {
	return 0
}

// BasicAtkCritDmg computes the basic atk damage with crit. a(1+A)(1+x+xy)(1+P)
func BasicAtkCritDmg() int {
	return 0
}

// SkillAtkDmg computes the skill atk damage. a(1+A)(1+J+T)(1+P)T
func SkillAtkDmg() int {
	return 0
}

// SkillAtkCritDmg computes the skill atk damage with crit. a(1+A)(1+x+xy)(1+J+T)(1+P)T
func SkillAtkCritDmg() int {
	return 0
}
