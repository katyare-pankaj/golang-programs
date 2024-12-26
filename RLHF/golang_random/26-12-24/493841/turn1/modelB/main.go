package main  
import (  
    "encoding/csv"  
    "encoding/json"  
    "fmt"  
    "io"  
    "os"  
    "strconv"  
)  

// Define a struct to represent the data structure
type User struct {  
    ID   int    `json:"id"`  
    Name string `json:"name"`  
    Age  int    `json:"age"`  
}  

// Function to read data from a CSV file and validate it
func readDataFromCSV(fileName string) ([]User, error) {  
    file, err := os.Open(fileName)  
    if err != nil {  
        return nil, err  
    }  
    defer file.Close()  

    reader := csv.NewReader(file)  
    users := []User{}  

    for {  
        record, err := reader.Read()  
        if err == io.EOF {  
            break  
        }  
        if err != nil {  
            return nil, err  
        }  

        // Validate the CSV data
        if len(record) != 3 {  
            return nil, fmt.Errorf("Invalid CSV record: %v", record)  
        }  

        id, err := strconv.Atoi(record[0])  
        if err != nil {  
            return nil, fmt.Errorf("Invalid ID: %v", record[0])  
        }  

        age, err := strconv.Atoi(record[2])  
        if err != nil {  
            return nil, fmt.Errorf("Invalid Age: %v", record[2])  
        }  

        user := User{  
            ID:   id,  
            Name: record[1],  
            Age:  age,  
        }  
        users = append(users, user)  
    }  
    return users, nil  
}  

// Function to convert data from CSV to JSON format and validate it
func convertCSVtoJSON(users []User) ([]byte, error) {  
  // Validate the user data before converting to JSON
    for _, user := range users {  
        if user.ID <= 0 {  
            return nil, fmt.Errorf("Invalid user ID: %d", user.ID)  
        }  
        if user.Name == "" {  
            return nil, fmt.Errorf("Invalid user name: %s", user.Name)  
        }  
        if user.Age <= 0 {  
            return nil, fmt.Errorf("Invalid user age: %d", user.Age)  
        }  
    }  

    jsonData, err := json.Marshal(users)  
    if err != nil {  
        return nil, err  
    }  
    return jsonData, nil  
}  

// Function to write data to a JSON file
func writeDataToJSON(fileName string, jsonData []byte) error {  
    file, err := os.Create(fileName)  
    if err != nil {  
        return err  
    }  
    defer file.Close()  

    _, err = file.Write(jsonData)  
    if err != nil {  
        return err  
    }  
    return nil  
}  

func main() {  
    // Read data from CSV file  
    csvFileName := "users.csv"  
    users, err := readDataFromCSV(csvFileName)  
    if err != nil {  
        fmt.Println("Error reading CSV file:", err)  
        return  
    }  

    // Convert data from CSV to JSON format
    jsonData, err := convertCSVtoJSON(users)  
    if err != nil {  
        fmt.Println("Error converting CSV to JSON:", err)  
        return  
    }  

    // Write data to JSON file  
    jsonFileName := "users.json"  
    err = writeDataToJSON(jsonFileName, jsonData)