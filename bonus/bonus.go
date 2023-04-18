package bonus

import (
	"fmt"
)

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	if characterLevel <= 0 && enemyLevel <= 0 {
		return 0, fmt.Errorf("nível inválido")
	}
	var damage int
	var diferenca int
	diferenca = characterLevel - enemyLevel
	if diferenca > 20 {
		damage = characterLevel * 5
	} else if characterLevel > enemyLevel {
		damage = characterLevel * 10
	}
	if diferenca < 20 {
		damage = characterLevel * 2
	} else if characterLevel < enemyLevel {
		damage = characterLevel * 5
	}
	if characterLevel == enemyLevel {
		damage = characterLevel * 7
	}
	if characterLevel > 100 {
		damage = characterLevel * 2
	}
	return damage, nil
}
