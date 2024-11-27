package main  
  
import (  
    "flag"  
    "fmt"
)  
  
var username = flag.String("username", "", "Username")  
var password = flag.String("password", "", "Password")  
  
func main() {  
    flag.Parse()  

    correctUsername := "myuser"
    correctPassword := "mypassword"

    if *username == correctUsername && *password == correctPassword {
        fmt.Println("Login successful")  
    } else {  
        fmt.Println("Invalid credentials")  
    }  
}  