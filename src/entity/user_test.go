package entity

import (
	"fmt"
	"testing"
)
type test struct {
	data User
	answer string
}

func TestUserMustBeReturnAnErrorWhenPropertieIsNotDefined(t *testing.T) {
	tests := []test{
		{
			data: 
				User{
					Name: "",
					LastName: "UserLastName",
					Age: 20,
				},
				answer: "Name must be defined",
		},
		{
			data: User{
				Name: "UserTestName",
				LastName: "",
				Age: 20,
		},
			answer: "Last Name must be defined",
		},
		{
			data: User{
				Name: "UserTestName",
				LastName: "UserLastName",
				Age: 0,
			},
			answer: "Age must be greater than zero and less then 120",
		},
	} 
	for _, v := range tests{
		fmt.Println(v.data)
		isDefined := v.data.checkPropertiesMustBeDefined()
		if isDefined == nil {
			t.Error("Expected: ", v.answer, "Got: ", isDefined)
		}
	}
}


// func TestUserNameMustBeDefined(t *testing.T) {
// 	u := User{
// 		Name: "",
// 		LastName: "UserLastName",
// 		Age: 20,
// 	}
// 	expect := "Name must be defined"
// 	isDefined := u.checkPropertiesMustBeDefined()
// 	if isDefined == nil {
// 		t.Error("Expected: ", expect, "Got: ", isDefined)
// 	}
// }

// func TestUserLastNameMustBeDefined(t *testing.T) {
// 	u := User{
// 		Name: "UserTestName",
// 		LastName: "",
// 		Age: 20,
// 	}
// 	expect := "Last Name must be defined"
// 	isDefined := u.checkPropertiesMustBeDefined()
// 	if isDefined == nil {
// 		t.Error("Expected: ", expect, "Got: ", isDefined)
// 	}
// }

// func TestUserAgeMustBeDefined(t *testing.T) {
// 	u := User{
// 		Name: "UserTestName",
// 		LastName: "UserTestLastName",
// 		Age: 0,
// 	}
// 	expect := "Age must be greater than zero and less then 120"
// 	isDefined := u.checkPropertiesMustBeDefined()
// 	if isDefined == nil {
// 		t.Error("Expected: ", expect, "Got: ", isDefined)
// 	}
// }