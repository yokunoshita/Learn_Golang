package databases

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// TestExeSql demonstrates how to execute a simple SQL insert statement using ExecContext.
func TestExeSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, nama) VALUES('fifth', 'Yukino')"

	_, err := db.ExecContext(ctx, query) // ExecContext is used to execute a query without returning any rows
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted Successfuly")
}

// TestSelectSql demonstrates how to select data from a database table and scan the results into Go variables.
func TestSelectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, nama, email FROM customer"

	rows, err := db.QueryContext(ctx, query) // QueryContext is used to execute a query that returns rows
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, nama, email string
		err := rows.Scan(&id, &nama, &email)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Nama:", nama)
		fmt.Println("Email:", email)
		fmt.Println("-----")
	}
}

// TestComplexSelect demonstrates how to select multiple data types from a database table and scan them into appropriate Go variables.
func TestComplexSelect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// query := "SELECT * FROM customer"
	query := "SELECT id, nama, email, balance, rating, birth_date, created_at, married FROM customer" // best practice to avoid using SELECT *

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, nama, email string
		var balance int32
		var rating float32
		var created_at, birthdate time.Time
		var married bool

		err := rows.Scan(&id, &nama, &email, &balance, &rating, &birthdate, &created_at, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Nama:", nama)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birthdate:", birthdate)
		fmt.Println("Created At:", created_at)
		fmt.Println("Married:", married)
		fmt.Println("-----")
	}
}

// TestSqlWithParams demonstrates how to use parameters in an SQL query to prevent SQL injection attacks.
func TestSqlWithParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "user1"
	password := "user1"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	// fmt.Println(query)

	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login success:", username)
	} else {
		fmt.Println("Failed to login: User not found")
	}

}

// TestExeWithParams demonstrates how to execute an SQL statement with parameters to insert a new user into the database.
func TestExeWithParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "user2"
	password := "user2"

	query := "INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Inserted Successfuly")
}

// TestLastInsertId demonstrates how to retrieve the ID of the last inserted record in a database table.
func TestLastInsertId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "yokubo@go.dev"
	comment := "This is a test comment 1"

	querry := "INSERT INTO comments(email, comment) VALUES (?, ?) "
	result, err := db.ExecContext(ctx, querry, email, comment)
	if err != nil {
		panic(err)
	}

	insert, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id:", insert)
}

// TestPrepareStatement demonstrates the use of prepared statements for executing the same SQL statement multiple times with different parameters.
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "Yoku" + strconv.Itoa(i) + "@go.dev"
		comment := "This is comment number " + strconv.Itoa(i)
		res, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertedId, _ := res.LastInsertId()
		fmt.Println("Success insert new comment with id:", lastInsertedId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {

		email := "Yoku" + strconv.Itoa(i) + "@go.dev"
		comment := "This is comment number " + strconv.Itoa(i)

		_, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	}
	fmt.Println("Success insert new comments")
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
