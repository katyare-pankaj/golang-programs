package dbOperation

import (
        "database/sql"
        "time"
        "fmt"
)
// This Function transforms our error or NULL handling into things using something like a "withClose"
type user struct {
        ID int       `json:"id"`
        Username string `json:"username"`
        Email   string `json:"email"`
}


// example query
query := `SELECT id, username, email FROM users`
        rows, err := db.Query(query)
        if err != nil {
                fmt.Println(err)
                return nil, err
        }
        defer rows.Close()

   // Our data might have lots of details in many rows-encapsulating retrieval with unique
        var users []user
        for rows.Next() {
                var u user
                if err := rows.Scan(&u.ID, &u.Username, &u.Email); err != nil {
                        fmt.Println(err)
                        return nil, err
                }
                users = append(users, u)
        }