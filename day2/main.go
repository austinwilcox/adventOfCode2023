package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//NOTE: Which games would be possible in the list if there are only:
//12 red cubes
//13 green cubes
//14 blue cubes
func main() {
  example, err := os.ReadFile("input.txt"); if err != nil {
    log.Fatal(err)
  }
  splitArr := strings.Split(string(example), "\n")
  sum := 0
  for _, line := range splitArr {
    if line == "" {
      continue
    }
    sum += PartTwo(line)
  }

  fmt.Println(sum)
}

func PartOne(line string) int {
  dp := make(map[string]int, 3)
  dp["red"] = 12
  dp["green"] = 13
  dp["blue"] = 14

  var regex = regexp.MustCompile(`^Game (\d+)`)
  gameId, _ := strconv.Atoi(regex.FindStringSubmatch(line)[1])

  splitArr := strings.Split(strings.Split(line, ":")[1], ";")
  for _, game := range splitArr {
    for _, cube := range strings.Split(game, ",") {
      cube = strings.Trim(cube, " ")
      values := strings.Split(cube, " ")
      intVal, _ := strconv.Atoi(values[0])
      if dp[values[1]] < intVal {
        return 0
      }
    }
  }

  return gameId
}

func PartTwo(line string) int {
  dp := make(map[string]int, 3)
  dp["red"] = 0
  dp["green"] = 0
  dp["blue"] = 0

  splitArr := strings.Split(strings.Split(line, ":")[1], ";")
  for _, game := range splitArr {
    for _, cube := range strings.Split(game, ",") {
      cube = strings.Trim(cube, " ")
      values := strings.Split(cube, " ")
      intVal, _ := strconv.Atoi(values[0])
      if dp[values[1]] < intVal {
        dp[values[1]] = intVal
      }
    }
  }

  return dp["red"] * dp["blue"] * dp["green"]
}
