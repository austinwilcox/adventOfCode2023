package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
  file, err := os.ReadFile("./input.txt"); if err != nil {
    log.Fatal(err)
  }
  fileContents := string(file)

  grid := make([][]string, len(strings.Split(fileContents, "\n")) - 1)
  for i:=0; i < len(strings.Split(fileContents, "\n"))-1; i++ {
    line := strings.Split(fileContents, "\n")[i]
    grid[i] = make([]string, len(line))
    for j:= 0; j < len(grid[i]); j++ {
      grid[i][j] = string(line[j])
    }
  }

  fmt.Println(PartOne(grid))
}

func PartOne(grid [][]string) int {
  sum := 0
  ratios := 0
  var specialPoints []SpecialPoint
  hashValues := make(map[int]int, 100)
  for y:=0; y < len(grid); y++ {
    for x:=0; x < len(grid[y])-1; x++ {
      val, err :=  strconv.Atoi(grid[y][x]); if err != nil {
        continue
      }

      fullVal := fmt.Sprint(val)
      addValueToSum := false
      var point Point
      for {
        haveASymbol, samplePoint := doSiblingsHaveASymbol(x, y, grid)
        if haveASymbol {
          addValueToSum = true
          point = samplePoint
          // points = append(points, point)
        }
        x++
        if x > len(grid[y])-1 {
          break
        }
        temp, err :=  strconv.Atoi(grid[y][x]); if err != nil {
          break
        }
        fullVal += fmt.Sprint(temp)
      }

      if addValueToSum {
        valueToAdd, err :=  strconv.Atoi(fullVal); if err != nil {
          break
        }
        cantor := cantorsPairing(float64(point.x), float64(point.y))
        if hashValues[cantor] != 0 {
          ratios += hashValues[cantor] * valueToAdd
        } else {
          hashValues[cantor] = valueToAdd
          specialPoints = append(specialPoints, SpecialPoint{p: point, val: fullVal})
        }
        sum += valueToAdd
      }
    }
  }

  fmt.Println(ratios)
  return sum
}

type SpecialPoint struct {
  p Point
  val string
}

type Point struct {
  x int
  y int
}

func doSiblingsHaveASymbol(x int, y int, grid [][]string) (bool, Point) {
  symbols := []string{
    "#","+", "=", "_", "/", "-", "=", "%", "*", "@", "$", "&",
  }
  // var options []Test
  for i:=y-1; i <= y+1; i++ {
    for j:=x-1; j <= x+1; j++ {
      if i == y && x == j {
        continue
      }
      if i < 0 || j < 0 {
        continue
      }
      if i > len(grid)-1 {
        continue
      }
      if j > len(grid[i])-1 {
        continue
      }

      test := []rune(grid[i][j])
      for _, symbol := range symbols {
        if string(test[0]) == symbol {
          return true, Point{x: i, y: j}
        }
      }
    }
  }

  return false, Point{-1, -1}
}

func cantorsPairing(x float64, y float64) int {
  return int(math.Floor((math.Pow(x, 2.0) + (3.0*x) + (2 * x * y) + y + math.Pow(y, 2.0)) / 2))
}
