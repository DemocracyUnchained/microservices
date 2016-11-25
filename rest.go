
/************************************************************************************************
 * Copyright 2016 Democracy Unchained -- http://DemocracyUnchained.org
 *
 * Licensed under the GNU General Public License. See LICENSE.txt for details.
 *
 * Note: other non-source-code works from Democracy Unchained are typically licensed differently
 * (typically Creative Commons with Attribution).  Please refer to the relevant licensing details
 * in those projects.
 *************************************************************************************************/

package main 

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
      "net/http"
      "encoding/json"
      "fmt"
//      "html"
      "log"
      "github.com/gorilla/mux"
//      "time"
	"strconv"
)

type State struct {
     Id		int		`json:"id"`
     Name  	string		`json:"name"`
//     Joined     time.Time	`json:"joined"`
     Joined	string		`json:"joined"`
     Is_state	bool		`json:"is_state"`
}

type States []State

type Population struct {
     People	int		`json:"people"`
     Year	int		`json:"year"`
     Type	string		`json:"type"`
}

type Populations []Population

type ElectoralVote struct {
     Votes	   int		`json:"votes"`
     CensusYear	   int		`json:"census_year"`
}

type ElectoralVotes []ElectoralVote

type StateReport struct {
     State	 	`json:"state"`
     ElectoralVotes	`json:"electoral_votes"`
     Population		`json:"population_recent"`
     Populations	`json:"populations"`
     
}

type StateReports []StateReport

// Global sql.DB to access the database by all handlers
var db *sql.DB 
var err error

func InitDB() {

    // Create an sql.DB and check for errors
    db, err = sql.Open("mysql", "readonly:cnc93y3ghniwy9384nfihsd93f938nds@/democracy")
    if err != nil {
       panic(err.Error())
       }			

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
       panic(err.Error())
    }			

}

func main() {
    
    InitDB()

      router := mux.NewRouter().StrictSlash(true)
      router.HandleFunc("/", Index)
      router.HandleFunc("/states", StatesIndex)
      router.HandleFunc("/states/{stateName}", StatesShow)

      log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintln(w, "Welcome!")
     }

func StatesIndex(w http.ResponseWriter, r *http.Request) {

     rows, err := db.Query("SELECT id,name,joined,is_state FROM states order by id asc")
     checkErr(err)

      states := States{}

      for rows.Next() {
              state := State{}
	      err := rows.Scan(&state.Id, &state.Name, &state.Joined, &state.Is_state)
	      checkErr(err)
	      states = append(states, state)
	      }

      json.NewEncoder(w).Encode(states)
    
}

func StateQuery(stateName string) string {

     if _, err := strconv.Atoi(stateName); err == nil {

       return "SELECT id,name,joined,is_state FROM states where id=?"

     } else {

       return "SELECT id,name,joined,is_state FROM states where name=?"

     }

     return "ERROR" // TODO

}

func StatesShow(w http.ResponseWriter, r *http.Request) {
     vars := mux.Vars(r)
     stateName := vars["stateName"]

    rows, err := db.Query(StateQuery(stateName), stateName)

    checkErr(err);

    // TODO: handle error (missing state)

      state_reports := StateReports{}

      for rows.Next() {

      	state_report := StateReport{}
	err := rows.Scan(&state_report.State.Id, &state_report.State.Name, &state_report.State.Joined, &state_report.State.Is_state)
	checkErr(err)

	population_rows, err := db.Query("SELECT people,year,type FROM populations where state_id=? order by year,type asc",state_report.State.Id)
	checkErr(err)

	for population_rows.Next() {
		population := Population{}
		err := population_rows.Scan(&population.People,&population.Year,&population.Type)
		checkErr(err)
		state_report.Populations = append(state_report.Populations,population)

		state_report.Population = population
	}	

	electoral_rows, err := db.Query("SELECT votes,census_year FROM electoral_votes where state_id=? order by census_year desc",state_report.State.Id)
	checkErr(err)

	for electoral_rows.Next() {
	    electoral_vote := ElectoralVote{}
	    err := electoral_rows.Scan(&electoral_vote.Votes,&electoral_vote.CensusYear)
	    checkErr(err)
	    state_report.ElectoralVotes = append(state_report.ElectoralVotes,electoral_vote)
	}

	state_reports = append(state_reports, state_report)      	  

      }

    json.NewEncoder(w).Encode(state_reports)

}

func checkErr(err error) {
     if err != nil {
     	panic(err)
	}
}