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

type Entity struct {
	health float64
	damage float64
	armor  float64
}

type Equipment struct {
	name   string
	cost   int
	damage float64
	armor  float64
}

var weapons = []Equipment{
	{name: "Dagger", cost: 8, damage: 4, armor: 0},
	{name: "Shortsword", cost: 10, damage: 5, armor: 0},
	{name: "Warhammer", cost: 25, damage: 6, armor: 0},
	{name: "Longsword", cost: 40, damage: 7, armor: 0},
	{name: "Greataxe", cost: 74, damage: 8, armor: 0},
}

var armors = []Equipment{
	{name: "None", cost: 0, damage: 0, armor: 0},
	{name: "Leather", cost: 13, damage: 0, armor: 1},
	{name: "Chainmail", cost: 31, damage: 0, armor: 2},
	{name: "Splintmail", cost: 53, damage: 0, armor: 3},
	{name: "Bandedmail", cost: 75, damage: 0, armor: 4},
	{name: "Platemail", cost: 102, damage: 0, armor: 5},
}

var rings = []Equipment{
	{name: "Damage +1", cost: 25, damage: 1, armor: 0},
	{name: "Damage +2", cost: 50, damage: 2, armor: 0},
	{name: "Damage +3", cost: 100, damage: 3, armor: 0},
	{name: "Defense +1", cost: 20, damage: 0, armor: 1},
	{name: "Defense +2", cost: 40, damage: 0, armor: 2},
	{name: "Defense +3", cost: 80, damage: 0, armor: 3},
}

func parseInput(input string) Entity {
	lines := aoc.SplitAndTrim(input, "\n")

	values := funk.Map(lines, func(line string) float64 {
		value, _ := strconv.ParseFloat(strings.Split(line, ": ")[1], 64)
		return value
	}).([]float64)

	return Entity{
		health: values[0],
		damage: values[1],
		armor:  values[2],
	}
}

func solution(input string) int {
	boss := parseInput(input)

	minimumGold := math.MaxInt

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

					if playerTurns <= bossTurns && totalCost < minimumGold {
						minimumGold = totalCost
					}
				}
			}
		}
	}

	return minimumGold
}

func TestSolution(t *testing.T) {
	input, _ := os.ReadFile("../inputs/21.txt")
	t.Log(solution(string(input)))
}
