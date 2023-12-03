package main

import (
	"testing"
)

type TestPair struct {
  expected string
  actual string
}

var testCases = []TestPair{
  { expected: "two1nine", actual: "29"},
  { expected: "eightwothree", actual: "83"},
  { expected: "abcone2threexyz", actual: "13"},
  { expected: "xtwone3four", actual: "24"},
  { expected: "4nineeightseven2", actual: "42"},
  { expected: "zoneight234", actual: "14"},
  { expected: "7pqrstsixteen", actual: "76"},
  { expected: "4rcs6bhbbgzhsstwomnineksbxfzj8", actual: "48"},
  { expected: "25gmh12threeltfnfdrxhh5", actual: "25"},
  { expected: "57four", actual: "54"},
  { expected: "mqgdhfour67", actual: "47"},
  { expected: "37ninetxkddhfive", actual: "35"},
  { expected: "rzrsskzrlzjbcgthreeghbqhdpxfvgjfqclcf4", actual: "34"},
  { expected: "fourvone2vbpltlrj", actual: "42"},
  { expected: "xz5four3nineseven", actual: "57"},
  { expected: "1szrhcmzkftwo9eight2ltjmgjzcblzone", actual: "11"},
  { expected: "zlnkddtgsb1sixsxvjxgxp2", actual: "12"},
  { expected: "26sixpzpsixtwozqff", actual: "22"},
  { expected: "seven99fzqxfmttfgxm", actual: "79"},
  { expected: "9twonineonefourpttbgkxt8two", actual: "92"},
  { expected: "fv9", actual: "99"},
  { expected: "5qcmjsfk6zxjld1", actual: "51"},
  { expected: "fkjstnvmchsr9q699", actual: "99"},
  { expected: "nine78three", actual: "93"},
}

func TestDetectingStringNumberValues(t *testing.T) {
  for _, pair := range testCases {
    v := GetNumbersFromStrings(pair.expected)
    if v != pair.actual {
      t.Error(
        "For", pair.expected,
        "expected", pair.actual,
        "got", v,
      )
    }
  }
}
