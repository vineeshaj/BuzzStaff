package dashboard

import (
	dbs "buzzstaff/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type QualityAssessmentForm struct {
	Id                                                               int       `json:"Id"`
	Emp_id                                                           int       `json:"Emp_id"`
	Email                                                            string    `json:"Email"`
	Entry_date                                                       time.Time `json:"Entry_date"`
	Program_assessment                                               string    `json:"Program_assessment"`
	Name_of_the_district                                             string    `json:"Name_of_the_district"`
	Name_of_the_taluk                                                string    `json:"Name_of_the_taluk"`
	Name_of_the_village_and_the_venue_of_meeting_or_training         string    `json:"Name_of_the_village_and_the_venue_of_meeting_or_training"`
	Day1_or_day2                                                     string    `json:"Day1_or_day2"`
	Name_of_the_trainer_being_evaluated                              string    `json:"Name_of_the_trainer_being_evaluated"`
	Check_which_ones_the_trainer_did_not_do                          string    `json:"Check_which_ones_the_trainer_did_not_do"`
	How_many_women_attended_the_training_session                     string    `json:"How_many_women_attended_the_training_session"`
	Were_the_women_interactive                                       string    `json:"Were_the_women_interactive"`
	Did_any_women_leave_the_session_during_or_after_the_first_module string    `json:"Did_any_women_leave_the_session_during_or_after_the_first_module"`
	If_so_how_many                                                   string    `json:"If_so_how_many"`
	Did_this_module_take_20_minutes_as_allotted                      string    `json:"Did_this_module_take_20_minutes_as_allotted"`
	Did_any_new_women_attend_the_training_session_during_this_module string    `json:"Did_any_new_women_attend_the_training_session_during_this_module"`
	During_the_debrief_the_trainer_did                               string    `json:"During_the_debrief_the_trainer_did"`
	During_the_debriefs_for_role_plays_the_trainer_did_not_ask       string    `json:"During_the_debriefs_for_role_plays_the_trainer_did_not_ask"`
	Did_the_trainer_leave_the_women_to_read_role_play_card           string    `json:"Did_the_trainer_leave_the_women_to_read_role_play_card"`
	Did_the_groups_engage_and_interact_among_themselves_well         string    `json:"Did_the_groups_engage_and_interact_among_themselves_well"`
	Were_the_participants_responsive_during_the_debriefing           string    `json:"Were_the_participants_responsive_during_the_debriefing"`
	Did_this_module_take_30_minutes_as_allotted                      string    `json:"Did_this_module_take_30_minutes_as_allotted"`
	How_many_women_remained_by_the_end_of_this_training_session      string    `json:"How_many_women_remained_by_the_end_of_this_training_session"`
	How_many_are_likely_to_come_back                                 string    `json:"How_many_are_likely_to_come_back"`
	Was_the_recap_done                                               string    `json:"Was_the_recap_done"`
	Did_the_recap_take_15_minutes_as_allotted                        string    `json:"Did_the_recap_take_15_minutes_as_allotted"`
	During_the_debrief_did_the_trainer_did_not_do_the_following      string    `json:"During_the_debrief_did_the_trainer_did_not_do_the_following"`
	Check_which_instructions_the_trainer_did_not_do                  string    `json:"Check_which_instructions_the_trainer_did_not_do"`
	Repeat_the_activity_with_the_second_volunteer                    string    `json:"Repeat_the_activity_with_the_second_volunteer"`
	During_the_debrief_did_the_trainer_not_ask                       string    `json:"During_the_debrief_did_the_trainer_not_ask"`
	The_trainer_did_not_ask                                          string    `json:"The_trainer_did_not_ask"`
	What_did_the_trainer_not_do                                      string    `json:"What_did_the_trainer_not_do"`
	No_of_participants_at_the_start_of_the_session                   string    `json:"No_of_participants_at_the_start_of_the_session"`
	Assessment_of                                                    string    `json:"Assessment_of"`
	The_gf_competently_carried_out_the_following_functions           string    `json:"The_gf_competently_carried_out_the_following_functions"`
	The_gf_carried_out_the_functions_before_the_training_started     string    `json:"The_gf_carried_out_the_functions_before_the_training_started"`
	How_many_stories_of_success_or_change_emerged_from_the_recap     string    `json:"How_many_stories_of_success_or_change_emerged_from_the_recap"`
	Mention_name_of_the_gelathis_with_success_stories                string    `json:"Mention_name_of_the_gelathis_with_success_stories"`
	Check_which_ones_the_gelathi_facilitator_did_not_do              string    `json:"Check_which_ones_the_gelathi_facilitator_did_not_do"`
	Number_of_enrolled_gelathis_in_the_circle                        string    `json:"Number_of_enrolled_gelathis_in_the_circle"`
	No_of_attended_gelathis                                          string    `json:"No_of_attended_gelathis"`
	Level_of_participation                                           string    `json:"Level_of_participation"`
	The_gf_covered_the_following_things_in_the_training              string    `json:"The_gf_covered_the_following_things_in_the_training"`
	Rate_the_gelathi_facilitator                                     string    `json:"Rate_the_gelathi_facilitator"`
	What_worked_in_the_training                                      string    `json:"What_worked_in_the_training"`
	What_can_be_better_next_time                                     string    `json:"What_can_be_better_next_time"`
	Any_further_training_required_by_the_gf_of_modules_delivered     string    `json:"Any_further_training_required_by_the_gf_of_modules_delivered"`
	Anything_in_the_training_or_gf_needs_to_be_worked_on_priority    string    `json:"Anything_in_the_training_or_gf_needs_to_be_worked_on_priority"`
	Details_of_success_stories_to_be_collected_from_gelathis_by_gf   string    `json:"Details_of_success_stories_to_be_collected_from_gelathis_by_gf"`
	Deadline_to_collect_the_stories                                  string    `json:"Deadline_to_collect_the_stories"`
	End_time_of_the_training                                         string    `json:"End_time_of_the_training"`
	No_of_participants_at_end_of_the_session                         string    `json:"No_of_participants_at_end_of_the_session"`
	Any_other_comments_about_the_gelathi_facilitator                 string    `json:"Any_other_comments_about_the_gelathi_facilitator"`
	Check_which_ones_the_gelathi_did_not_do                          string    `json:"Check_which_ones_the_gelathi_did_not_do"`
	During_the_debrief_did_the_gelathi                               string    `json:"During_the_debrief_did_the_gelathi"`
	During_the_debriefs_for_role_plays_the_gelathi_did_not_ask       string    `json:"During_the_debriefs_for_role_plays_the_gelathi_did_not_ask"`
	Check_which_instructions_the_gelathi_did_not_do                  string    `json:"Check_which_instructions_the_gelathi_did_not_do"`
	Did_the_debrief_done_by_gelathi                                  string    `json:"Did_the_debrief_done_by_gelathi"`
	The_gelathi_did_not_ask                                          string    `json:"The_gelathi_did_not_ask"`
	During_the_debrief_did_the_gelathi_not_ask                       string    `json:"During_the_debrief_did_the_gelathi_not_ask"`
}

var db = dbs.Connection()

func HandleFunc() {
	mux := http.NewServeMux()
	//----------------------------------Add Quality Assessment Form---------------------------------------
	mux.HandleFunc("/addQualityAssessmentForm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method != http.MethodPost {
			w.WriteHeader(405) // Return 405 Method Not Allowed.
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
			return
		}
		var qualityassessment QualityAssessmentForm
		var Id int
		Id = 0
		err := json.NewDecoder(r.Body).Decode(&qualityassessment)
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"Message": "Invalid Input Syntax", "Status Code": "400 "})
			return
		}

		err = db.QueryRow("Select Id from QualityAssessmentForm where Id=$1", qualityassessment.Id).Scan(&Id)
		Id = Id + 1
		addStatement := `INSERT INTO QualityAssessmentForm (
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
			) 
			VALUES (
				$1,$2,LOCALTIMESTAMP(0),$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59)`
		_, err = db.Exec(addStatement,
			qualityassessment.Emp_id,
			qualityassessment.Email,
			qualityassessment.Program_assessment,
			qualityassessment.Name_of_the_district,
			qualityassessment.Name_of_the_taluk,
			qualityassessment.Name_of_the_village_and_the_venue_of_meeting_or_training,
			qualityassessment.Day1_or_day2,
			qualityassessment.Name_of_the_trainer_being_evaluated,
			qualityassessment.Check_which_ones_the_trainer_did_not_do,
			qualityassessment.How_many_women_attended_the_training_session,
			qualityassessment.Were_the_women_interactive,
			qualityassessment.Did_any_women_leave_the_session_during_or_after_the_first_module,
			qualityassessment.If_so_how_many,
			qualityassessment.Did_this_module_take_20_minutes_as_allotted,
			qualityassessment.Did_any_new_women_attend_the_training_session_during_this_module,
			qualityassessment.During_the_debrief_the_trainer_did,
			qualityassessment.During_the_debriefs_for_role_plays_the_trainer_did_not_ask,
			qualityassessment.Did_the_trainer_leave_the_women_to_read_role_play_card,
			qualityassessment.Did_the_groups_engage_and_interact_among_themselves_well,
			qualityassessment.Were_the_participants_responsive_during_the_debriefing,
			qualityassessment.Did_this_module_take_30_minutes_as_allotted,
			qualityassessment.How_many_women_remained_by_the_end_of_this_training_session,
			qualityassessment.How_many_are_likely_to_come_back,
			qualityassessment.Was_the_recap_done,
			qualityassessment.Did_the_recap_take_15_minutes_as_allotted,
			qualityassessment.During_the_debrief_did_the_trainer_did_not_do_the_following,
			qualityassessment.Check_which_instructions_the_trainer_did_not_do,
			qualityassessment.Repeat_the_activity_with_the_second_volunteer,
			qualityassessment.During_the_debrief_did_the_trainer_not_ask,
			qualityassessment.The_trainer_did_not_ask,
			qualityassessment.What_did_the_trainer_not_do,
			qualityassessment.No_of_participants_at_the_start_of_the_session,
			qualityassessment.Assessment_of,
			qualityassessment.The_gf_competently_carried_out_the_following_functions,
			qualityassessment.The_gf_carried_out_the_functions_before_the_training_started,
			qualityassessment.How_many_stories_of_success_or_change_emerged_from_the_recap,
			qualityassessment.Mention_name_of_the_gelathis_with_success_stories,
			qualityassessment.Check_which_ones_the_gelathi_facilitator_did_not_do,
			qualityassessment.Number_of_enrolled_gelathis_in_the_circle,
			qualityassessment.No_of_attended_gelathis,
			qualityassessment.Level_of_participation,
			qualityassessment.The_gf_covered_the_following_things_in_the_training,
			qualityassessment.Rate_the_gelathi_facilitator,
			qualityassessment.What_worked_in_the_training,
			qualityassessment.What_can_be_better_next_time,
			qualityassessment.Any_further_training_required_by_the_gf_of_modules_delivered,
			qualityassessment.Anything_in_the_training_or_gf_needs_to_be_worked_on_priority,
			qualityassessment.Details_of_success_stories_to_be_collected_from_gelathis_by_gf,
			qualityassessment.Deadline_to_collect_the_stories,
			qualityassessment.End_time_of_the_training,
			qualityassessment.No_of_participants_at_end_of_the_session,
			qualityassessment.Any_other_comments_about_the_gelathi_facilitator,
			qualityassessment.Check_which_ones_the_gelathi_did_not_do,
			qualityassessment.During_the_debrief_did_the_gelathi,
			qualityassessment.During_the_debriefs_for_role_plays_the_gelathi_did_not_ask,
			qualityassessment.Check_which_instructions_the_gelathi_did_not_do,
			qualityassessment.Did_the_debrief_done_by_gelathi,
			qualityassessment.The_gelathi_did_not_ask,
			qualityassessment.During_the_debrief_did_the_gelathi_not_ask)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": "", "Status Code": "400", "Error": err})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"Status Code": "200 OK", "Message": "Recorded Successfully"})
	})
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8000", handler)
}
