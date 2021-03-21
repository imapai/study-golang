package main

import (
	"errors"
	"fmt"
)

func f1(ary int) (int, error) {
	if ary == 42 {
		return -1, errors.New("can't work with 42")
	}
	return ary + 1, nil
}

type aryError struct {
	code    int
	message string
}

func (e *aryError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}
func f2(ary int) (int, error) {
	if ary == 42 {
		return -1, &aryError{ary, "can't work with 42"}
	}
	return ary + 2, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)
		} else {
			fmt.Println("f1 work", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 work", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*aryError); ok {
		fmt.Println(ae.code)
		fmt.Println(ae.message)
	}
}
