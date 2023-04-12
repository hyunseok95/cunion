package controlstructures

import (
	"fmt"
	"os"
)

func IfFunc() error {
	var x int

	if x > 0 {
		fmt.Println("if test")
	}

	fmt.Println(x)

	if err := os.Chmod("./README.md", 0644); err != nil {
		return err
	}

	f, err := os.Open("./README.md")

	if err != nil {
		return err
	}

	d, err := f.Stat()
	if err != nil {
		f.Close()
		return err
	}
	fmt.Println(d.Mode())
	return nil
}

func ForFunc() error {

	var (
		sum int
	)

	for i := 1; i <= 10; i++ {
		sum += i
	}

	oldMap := map[string]uint8{
		"first":  1,
		"second": 2,
	}

	newMap := map[uint8]string{3: "third"}

	for key, value := range oldMap {
		newMap[value] = key

	}

	fmt.Println(newMap)

	sl := []byte{1, 2, 3}

	for _, value := range sl {
		fmt.Println(value)
	}

	str := "日本\x80語"
	for pos, char := range str {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	reverseString2 := "!olleH"
	reverseString := []byte(reverseString2)

	for i, j := 0, len(reverseString)-1; i < j; i, j = i+1, j-1 {
		reverseString[i], reverseString[j] = reverseString[j], reverseString[i]
	}

	fmt.Println(string(reverseString))
	return nil
}

func SwitchFunc(i interface{}) (result error) {

	/*
		go 언어에서 스위치는 C 언어보다 더 일반적인 표현이 가능하다.
		- 표현식은 상수이거나 꼭 정수일 필요가 없다
		- 조건식이 참일때가지 위에서부터 바닥까지 검사한다.
		- if else 쓰지말고 case를 써라
	*/

	const (
		swap byte = 'a' - 'A'
	)

	var (
		val   string
		chars []byte
	)

	switch v := i.(type) {
	case string:
		val = v
	case []string:
		val = v[0]
	default:
		val = "default"
	}

	chars = []byte(val)

	fmt.Println(string(chars))

OuterLoop:
	for n := 0; n < len(chars); n++ {
		switch {
		case chars[n] == 'k':
			fmt.Println("inner break")
			continue
		case chars[n] == 'z':
			fmt.Println("outer break")
			break OuterLoop
		case chars[n] >= 'a' && chars[n] <= 'z':
			chars[n] = chars[n] - swap
		}

	}

	fmt.Println(string(chars))

	// compare function

	var (
		a []byte
		b []byte
	)

	switch i := i.(type) {
	case []string:
		a, b = []byte(i[0]), []byte(i[1])
	}

	res := func(a, b []byte) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			switch {
			case a[i] > b[i]:
				return 1
			case a[i] < b[i]:
				return -1
			}
		}
		switch {
		case len(a) > len(b):
			return 1
		case len(a) < len(b):
			return -1
		}
		return 0
	}(a, b)

	fmt.Println(res)
	return
}
