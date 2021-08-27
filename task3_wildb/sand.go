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
		fmt.Printf("Эхо: %s, %T\n", txt, txt)
	}
}
