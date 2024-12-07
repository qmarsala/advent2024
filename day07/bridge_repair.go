package day07

import "fmt"

func RunDay(name string) {
	fmt.Printf("%v wip\n", name)
}

func SumSolvableEquations(input map[int64][]int64) (sum int64) {
	for k, v := range input {
		if Sum(v) == k {
			sum += Sum(v)
		}
		if Product(v) == k {
			sum += Sum(v)
		}
	}
	return
}

func Sum(numbers []int64) (sum int64) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func Product(numbers []int64) (product int64) {
	product = 1
	for _, n := range numbers {
		product *= n
	}
	return
}
