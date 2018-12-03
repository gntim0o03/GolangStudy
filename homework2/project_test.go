package main

import (
	"fmt"
	"testing"
)

func TestGetMachineGame(t *testing.T) {
	ids := []string{"1", "2", "3"}

	for _, id := range ids {
		result := GetMachineGame(id)
		fmt.Println("ID :", id, "結果 :", result)

		if result == 0 {
			t.Error("Error")
		}
		if result == 1 {
			t.Log("Success")
		}
	}

}
