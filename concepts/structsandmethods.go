package concepts

import (
	"fmt"
	"errors"
)

type user struct {
	name, role, email string
	age int
}

func (u user) salary() (int, error) {
	if u.role == "" {
		return 0, errors.New("I cannot handle empty roles")
	}
	switch u.role {
	case "Developer":
		return 100, nil
	case "Architect":
		return 200, nil
	}
	return 0, errors.New(fmt.Sprintf("I am not able to handle '%s' role", u.role))
}

func (u *user) updateEmail(email string) {
	u.email = email
}

func StructExample() {
	user := user{
		name:  "Gagan",
		role:  "Developer",
		age:   27,
	}
	fmt.Println(user)
}

func StructWithMethodsExample() {
	gagan := user{
		name:  "Gagan",
		role:  "Developer",
		age:   27,
	}
	futureGagan := user{
		name:  "Gagan",
		role:  "Architect",
		age:   28,
	}
	userOne := user{
		name:  "Gagan",
		role:  "Developer",
		email: "usser@rmail.com",
		age:   27,
	}
	fmt.Println(userOne)
	fmt.Println(gagan.salary())
	fmt.Println(futureGagan.salary())
	userOne.updateEmail("user@gmail.com")
	fmt.Println(userOne)
	userTwo := user{
		name:  "Gagan",
		role:  "Designer",
		email: "usser@rmail.com",
		age:   27,
	}
	salary, err := userTwo.salary()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salary: ", salary)
	}
	userThree := user{
		name:  "Gagan",
		role:  "",
		email: "usser@rmail.com",
		age:   27,
	}
	salary, err = userThree.salary()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salary: ", salary)
	}
}
