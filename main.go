// package main

// import (
// 	"database/sql"
// 	_ "github.com/go-sql-driver/mysql"
// 	"log"
// 	// "github.com/joho/godotenv"
// 	"fmt"
// )

// func main() {
// 	db,err := sql.Open("mysql","****:*************************@/go_mysql_test")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	var (
// 		id int
// 		name string
// 	)

// 	rows,err := db.Query("SELECT id, name FROM users WHERE id = ?",1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err := rows.Scan(&id,&name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Fatal(err)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(name)
// }

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql","****:*************************@/go_mysql_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	id := 1
	var name string 
	err = db.QueryRow("SELECT name FROM users WHERE id = ?",id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}