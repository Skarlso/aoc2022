package main

import (
	"fmt"
	"os"
	"strings"
)

type blueprint struct {
	id         int
	oreRobot   oreRobot
	clayRobot  clayRobot
	obsRobot   obsRobot
	geodeRobot geodeRobot
}

type oreRobot struct {
	oreCost int
}

type clayRobot struct {
	oreCost int
}

type obsRobot struct {
	oreCost, clayCost int
}

type geodeRobot struct {
	oreCost, obsCost int
}

var globalBest int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]
	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	blueprints := make([]blueprint, 0)

	for _, line := range split {
		if line == "" {
			continue
		}

		var (
			id         int
			oreRobot   oreRobot
			clayRobot  clayRobot
			obsRobot   obsRobot
			geodeRobot geodeRobot
		)

		fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreRobot.oreCost, &clayRobot.oreCost, &obsRobot.oreCost, &obsRobot.clayCost, &geodeRobot.oreCost, &geodeRobot.obsCost)

		bp := blueprint{
			id:         id,
			oreRobot:   oreRobot,
			clayRobot:  clayRobot,
			obsRobot:   obsRobot,
			geodeRobot: geodeRobot,
		}

		blueprints = append(blueprints, bp)
	}

	result := 0
	for _, bp := range blueprints {

		result += bp.id * search(bp, 0, 0, 0, 24, 1, 0, 0, 0, 0)
		globalBest = 0
	}

	fmt.Println("quality level sum: ", result)
}

func search(bp blueprint, ore, clay, obs, time, oreRobots, clayRobots, obsRobots, geodeRobots, geodes int) int {
	if time == 0 || globalBest >= geodes+precalculateGeodResult(geodeRobots, geodeRobots+time-1) {
		return 0
	}
	if oreRobots >= bp.geodeRobot.oreCost && obsRobots >= bp.geodeRobot.obsCost {
		return precalculateGeodResult(geodeRobots, geodeRobots+time-1)
	}

	oreLimitHit := oreRobots >= max(bp.geodeRobot.oreCost, max(bp.clayRobot.oreCost, bp.obsRobot.oreCost))
	clayLimitHit := clayRobots >= bp.obsRobot.clayCost
	obsLimitHit := obsRobots >= bp.geodeRobot.obsCost
	best := 0

	if !oreLimitHit {
		best = max(
			best,
			geodeRobots+search(
				bp, ore+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.oreRobot.oreCost && !oreLimitHit {
		best = max(
			best,
			geodeRobots+search(
				bp, ore-bp.oreRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots+1, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.clayRobot.oreCost && !clayLimitHit {
		best = max(
			best, geodeRobots+search(
				bp, ore-bp.clayRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots+1, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.obsRobot.oreCost && clay >= bp.obsRobot.clayCost && !obsLimitHit {
		best = max(
			best, geodeRobots+search(
				bp, ore-bp.obsRobot.oreCost+oreRobots, clay-bp.obsRobot.clayCost+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots+1, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.geodeRobot.oreCost && obs >= bp.geodeRobot.obsCost {
		best = max(
			best, geodeRobots+search(
				bp, ore-bp.geodeRobot.oreCost+oreRobots, clay+clayRobots, obs-bp.geodeRobot.obsCost+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots+1, geodes+geodeRobots))
	}

	globalBest = max(best, globalBest)
	return best
}

func precalculateGeodResult(first, last int) int {
	return last*(last+1)/2 - ((first - 1) * first / 2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
