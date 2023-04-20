package bonus

import (
	"fmt"
)

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	if characterLevel == 0 || enemyLevel == 0 {
		return 0, fmt.Errorf("nível inválido")
	}
	var damage int

	if characterLevel > 100 {
		damage = characterLevel * 20
	} else if enemyLevel > 100 {
		damage = characterLevel * 2
	} else if characterLevel > enemyLevel {
		damage = characterLevel * 10
	} else if characterLevel < enemyLevel {
		damage = characterLevel * 5
	} else {
		damage = characterLevel * 7
	}
	if dif := characterLevel - enemyLevel; dif > 20 {
		damage *= 5
	} else if dif < -20 {
		damage *= 2
	}
	return damage, nil
}
