package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RED_MAX   int = 12
	GREEN_MAX int = 13
	BLUE_MAX  int = 14
)

type Game struct {
	Id   int
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func (s *Set) isValid(redMax, greenMax, blueMax int) bool {
	return s.Red <= redMax && s.Green <= greenMax && s.Blue <= blueMax
}

func minSetColors(sets []Set) (int, int, int) {
	red, green, blue := 0, 0, 0
	for _, set := range sets {
		if set.Red > red {
			red = set.Red
		}
		if set.Green > green {
			green = set.Green
		}
		if set.Blue > blue {
			blue = set.Blue
		}
	}

	return red, green, blue
}

func main() {
	file, err := os.Open("day2.txt")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	defer file.Close()

	var games []Game
	powerOfAllSets := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		gameSplit := strings.Split(line, ": ")
		gameNumberSplit := strings.Split(gameSplit[0], " ")
		gameNumber, err := strconv.Atoi(gameNumberSplit[1])
		if err != nil {
			fmt.Println("Converting game number error:", err)
			return
		}
		var game Game
		game.Id = gameNumber
		gameSets := strings.Split(gameSplit[1], "; ")
		isAllSetsValid := true
		for _, set := range gameSets {
			s := Set{}
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				cubeSplit := strings.Split(cube, " ")
				color := cubeSplit[1]
				amount, err := strconv.Atoi(cubeSplit[0])
				if err != nil {
					fmt.Println("Converting cube color amount error:", err)
					return
				}
				switch color {
				case "red":
					s.Red += amount
					break
				case "green":
					s.Green += amount
					break
				case "blue":
					s.Blue += amount
					break
				}
			}
			if !s.isValid(RED_MAX, GREEN_MAX, BLUE_MAX) {
				isAllSetsValid = false
			}
			game.Sets = append(game.Sets, s)
		}
		r, g, b := minSetColors(game.Sets)
		powerOfAllSets += r*b*g
		if isAllSetsValid {
			games = append(games, game)
		}
	}

	idSum := 0

	for _, game := range games {
		idSum += game.Id
	}
	fmt.Println("IDs: ", idSum)
	fmt.Println("Power of all sets:", powerOfAllSets)
}
