package main

import (
	"fmt"
	"math/rand"
)

func getStat(p Pokemon, statName string) int {
	for _, s := range p.Stats {
		if s.Stat.Name == statName {
			return s.BaseStat
		}
	}
	return 0
}

type fighter struct {
	name    string
	hp      int
	attack  int
	defense int
	speed   int
}

func newFighter(p Pokemon) fighter {
	return fighter{
		name:    p.Name,
		hp:      getStat(p, "hp") + 60,
		attack:  getStat(p, "attack"),
		defense: getStat(p, "defense"),
		speed:   getStat(p, "speed"),
	}
}

func calcDamage(attacker, defender fighter) int {
	defense := defender.defense
	if defense < 1 {
		defense = 1
	}
	damage := 22*attacker.attack/defense + 2
	damage = damage * (85 + rand.Intn(16)) / 100
	if damage < 1 {
		damage = 1
	}
	return damage
}
func simulateBattle(a, b fighter) string {
	attacker, defender := &a, &b
	if b.speed > a.speed || (b.speed == a.speed && rand.Intn(2) == 0) {
		attacker, defender = &b, &a
	}
	fmt.Printf("%s is faster and moves first!\n", attacker.name)

	const maxRounds = 100
	for round := 1; round <= maxRounds; round++ {
		damage := calcDamage(*attacker, *defender)
		defender.hp -= damage
		if defender.hp < 0 {
			defender.hp = 0
		}
		fmt.Printf("%s hits %s for %d damage! (%s HP: %d)\n",
			attacker.name, defender.name, damage, defender.name, defender.hp)

		if defender.hp == 0 {
			fmt.Printf("%s fainted!\n", defender.name)
			return attacker.name
		}
		attacker, defender = defender, attacker
	}
	fmt.Println("The battle went on too long and ended in a draw!")
	return ""
}
