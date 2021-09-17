// package main
 
// import (
// 	"fmt"
// )
 
// func main() {
 
// 	var name string
// 	fmt.Println("Как тебя зовут?")
// 	fmt.Scanf("%s\n", &name)
 
// 	var age int
// 	fmt.Println("Сколько тебе лет?")
// 	fmt.Scanf("%d\n", &age)
 
// 	fmt.Printf("Привет, %s, твой возраст - %d\n", name, age)
 
// }

package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func main() {
	// Scanner может просканировать построчный ввод
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Эхо: %s\n", txt)
	}
}