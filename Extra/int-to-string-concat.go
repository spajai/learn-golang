package main

import (
	"fmt"
)

func main() {
	num := 1
	email := "bob@example.com"
	concatenated := fmt.Sprintf("%d%s", num, email)

	fmt.Println(concatenated)
}
package main

import (
	"fmt"
)

func main() {
	num := 1
	email := "bob@example.com"
	concatenated := fmt.Sprint(num, email)

	fmt.Println(concatenated)
}


package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 1
	email := "bob@example.com"
	concatenated := strconv.Itoa(num) + email

	fmt.Println(concatenated)
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 1
	email := "bob@example.com"
	concatenated := strings.Join([]string{strconv.Itoa(num), email}, "")

	fmt.Println(concatenated)
}