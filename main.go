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
	//---------------------Endpoint to filter the dashboard data(Date Range) for QAF------------------------
	http.HandleFunc("/filterDate", func(w http.ResponseWriter, r *http.Request) {
		d.FilterDate(w, r, db)
	})
	//---------------------Endpoint to filter the dashboard data (Employee Id) for QAF------------------------
	http.HandleFunc("/filterEmpId", func(w http.ResponseWriter, r *http.Request) {
		d.FilterEmpId(w, r, db)
	})
	//-------------------------Multifilter for dashboard----------------------------
	http.HandleFunc("/filterDashboard", func(w http.ResponseWriter, r *http.Request) {
		d.FilterDashboard(w, r, db)
	})
	//---------------------Endpoint to filter the dashboard data for QAF------------------------
	http.HandleFunc("/editQualityAssessmentForm", func(w http.ResponseWriter, r *http.Request) {
		d.EditQualityAssessmentForm(w, r, db)
	})
	//---------------------List Quality Assessment Form-----------------------------------------
	http.HandleFunc("/listQualityAssessmentForm", func(w http.ResponseWriter, r *http.Request) {
		d.ListQualityAssessmentForm(w, r, db)
	})
	//---------------------List Quality Assessment Form-----------------------------------------
	http.HandleFunc("/getUserData", func(w http.ResponseWriter, r *http.Request) {
		d.GetUserData(w, r, db)
	})
	//---------------------List Quality Assessment Form-----------------------------------------
	http.HandleFunc("/getEmpData", func(w http.ResponseWriter, r *http.Request) {
		d.GetEmpData(w, r, db)
	})
	//---------------------Endpoint to filter the dashboard data(Taluk) for QAF-----------------------------------------
	http.HandleFunc("/filterTaluk", func(w http.ResponseWriter, r *http.Request) {
		d.FilterTaluk(w, r, db)
	})
	//---------------------Endpoint to filter the dashboard data(District) for QAF-----------------------------------------
	http.HandleFunc("/filterDistrict", func(w http.ResponseWriter, r *http.Request) {
		d.FilterDistrict(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
