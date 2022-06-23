package main

import "testing"

func TestProcessInputTabularMethod(t *testing.T) {

	cases := []struct {
		name        string
		num1        int
		num2        int
		res         int
		expectError bool
	}{
		{
			name:        "Case 1",
			num1:        1,
			num2:        2,
			res:         3,
			expectError: false,
		},
		{
			name:        "Case 2",
			num1:        100,
			num2:        200,
			res:         300,
			expectError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			myVariable := myStruct{}
			res, err := processInput(myVariable, tc.num1, tc.num2)
			if err != nil && !tc.expectError {
				t.Fatal(err)
			}
			if res != tc.res {
				t.Fatalf("Expected result %d got %d", tc.res, res)
			}
		})
	}
}
