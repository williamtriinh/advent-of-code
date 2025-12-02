package main

import (
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/thoas/go-funk"
	"github.com/williamtriinh/advent-of-code/pkg/aoc"
	"gonum.org/v1/gonum/stat/combin"
)

// Recharge: guaranteed 202 mana by the time player casts another spell

func parseInput(input string) Entity {
	lines := aoc.SplitAndTrim(input, "\n")

	values := funk.Map(lines, func(line string) float64 {
		value, _ := strconv.ParseFloat(strings.Split(line, ": ")[1], 64)
		return value
	}).([]float64)

	return Entity{
		health: values[0],
		damage: values[1],
	}
}

func solution(input string) int {
	boss := parseInput(input)

	maximumGold := 0

	for _, weapon := range weapons {
		weaponCost := weapon.cost

		for _, armor := range armors {
			armorCost := armor.cost

			for i := 0; i <= 2; i++ {
				for _, ringIndices := range combin.Combinations(len(rings), i) {
					var ringDamage float64 = 0
					var ringArmor float64 = 0
					ringCost := 0

					for _, ringIndex := range ringIndices {
						ringDamage += rings[ringIndex].damage
						ringArmor += rings[ringIndex].armor
						ringCost += rings[ringIndex].cost
					}

					bossTurns := math.Ceil(100 / math.Max(boss.damage-(armor.armor+ringArmor), 1))
					playerTurns := math.Ceil(boss.health / math.Max((weapon.damage+ringDamage)-boss.armor, 1))

					totalCost := weaponCost + armorCost + ringCost

					if playerTurns > bossTurns && totalCost > maximumGold {
						maximumGold = totalCost
					}
				}
			}
		}
	}

	return maximumGold
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/22.txt")
	t.Log(solution(string(input)))
}
