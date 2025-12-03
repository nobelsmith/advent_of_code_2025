package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func getFactorPairs(n int) [][2]int {
	pairs := [][2]int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 && i != 1 {
			pairs = append(pairs, [2]int{i, n / i})
		}
	}
	return pairs
}

func checkIfPairsAreSame(list []string) bool {
	first := list[0]
	for _, s := range list {
		if s != first {
			return false
		}
	}
	return true
}

func main() {
	var invalidIds []int
	in := `3299143-3378031,97290-131156,525485-660941,7606-10180,961703-1031105,6856273537-6856492968,403537-451118,5330-7241,274725-384313,27212572-27307438,926609-954003,3035-3822,161-238,22625-31241,38327962-38415781,778-1155,141513-192427,2-14,47639-60595,4745616404-4745679582,1296-1852,80-102,284-392,4207561-4292448,404-483,708177-776613,65404-87389,5757541911-5757673432,21-38,485-731,1328256-1444696,11453498-11629572,41-66,2147-3014,714670445-714760965,531505304-531554460,4029-5268,3131222053-3131390224`
	// in := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	r := csv.NewReader(strings.NewReader(in))

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, data := range line {
			bounds := strings.Split(data, "-")
			lowerBound, err := strconv.Atoi(bounds[0])
			if err != nil {
				log.Fatal(err)
			}
			upperBound, err := strconv.Atoi(bounds[1])
			if err != nil {
				log.Fatal(err)
			}

			for j := lowerBound; j < upperBound+1; j++ {
				candidate := strconv.Itoa(j)
				var candidateIsInvalid bool
				// check to see if all of the digits are the same
				if strings.Repeat(string(candidate[0]), len(candidate)) == candidate {
					candidateIsInvalid = true
					continue
				}
				// check all factor pairs
				factorPairs := getFactorPairs(len(candidate))
				for _, pair := range factorPairs {
					fmt.Println("candidate", candidate, "candidate length:", len(candidate), "pair:", pair)
					// check the first number the second number of times
					var comparablePairs []string
					var startPosition = 0
					for i := 1; i <= pair[1]; i++ {
						endPosition := startPosition + pair[0]
						fmt.Println("staring:", startPosition, "ending", endPosition)
						comparablePairs = append(comparablePairs, candidate[startPosition:endPosition])
						startPosition = endPosition
					}
					if checkIfPairsAreSame(comparablePairs) {
						candidateIsInvalid = true
						continue
					}
					if pair[0] != pair[1] {
						// check the second number the first number of times
						comparablePairs = nil
						startPosition = 0
						for i := 1; i <= pair[0]; i++ {
							endPosition := startPosition + pair[1]
							fmt.Println("staring:", startPosition, "ending", endPosition)
							comparablePairs = append(comparablePairs, candidate[startPosition:endPosition])
							startPosition = endPosition
						}
						if checkIfPairsAreSame(comparablePairs) {
							candidateIsInvalid = true
						}
					}
				}
				if candidateIsInvalid {
					invalidIds = append(invalidIds, j)
				}
			}
		}
	}
	// Sum up all of the invalid IDS
	var result int
	for i := 0; i < len(invalidIds); i++ {
		result += invalidIds[i]
	}
	fmt.Println("result: ", result)
}

// first guess 47477054026 too high
// second guess 47470921690 too low