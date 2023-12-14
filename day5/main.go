package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
  file, err := os.ReadFile("./example.txt"); if err != nil {
    log.Fatal(err)
  }
  PartOne(file)

}

func PartOne(fileContents []byte) {
  var seeds []string
  // seedToSoil := make(map[int]int, 100)
  // soilToFertilizer := make(map[int]int, 100)
  // fertilizerToWater := make(map[int]int, 100)
  // waterToLight := make(map[int]int, 100)
  // lightToTemperature := make(map[int]int, 100)
  // tempToHumidity := make(map[int]int, 100)
  // humidyToLocation := make(map[int]int, 100)

  currState := ""
  for _, line := range strings.Split(string(fileContents), "\n") {
    if line == "" {
      currState = ""
      continue
    }
    switch currState {
    case "seed-to-soil map:":
      fmt.Println("correct state")
    case "soil-to-fertilizer map:":
      fmt.Println("correct state")
    case "":
      fmt.Println("correct state")
    case "":
      fmt.Println("correct state")
    case "":
      fmt.Println("correct state")
    case "":
      fmt.Println("correct state")
    case "":
      fmt.Println("correct state")
    }
    if regexp.MustCompile("seeds:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("seed-to-soil map:").MatchString(line) {
      currState = "seed-to-soil map:"
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("soil-to-fertilizer map:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("fertilizer-to-water map:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("water-to-light map:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("light-to-temperature map:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("temperature-to-humidity map:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
    if regexp.MustCompile("humidity-to-location map:").MatchString(line) {
      seedNumbers := strings.Trim(strings.Split(line, ":")[1], " ")
      seeds = strings.Split(seedNumbers, " ")
      fmt.Println(seeds)
    }
  }

  fmt.Println("HIT")
}
