//Package dog calculates dog years with reference to human years
package dog

import (
	"fmt"

	"github.com/sirajummprince/go-exercises/exerciseseventythree/dog"
)

//Years converts human years into dog years
func Years(humanYears int) int {
	return humanYears * 7
}

//ExampleDog displays an example for package dog
func ExampleDog() {
	dogYears := dog.Years(10)
	fmt.Println(dogYears)
	//Output:
	//70
}
