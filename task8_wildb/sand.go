// package main

// import (
// 	"fmt"
// 	"log"
// 	"os/exec"
// )

// func main() {
	
// 	path, err := exec.LookPath("sleep")
// 	if err != nil {
// 		log.Fatal("installing fortune is in your future")
// 	}
// 	fmt.Printf("fortune is available at %s\n", path)
// }

package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    out, err := exec.Command("pwd").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(out))
}