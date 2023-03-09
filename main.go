package main

import (
	d "buzzstaff-go/QualityAssessment"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "bdms_staff_admin:sfhakjfhyiqundfgs3765827635@tcp(buzzwomendatabase-new.cixgcssswxvx.ap-south-1.rds.amazonaws.com:3306)/bdms_staff?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")
	//-------------------Endpoint for adding data to QAF---------------------------------------
	http.HandleFunc("/addQualityAssessmentForm", func(w http.ResponseWriter, r *http.Request) {
		d.AddQualityAssessmentForm(w, r, db)
	})
	//-------------------Endpoint to get the total dashboard data for QAF-----------------------
	http.HandleFunc("/getDashboard", func(w http.ResponseWriter, r *http.Request) {
		d.GetDashboard(w, r, db)
	})
	//---------------------Endpoint to filter the dashboard data for QAF------------------------
	http.HandleFunc("/filterdashboard", func(w http.ResponseWriter, r *http.Request) {
		d.FilterDashboard(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
