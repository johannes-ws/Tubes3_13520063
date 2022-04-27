package main

import (
	"context"
	"database/sql"
	"fmt"
)

var database *sql.DB

var (
	port     = 
	password = 
	user     = 
)

var dbContext = context.Background()

var connectionString = fmt.Sprintf("user id=%s;password=%s;port=%d", user, password, port)
