package main

import (
	"fmt"
	"library/api"
)

func main() {
	response, err := api.Process("Съешь ещё этих мягких французских булок, да выпей же чаю.")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(response.Result)
}
