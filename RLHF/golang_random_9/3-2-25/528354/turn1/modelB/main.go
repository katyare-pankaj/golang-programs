package main  
import (  
    "crypto/hmac"  
    "crypto/sha256"  
    "encoding/base64"  
    "encoding/json"  
    "fmt"  
    "log"  
    "net/http"  
    "time"  
  
    "github.com/gorilla/mux"  
    "github.com/jinzhu/gorm"  
    _ "github.com/jinzhu/gorm/dialects/mysql"  
    "golang.org/x/crypto/bcrypt"  
)  
// User struct  
type User struct {  
    ID        int       `json:"id"`  
    Username  string    `json:"username"`  
    Password  string    `json:"password"`  
    Email     string    `json:"email"`  
    Role      string    `json:"role"`  
    CreatedAt time.Time `json:"createdAt"`  
    UpdatedAt time.Time `json:"updatedAt"`  
}  
// Activity struct  
type Activity struct {  
    ID        int       `json:"id"`  
    UserID    int       `json:"userID"`  
    Activity  string    `json:"activity"`  
    Timestamp time.Time `json:"timestamp"`  
}  
// JWT struct  
type JWT struct {  
    Token string `json:"token"`  
}  
var db *gorm.DB  
var jwtKey = []byte("secret-key") // Replace with your own secret key  
func main() {  
    // Initialize database connection  
    var err error  
    db, err = gorm.Open("mysql", "username:password@/dbname?charset=utf8&parseTime=True&loc=Local")  
    if err != nil {  
        panic("failed to connect database")  
    }  
    defer db.Close()  
    // Migrate models  
    db.AutoMigrate(&User{}, &Activity{})  
    // Create router  
    router := mux.NewRouter()  
    // Register routes  
    router.HandleFunc("/signup", SignupHandler).Methods("POST")  
    router.HandleFunc("/login", LoginHandler).Methods("POST")  
    router.HandleFunc("/files/rename", RenameFileHandler).Methods("POST")  
    router.HandleFunc("/files/move", MoveFileHandler).Methods("POST")  
    router.HandleFunc("/activity", ActivityLogHandler).Methods("GET")  
    // Start server  
    fmt.Println("Server is running on port 8000")  
    log.Fatal(http.ListenAndServe(":8000", router))  
}  
// SignupHandler handles user signup  
func SignupHandler(w http.ResponseWriter, r *http.Request) {  
    var user User  
    json.NewDecoder(r.Body).Decode(&user)  
    // Hash the password  
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)  
    if err != nil {  
        http.Error(w, err.Error(), http.StatusInternalServerError)  
        return  
    }  
    user.Password = string(hashedPassword)  
    // Create the user  
    db.Create(&user)  
    json.NewEncoder(w).Encode(user)  
}  
// LoginHandler handles user login  
func LoginHandler(w http.ResponseWriter, r *http.Request) {  
    var user User  
    json.NewDecoder(r.Body).Decode(&user)  
    var existingUser User  
    db.Where("username = ?", user.Username).First(&existingUser)  
    if existingUser.ID == 0 {  
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)  