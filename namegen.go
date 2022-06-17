package namegen

import (
	"errors"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate one returns one generated name
func GenerateOne(list1 []string, list2[]string)(string, error){
	separator:="-"
    if len(list1) == 0 {
        return "", errors.New("list 1 is empty, please provide some lists to build your names")
    }
    if len(list2) == 0 {
        return "", errors.New("list 2 is empty, please provide the list for the second part of your names")
    }

	part1:=list1[rand.Intn(len(list1))]
	part2:=list2[rand.Intn(len(list2))]

	return part1 + separator + part2, nil 
}