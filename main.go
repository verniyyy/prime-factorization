package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	inputData := 139819680

	data := NewData(int64(inputData))
	for key := range data.PrimeNumbers {
		Recursive(key, data)
	}
	data.Formula()

	return
}

func SysAbend(err error) {
	exitStatus := 2
	fmt.Printf("Error { %v }\n", err)
	os.Exit(exitStatus)
}

type Data struct {
	PrimeNumbers map[int64]int64
	inputData    int64
	MainData     int64
}

func getKeys(m map[int64]int64) []int {
	var keys []int
	for k := range m {
		keys = append(keys, int(k))
	}
	return keys
}

func NewData(inputData int64) *Data {
	return &Data{
		map[int64]int64{2: 0, 3: 0, 5: 0, 7: 0, 11: 0, 13: 0, 17: 0, 23: 0, 29: 0, 31: 0, 37: 0, 41: 0},
		inputData,
		inputData,
	}
}

func (d *Data) ViewingDivision(n int64) error {
	d.Println("befor")

	err := d.Division(n)
	if err != nil {
		return err
	}

	d.Println("after")
	fmt.Println()

	return nil
}

func (d *Data) Formula() error {
	keys := getKeys(d.PrimeNumbers)
	sort.Ints(keys)

	fmt.Printf("%v =", d.inputData)

	headTerm := true
	var value int64
	var key int64
	for i := range keys {
		key = int64(keys[i])
		value = d.PrimeNumbers[int64(key)]

		if value == 0 {
			continue
		}
		if headTerm {
			fmt.Printf(" %v ^ %v", key, value)
			headTerm = false
			continue
		}
		fmt.Printf(" + %v ^ %v", key, value)
	}

	if d.MainData != 1 {
		fmt.Printf(" + %v", d.MainData)
	}

	fmt.Println()
	return nil
}

func (d *Data) Println(prefix string) error {
	fmt.Printf("%s: %v\n", prefix, d)
	return nil
}

func (d *Data) Division(n int64) error {
	_, exsists := d.PrimeNumbers[n]
	if !exsists {
		return fmt.Errorf("invalid number: n=%v", n)
	}

	err := d.division(n)
	if err != nil {
		return err
	}

	d.countUpPrimeNumber(n)

	return nil
}

func (d *Data) division(n int64) error {
	if d.MainData%n != 0 {
		return fmt.Errorf("DivisionError")
	}
	d.MainData /= n
	return nil
}

func (d *Data) countUpPrimeNumber(p int64) error {
	d.PrimeNumbers[p] += 1
	return nil
}

func Recursive(n int64, data *Data) (*Data, error) {
	err := data.Division(n)
	if err != nil {
		return nil, err
	}

	_, err = Recursive(n, data)
	if err != nil {
		return data, nil
	}

	return data, err
}
