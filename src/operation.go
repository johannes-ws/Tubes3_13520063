package main

import (
	"fmt"
)

func InsertPenyakit(nama, sequence string) error {
	var err error

	err = database.PingContext(dbContext)
	if err != nil {
		fmt.Printf("Error checking db connection: %v", err)
	}

	queryStatement := fmt.Sprintf("INSERT INTO penyakit (nama, sequence) VALUES ( '%v', '%v');", nama, sequence)
	query, err := database.Prepare(queryStatement)
	if err != nil {
		fmt.Printf("Query err: %v", err)
	}

	_, queryErr := query.QueryContext(dbContext)

	if queryErr != nil {
		fmt.Printf("Query err: %v", queryErr)
	}

	return nil
}
