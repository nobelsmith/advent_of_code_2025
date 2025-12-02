package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
	var invalidIds []int
	in := `3299143-3378031,97290-131156,525485-660941,7606-10180,961703-1031105,6856273537-6856492968,403537-451118,5330-7241,274725-384313,27212572-27307438,926609-954003,3035-3822,161-238,22625-31241,38327962-38415781,778-1155,141513-192427,2-14,47639-60595,4745616404-4745679582,1296-1852,80-102,284-392,4207561-4292448,404-483,708177-776613,65404-87389,5757541911-5757673432,21-38,485-731,1328256-1444696,11453498-11629572,41-66,2147-3014,714670445-714760965,531505304-531554460,4029-5268,3131222053-3131390224`
	// in := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862`
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
				if len(candidate)%2 == 0 {
					half := len(candidate) / 2
					firstHalf := candidate[half:]
					secondHalf := candidate[:half]
					if firstHalf == secondHalf {
						invalidIds = append(invalidIds, j)
					}
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
