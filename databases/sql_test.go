package databases

import (
	"context"
	"fmt"
	"testing"
	"time"
)

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
