package mainTask2


import (
"fmt"
"github.com/nurtuganulyn/libTask2"
)

func main() {
	fmt.Println( libTask2.ChangeRegistCharacter('Z'))
	fmt.Println( libTask2.ChangeRegistCharacter('z'))

	r1, r2, err := libTask2.QuadraticRoot(1, -5, -9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1, r2)
	}
	fmt.Println(libTask2.GenerateUUID())
}
