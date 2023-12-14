package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  file, err := os.ReadFile("./input.txt"); if err != nil {
    log.Fatal(err)
  }
  lines := strings.Split(string(file), "\n")
  
  // finalSum := 0
  dp := make(map[int]int, 100)
  for _, line := range lines {
    if line == "" {
      continue
    }

    // finalSum += PartOne(line)
    PartTwo(line, dp)
  }
  sum := 0
  for _, v := range dp {
    sum += v
  }
  fmt.Println(sum)

  // fmt.Println(finalSum)
}

func PartOne(line string) int {
  ht := make(map[int]int, 100)

  splitStr := strings.Split(line, ":")
  gameNumbers := strings.Split(splitStr[1], "|")
  winningNumbers := gameNumbers[0]
  myNumbers := gameNumbers[1]

  for _, str := range strings.Split(strings.Trim(winningNumbers, " "), " ") {
    number, err := strconv.Atoi(str); if err != nil {
      continue
    }

    ht[number] = 1
  }

  sum := 0
  numberOfCards := 0
  for _, str := range strings.Split(strings.Trim(myNumbers, " "), " ") {
    number, err := strconv.Atoi(str); if err != nil {
      continue
    }
    if ht[number] == 1 {
      numberOfCards++
      if sum == 0 {
        sum = 1
      } else {
        sum *= 2
      }
    }

    ht[number] = 1
  }


  return sum
}

func PartTwo(line string, dp map[int]int) []int {
  ht := make(map[int]int, 100)
  var regex = regexp.MustCompile(`Card\s+(\d+)`)
  cardId, _ := strconv.Atoi(regex.FindStringSubmatch(line)[1])
  dp[cardId] += 1

  splitStr := strings.Split(line, ":")
  gameNumbers := strings.Split(splitStr[1], "|")
  winningNumbers := gameNumbers[0]
  myNumbers := gameNumbers[1]

  for _, str := range strings.Split(strings.Trim(winningNumbers, " "), " ") {
    number, err := strconv.Atoi(str); if err != nil {
      continue
    }

    ht[number] = 1
  }

  numberOfCards := cardId
  for _, str := range strings.Split(strings.Trim(myNumbers, " "), " ") {
    number, err := strconv.Atoi(str); if err != nil {
      continue
    }
    if ht[number] == 1 {
      numberOfCards++
      dp[numberOfCards] += dp[cardId]
    }

    ht[number] = 1
  }

  fmt.Println(cardId, dp);


  return []int{}
}
