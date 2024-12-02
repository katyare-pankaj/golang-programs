// +build linux

package main  
import (  
    "fmt"
)  

func main() {  
    fmt.Println("This code will be compiled only on Linux.")
} 

// +build windows

package main
import (  
    "fmt"
)  

func main() {  
    fmt.Println("This code will be compiled only on Windows.")
} 