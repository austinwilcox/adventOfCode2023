package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  file, err := os.ReadFile("./input.txt"); if err != nil {
    log.Fatal(err)
  }
  lines := strings.Split(string(file), "\n")
  
  finalSum := 0
  for _, line := range lines {
    if line == "" {
      continue
    }

    finalSum += PartOne(line)
  }

  fmt.Println(finalSum)
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
  for _, str := range strings.Split(strings.Trim(myNumbers, " "), " ") {
    number, err := strconv.Atoi(str); if err != nil {
      continue
    }
    if ht[number] == 1 {
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
