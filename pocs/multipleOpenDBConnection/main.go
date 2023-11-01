package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println("This is a POC program to check if `n` PostgreSQL DB connections can stay open")
	var n int
	fmt.Print("Enter the number of connections to create: ")
	fmt.Scan(&n)

	connections := make([]*pgx.Conn, n)

	urlExample := "postgres://pond_application@localhost:5432/pond"

	for i := 0; i < n; i++ {
		conn, err := pgx.Connect(context.Background(), urlExample)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		connections = append(connections, conn)
	}

	fmt.Printf("Created %d PostgreSQL connections.\n", n)

	// Keep the connections open as long as needed
	// You can use these connections for database operations
	// Don't forget to close the connections when done with them

	// To close all the connections:
	// for _, conn := range connections {
	// 	conn.Close(context.Background())
	// }
}
