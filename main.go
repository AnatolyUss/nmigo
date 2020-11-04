package main

import (
	"runtime"
	//"time"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	PrintMemUsage()

	db, err := sql.Open("mysql", "root:asd@/staff")

	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	query := "SELECT id, employee_id FROM salary -- WHERE id < 500000"
	rows, err := db.Query(query)

	PrintMemUsage()

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columns, err := rows.Columns()

	if err != nil {
		panic(err.Error())
	}

	columnsLen := len(columns)

	// Make a slice for the values
	values := make([]sql.RawBytes, columnsLen)

	valuesLen := len(values)
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, valuesLen)

	// columnsLen: 2 valuesLen: 2
	fmt.Printf("columnsLen: %d valuesLen: %d \n", columnsLen, valuesLen)

	for i := range values {
		scanArgs[i] = &values[i]
	}

	PrintMemUsage()

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		if err = rows.Scan(scanArgs...); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string

		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			if value == "2000" {
				PrintMemUsage()
				fmt.Println(columns[i], ": ", value)
			}

			if value == "650000" {
				PrintMemUsage()
				fmt.Println(columns[i], ": ", value)
			}

			if value == "1000000" {
				PrintMemUsage()
				fmt.Println(columns[i], ": ", value)
			}

			if value == "2000000" {
				PrintMemUsage()
				fmt.Println(columns[i], ": ", value)
			}

			if value == "2844046" {
				PrintMemUsage()
				fmt.Println(columns[i], ": ", value)
			}
		}

		// fmt.Println("-----------------------------------")
	}

	PrintMemUsage()

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	///////////////////////////////////////////////////////////////
	//db, err := sql.Open("mysql", "root:7370142533@/sample_staff")
	//
	//if err != nil {
	//	panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	//}
	//
	//defer db.Close()
	//
	//stmtOut, err := db.Prepare("SELECT id, employee_id FROM salary WHERE id = ?")
	//
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//
	//defer stmtOut.Close()
	//
	//var id int // we "scan" the result in here
	//var employeeId int
	//
	//if err = stmtOut.QueryRow(13).Scan(&id, &employeeId); err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//
	//fmt.Printf("id: %d employee_id: %d", id, employeeId)
	//fmt.Println()
	//
	//if err = stmtOut.QueryRow(1).Scan(&id, &employeeId); err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//
	//fmt.Printf("id: %d employee_id: %d", id, employeeId)
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	fmt.Println()
	fmt.Println("========================================")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	fmt.Println("========================================")
	fmt.Println()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
