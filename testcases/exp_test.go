package testcases

import (
	"calculator/app"
	"fmt"
	"testing"
)

func TestExpression(t *testing.T) {
	InitBasicOperationsResultMap()
	InitCombinedAdditionSubtractionOperationsMap()

	for indx, item := range BasicOperations {
		result, err := app.Valuate(item)
		if err != nil {
			t.Error("err: ", err)
		}

		if result != BasicOperationsResult[indx] {
			t.Errorf("failed %s : %s", item, fmt.Sprint(result))
		}
	}

	for indx, item := range CombinedAdditionSubtractionOperations {
		result, err := app.Valuate(item)
		if err != nil {
			t.Error("err: ", err)
		}

		if result != CombinedAdditionSubtraction[indx] {
			t.Errorf("failed %s : %s", item, fmt.Sprint(result))
		}
	}

}
