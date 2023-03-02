package main

import (
	ds "buzzstaff/dashboard"
	dbs "buzzstaff/database"
)

func main() {
	dbs.Connection()
	ds.HandleFunc()
}
