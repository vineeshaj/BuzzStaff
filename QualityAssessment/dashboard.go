package QualityAssessment

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Define a struct to hold data from the table
type QualityAssessmentForm struct {
	Id                                                               int    `json:"Id"`
	Emp_id                                                           string `json:"Emp_id"`
	Email                                                            string `json:"Email"`
	Entry_date                                                       string `json:"Entry_date"`
	Program_assessment                                               int    `json:"Program_assessment"`
	Name_of_the_district                                             string `json:"Name_of_the_district"`
	Name_of_the_taluk                                                string `json:"Name_of_the_taluk"`
	Name_of_the_village_and_the_venue_of_meeting_or_training         string `json:"Name_of_the_village_and_the_venue_of_meeting_or_training"`
	Day1_or_day2                                                     string `json:"Day1_or_day2"`
	Name_of_the_trainer_being_evaluated                              string `json:"Name_of_the_trainer_being_evaluated"`
	Check_which_ones_the_trainer_did_not_do                          string `json:"Check_which_ones_the_trainer_did_not_do"`
	How_many_women_attended_the_training_session                     string `json:"How_many_women_attended_the_training_session"`
	Were_the_women_interactive                                       string `json:"Were_the_women_interactive"`
	Did_any_women_leave_the_session_during_or_after_the_first_module string `json:"Did_any_women_leave_the_session_during_or_after_the_first_module"`
	If_so_how_many                                                   string `json:"If_so_how_many"`
	Did_this_module_take_20_minutes_as_allotted                      string `json:"Did_this_module_take_20_minutes_as_allotted"`
	Did_any_new_women_attend_the_training_session_during_this_module string `json:"Did_any_new_women_attend_the_training_session_during_this_module"`
	During_the_debrief_the_trainer_did                               string `json:"During_the_debrief_the_trainer_did"`
	During_the_debriefs_for_role_plays_the_trainer_did_not_ask       string `json:"During_the_debriefs_for_role_plays_the_trainer_did_not_ask"`
	Did_the_trainer_leave_the_women_to_read_role_play_card           string `json:"Did_the_trainer_leave_the_women_to_read_role_play_card"`
	Did_the_groups_engage_and_interact_among_themselves_well         string `json:"Did_the_groups_engage_and_interact_among_themselves_well"`
	Were_the_participants_responsive_during_the_debriefing           string `json:"Were_the_participants_responsive_during_the_debriefing"`
	Did_this_module_take_30_minutes_as_allotted                      string `json:"Did_this_module_take_30_minutes_as_allotted"`
	How_many_women_remained_by_the_end_of_this_training_session      string `json:"How_many_women_remained_by_the_end_of_this_training_session"`
	How_many_are_likely_to_come_back                                 string `json:"How_many_are_likely_to_come_back"`
	Was_the_recap_done                                               string `json:"Was_the_recap_done"`
	Did_the_recap_take_15_minutes_as_allotted                        string `json:"Did_the_recap_take_15_minutes_as_allotted"`
	During_the_debrief_did_the_trainer_did_not_do_the_following      string `json:"During_the_debrief_did_the_trainer_did_not_do_the_following"`
	Check_which_instructions_the_trainer_did_not_do                  string `json:"Check_which_instructions_the_trainer_did_not_do"`
	Repeat_the_activity_with_the_second_volunteer                    string `json:"Repeat_the_activity_with_the_second_volunteer"`
	During_the_debrief_did_the_trainer_not_ask                       string `json:"During_the_debrief_did_the_trainer_not_ask"`
	The_trainer_did_not_ask                                          string `json:"The_trainer_did_not_ask"`
	What_did_the_trainer_not_do                                      string `json:"What_did_the_trainer_not_do"`
	No_of_participants_at_the_start_of_the_session                   string `json:"No_of_participants_at_the_start_of_the_session"`
	Assessment_of                                                    string `json:"Assessment_of"`
	The_gf_competently_carried_out_the_following_functions           string `json:"The_gf_competently_carried_out_the_following_functions"`
	The_gf_carried_out_the_functions_before_the_training_started     string `json:"The_gf_carried_out_the_functions_before_the_training_started"`
	How_many_stories_of_success_or_change_emerged_from_the_recap     string `json:"How_many_stories_of_success_or_change_emerged_from_the_recap"`
	Mention_name_of_the_gelathis_with_success_stories                string `json:"Mention_name_of_the_gelathis_with_success_stories"`
	Check_which_ones_the_gelathi_facilitator_did_not_do              string `json:"Check_which_ones_the_gelathi_facilitator_did_not_do"`
	Number_of_enrolled_gelathis_in_the_circle                        string `json:"Number_of_enrolled_gelathis_in_the_circle"`
	No_of_attended_gelathis                                          string `json:"No_of_attended_gelathis"`
	Level_of_participation                                           string `json:"Level_of_participation"`
	The_gf_covered_the_following_things_in_the_training              string `json:"The_gf_covered_the_following_things_in_the_training"`
	Rate_the_gelathi_facilitator                                     string `json:"Rate_the_gelathi_facilitator"`
	What_worked_in_the_training                                      string `json:"What_worked_in_the_training"`
	What_can_be_better_next_time                                     string `json:"What_can_be_better_next_time"`
	Any_further_training_required_by_the_gf_of_modules_delivered     string `json:"Any_further_training_required_by_the_gf_of_modules_delivered"`
	Anything_in_the_training_or_gf_needs_to_be_worked_on_priority    string `json:"Anything_in_the_training_or_gf_needs_to_be_worked_on_priority"`
	Details_of_success_stories_to_be_collected_from_gelathis_by_gf   string `json:"Details_of_success_stories_to_be_collected_from_gelathis_by_gf"`
	Deadline_to_collect_the_stories                                  string `json:"Deadline_to_collect_the_stories"`
	End_time_of_the_training                                         string `json:"End_time_of_the_training"`
	No_of_participants_at_end_of_the_session                         string `json:"No_of_participants_at_end_of_the_session"`
	Any_other_comments_about_the_gelathi_facilitator                 string `json:"Any_other_comments_about_the_gelathi_facilitator"`
	Check_which_ones_the_gelathi_did_not_do                          string `json:"Check_which_ones_the_gelathi_did_not_do"`
	During_the_debrief_did_the_gelathi                               string `json:"During_the_debrief_did_the_gelathi"`
	During_the_debriefs_for_role_plays_the_gelathi_did_not_ask       string `json:"During_the_debriefs_for_role_plays_the_gelathi_did_not_ask"`
	Check_which_instructions_the_gelathi_did_not_do                  string `json:"Check_which_instructions_the_gelathi_did_not_do"`
	Did_the_debrief_done_by_gelathi                                  string `json:"Did_the_debrief_done_by_gelathi"`
	The_gelathi_did_not_ask                                          string `json:"The_gelathi_did_not_ask"`
	During_the_debrief_did_the_gelathi_not_ask                       string `json:"During_the_debrief_did_the_gelathi_not_ask"`
}

var DB *sql.DB

// ------------------------------------Adding Data into Quality Assessment Form-------------------------------------------------------------
func AddQualityAssessmentForm(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	// Decode request body into QualityAssessmentForm struct
	var q QualityAssessmentForm
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare SQL statement
	stmt, err := DB.Prepare(`INSERT INTO QualityAssessmentForm ( 
		Emp_id, 
		Email, 
		Entry_date, 
		Program_assessment,
		Name_of_the_district, 
		Name_of_the_taluk, 
		Name_of_the_village_and_the_venue_of_meeting_or_training, 
		Day1_or_day2,
		Name_of_the_trainer_being_evaluated,
		Check_which_ones_the_trainer_did_not_do,
		How_many_women_attended_the_training_session,
		Were_the_women_interactive,
		Did_any_women_leave_the_session_during_or_after_the_first_module,
		If_so_how_many,
		Did_this_module_take_20_minutes_as_allotted,
		Did_any_new_women_attend_the_training_session_during_this_module,
		During_the_debrief_the_trainer_did,
		During_the_debriefs_for_role_plays_the_trainer_did_not_ask,
		Did_the_trainer_leave_the_women_to_read_role_play_card,
		Did_the_groups_engage_and_interact_among_themselves_well,
		Were_the_participants_responsive_during_the_debriefing,
		Did_this_module_take_30_minutes_as_allotted,
		How_many_women_remained_by_the_end_of_this_training_session,
		How_many_are_likely_to_come_back,
		Was_the_recap_done,
		Did_the_recap_take_15_minutes_as_allotted,
		During_the_debrief_did_the_trainer_did_not_do_the_following,
		Check_which_instructions_the_trainer_did_not_do,
		Repeat_the_activity_with_the_second_volunteer,
		During_the_debrief_did_the_trainer_not_ask,
		The_trainer_did_not_ask,
		What_did_the_trainer_not_do,
		No_of_participants_at_the_start_of_the_session,
		Assessment_of,
		The_gf_competently_carried_out_the_following_functions,
		The_gf_carried_out_the_functions_before_the_training_started,
		How_many_stories_of_success_or_change_emerged_from_the_recap,
		Mention_name_of_the_gelathis_with_success_stories,
		Check_which_ones_the_gelathi_facilitator_did_not_do,
		Number_of_enrolled_gelathis_in_the_circle,
		No_of_attended_gelathis,
		Level_of_participation,
		The_gf_covered_the_following_things_in_the_training,
		Rate_the_gelathi_facilitator,
		What_worked_in_the_training,
		What_can_be_better_next_time,
		Any_further_training_required_by_the_gf_of_modules_delivered,
		Anything_in_the_training_or_gf_needs_to_be_worked_on_priority,
		Details_of_success_stories_to_be_collected_from_gelathis_by_gf,
		Deadline_to_collect_the_stories,
		End_time_of_the_training,
		No_of_participants_at_end_of_the_session,
		Any_other_comments_about_the_gelathi_facilitator,
		Check_which_ones_the_gelathi_did_not_do,
		During_the_debrief_did_the_gelathi,
		During_the_debriefs_for_role_plays_the_gelathi_did_not_ask,
		Check_which_instructions_the_gelathi_did_not_do,
		Did_the_debrief_done_by_gelathi,
		The_gelathi_did_not_ask,
		During_the_debrief_did_the_gelathi_not_ask
	) VALUES (?,?,LOCALTIMESTAMP(0),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute SQL statement with data from QualityAssessmentForm struct
	_, err = stmt.Exec(
		q.Emp_id,
		q.Email,
		q.Program_assessment,
		q.Name_of_the_district,
		q.Name_of_the_taluk,
		q.Name_of_the_village_and_the_venue_of_meeting_or_training,
		q.Day1_or_day2,
		q.Name_of_the_trainer_being_evaluated,
		q.Check_which_ones_the_trainer_did_not_do,
		q.How_many_women_attended_the_training_session,
		q.Were_the_women_interactive,
		q.Did_any_women_leave_the_session_during_or_after_the_first_module,
		q.If_so_how_many,
		q.Did_this_module_take_20_minutes_as_allotted,
		q.Did_any_new_women_attend_the_training_session_during_this_module,
		q.During_the_debrief_the_trainer_did,
		q.During_the_debriefs_for_role_plays_the_trainer_did_not_ask,
		q.Did_the_trainer_leave_the_women_to_read_role_play_card,
		q.Did_the_groups_engage_and_interact_among_themselves_well,
		q.Were_the_participants_responsive_during_the_debriefing,
		q.Did_this_module_take_30_minutes_as_allotted,
		q.How_many_women_remained_by_the_end_of_this_training_session,
		q.How_many_are_likely_to_come_back,
		q.Was_the_recap_done,
		q.Did_the_recap_take_15_minutes_as_allotted,
		q.During_the_debrief_did_the_trainer_did_not_do_the_following,
		q.Check_which_instructions_the_trainer_did_not_do,
		q.Repeat_the_activity_with_the_second_volunteer,
		q.During_the_debrief_did_the_trainer_not_ask,
		q.The_trainer_did_not_ask,
		q.What_did_the_trainer_not_do,
		q.No_of_participants_at_the_start_of_the_session,
		q.Assessment_of,
		q.The_gf_competently_carried_out_the_following_functions,
		q.The_gf_carried_out_the_functions_before_the_training_started,
		q.How_many_stories_of_success_or_change_emerged_from_the_recap,
		q.Mention_name_of_the_gelathis_with_success_stories,
		q.Check_which_ones_the_gelathi_facilitator_did_not_do,
		q.Number_of_enrolled_gelathis_in_the_circle,
		q.No_of_attended_gelathis,
		q.Level_of_participation,
		q.The_gf_covered_the_following_things_in_the_training,
		q.Rate_the_gelathi_facilitator,
		q.What_worked_in_the_training,
		q.What_can_be_better_next_time,
		q.Any_further_training_required_by_the_gf_of_modules_delivered,
		q.Anything_in_the_training_or_gf_needs_to_be_worked_on_priority,
		q.Details_of_success_stories_to_be_collected_from_gelathis_by_gf,
		q.Deadline_to_collect_the_stories,
		q.End_time_of_the_training,
		q.No_of_participants_at_end_of_the_session,
		q.Any_other_comments_about_the_gelathi_facilitator,
		q.Check_which_ones_the_gelathi_did_not_do,
		q.During_the_debrief_did_the_gelathi,
		q.During_the_debriefs_for_role_plays_the_gelathi_did_not_ask,
		q.Check_which_instructions_the_gelathi_did_not_do,
		q.Did_the_debrief_done_by_gelathi,
		q.The_gelathi_did_not_ask,
		q.During_the_debrief_did_the_gelathi_not_ask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data inserted successfully"))
}

// -------------------------------------Total Dashboard for Quality Assessment Form-------------------------------------------------------
func GetDashboard(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	rows, err := DB.Query(`SELECT 
    (SELECT COUNT(*) FROM QualityAssessmentForm WHERE Program_assessment=1) as SStraining,
    (SELECT COUNT(*) FROM QualityAssessmentForm WHERE Program_assessment=2) as GelathiProgram,
    (SELECT COUNT(*) FROM QualityAssessmentForm WHERE Program_assessment=3) as SSbyGelathi;`)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
	}
	var SStraining, GelathiProgram, SSbyGelathi int
	for rows.Next() {
		err := rows.Scan(&SStraining, &GelathiProgram, &SSbyGelathi)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(w).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
		}
	}
	dashboard := map[string]int{"SStraining": SStraining, "GelathiProgram": GelathiProgram, "SSbyGelathi": SSbyGelathi}
	ww, _ := json.Marshal(dashboard)
	w.Write(ww)
}

// --------------------------------------Filter Dashboard for Quality Assessment Form---------------------------------------------------

func FilterDashboard(w http.ResponseWriter, r *http.Request, DB *sql.DB) {
	var qua QualityAssessmentForm

	err := json.NewDecoder(r.Body).Decode(&qua)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": "Invalid request body"})
		return
	}

	if qua.Emp_id == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": "Emp_id is required"})
		return
	}

	rows, err := DB.Query(`SELECT 
		(SELECT COUNT(*) FROM QualityAssessmentForm WHERE Program_assessment=1 AND Emp_id=?) as SStraining,
		(SELECT COUNT(*) FROM QualityAssessmentForm WHERE Program_assessment=2 AND Emp_id=?) as GelathiProgram,
		(SELECT COUNT(*) FROM QualityAssessmentForm WHERE Program_assessment=3 AND Emp_id=?) as SSbyGelathi`, qua.Emp_id, qua.Emp_id, qua.Emp_id)
	if err != nil {
		log.Fatal(err)
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
		return
	}
	defer rows.Close()

	var SStraining, GelathiProgram, SSbyGelathi int
	for rows.Next() {
		err := rows.Scan(&SStraining, &GelathiProgram, &SSbyGelathi)
		if err != nil {
			log.Fatal(err)
			json.NewEncoder(w).Encode(map[string]interface{}{"status": "400 Bad Request", "Message": err})
			return
		}
	}
	dashboard := map[string]int{"SStraining": SStraining, "GelathiProgram": GelathiProgram, "SSbyGelathi": SSbyGelathi}
	ww, _ := json.Marshal(dashboard)
	w.Write(ww)
}
