
/************************************************************************************************
 * Copyright 2016 Democracy Unchained -- http://DemocracyUnchained.org
 *
 * Licensed under the GNU Affero General Public License. See LICENSE.txt for details.
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
  "./config"
)

// Code to implement the TextUnmarshaler interface for `duration`:
//
// type duration struct {
//  time.Duration
// }
//
// func (d *duration) UnmarshalText(text []byte) error {
//  var err error
//  d.Duration, err = time.ParseDuration(string(text))
//  return err
// }

for _, s := range favorites.Song {
    fmt.Printf("%s (%s)\n", s.Name, s.Duration)
}

type ZipCode struct {
  Zip     string  `json:"zip"`
  County  string  `json:"county"`
  City    string  `json:"city"`
}

type ZipCodes []ZipCode

type State struct {
  Id		int		`json:"id"`
  Name  	string		`json:"name"`
  //     Joined     time.Time	`json:"joined"`
  Joined	string		`json:"joined"`
  Is_state	bool		`json:"is_state"`
}

type States []State

type Population struct {
  People  int       `json:"people"`
  Year    int       `json:"year"`
  Type    string    `json:"type"`
}

type Populations []Population

type ElectoralVote struct {
  Votes       int   `json:"votes"`
  CensusYear  int   `json:"census_year"`
}

type ElectoralVotes []ElectoralVote

type Voter struct {
  StateId                   int     `json:"state_id"`
  Source                    string  `json:"source"`
  ElectionId                int     `json:"election_id"`
  BallotsCounted            int     `json:"ballots_counted"`
  VotingEligiblePopulation  int     `json:"voting_eligible_population"`
  VotingAgePopulation       int     `json:"voting_age_population"`
  IneligiblePrison          int     `json:"ineligible_prison"`
  IneligibleProbation       int     `json:"ineligible_probation"`
  IneligibleParole          int     `json:"ineligible_parole"`
  Turnout                   float32 `json:"turnout"`
}

type Voters []Voter

type Vs struct {
  Name  string                    `json:"name"`
  Votes int                       `json:"votes"`
  VotingEligiblePopulation int    `json:"voting_eligible_population"`
  BallotsCounted int              `json:"ballots_counted"`
  PercentAsPowerful float32       `json:"percent_as_powerful"`
  Turnout                   float32 `json:"turnout"`
}

type StateReport struct {
  State	 	                          `json:"state"`
  ElectoralVotes	                  `json:"electoral_votes"`
  Population		                    `json:"population_recent"`
  Populations	                      `json:"populations"`     
  Voter                             `json:"voters"`
  Vs                                `json:"vs"`
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

  // /states
  router.HandleFunc("/states", StatesIndex)
  router.HandleFunc("/states/{stateName}", StatesShow)

  // /voters
  router.HandleFunc("/voters", VotersIndex)
  router.HandleFunc("/voters/{stateId}", VotersShow)

  // /zipcodes
  router.HandleFunc("/zipcodes", ZipCodeIndex)
  router.HandleFunc("/zipcodes/{zipCode}", ZipCodeShow)

  log.Fatal(http.ListenAndServe(":8080", router))

}

func Index(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintln(w, "Welcome!")
     }

func ZipCodeIndex(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintln(w, "TODO")
}

func ZipCodeShow(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  zipCode := vars["zipCode"]

  rows, err := db.Query("SELECT zip,city,county from zip_codes WHERE zip=?", zipCode)

  checkErr(err);

  zip_code := ZipCode{}
  for rows.Next() {
    err := rows.Scan(&zip_code.Zip,&zip_code.City,&zip_code.County)
    checkErr(err)
  }

  json.NewEncoder(w).Encode(zip_code)

}

func VotersIndex(w http.ResponseWriter, r *http.Request) {
     fmt.Fprintln(w, "TODO")
}

func VotersShow(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)
  stateId := vars["stateId"]

  rows, err := db.Query("SELECT state_id,source,election_id,ballots_counted,voting_eligible_population,voting_age_population,ineligible_prison,ineligible_probation,ineligible_parole FROM voters WHERE state_id=?", stateId)

  checkErr(err);

  voters := Voters{}

  for rows.Next() {
    voter := Voter{}
    err := rows.Scan(&voter.StateId,&voter.Source,&voter.ElectionId,&voter.BallotsCounted,&voter.VotingEligiblePopulation,&voter.VotingAgePopulation,&voter.IneligiblePrison,&voter.IneligibleProbation,&voter.IneligibleParole)
    checkErr(err)
    voter.Turnout = float32(voter.BallotsCounted) / float32(voter.VotingEligiblePopulation)
    voters = append(voters,voter)
  }

  json.NewEncoder(w).Encode(voters)

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

    // Show the most recent election for this state
    // TODO:  really choose the correct one [right now the 2016 presidential election is hardcoded]

    voter_rows, err := db.Query("SELECT state_id,source,election_id,ballots_counted,voting_eligible_population,voting_age_population,ineligible_prison,ineligible_probation,ineligible_parole FROM voters WHERE state_id=? AND election_id=1", state_report.State.Id)
    checkErr(err)

    for voter_rows.Next() {
      voter := Voter{}
      err := voter_rows.Scan(&voter.StateId,&voter.Source,&voter.ElectionId,&voter.BallotsCounted,&voter.VotingEligiblePopulation,&voter.VotingAgePopulation,&voter.IneligiblePrison,&voter.IneligibleProbation,&voter.IneligibleParole)
      checkErr(err)
      voter.Turnout = float32(voter.BallotsCounted) / float32(voter.VotingEligiblePopulation)
      state_report.Voter = voter
    }

    population_rows, err := db.Query("SELECT people,year,type FROM populations where state_id=? order by year,type asc",state_report.State.Id)
    checkErr(err)

    // Show population history for this state

    for population_rows.Next() {
		  population := Population{}
		  err := population_rows.Scan(&population.People,&population.Year,&population.Type)
		  checkErr(err)
		  state_report.Populations = append(state_report.Populations,population)
  		state_report.Population = population
    }	

    // Show the most recent electoral history for this state

    electoral_rows, err := db.Query("SELECT votes,census_year FROM electoral_votes where state_id=? order by census_year desc",state_report.State.Id)
    checkErr(err)

    for electoral_rows.Next() {
	   electoral_vote := ElectoralVote{}
	   err := electoral_rows.Scan(&electoral_vote.Votes,&electoral_vote.CensusYear)
	   checkErr(err)
	   state_report.ElectoralVotes = append(state_report.ElectoralVotes,electoral_vote)
    }

    // Compare to the most powerful state (Wyoming)

    powerful_rows, err := db.Query("SELECT name,votes,voting_eligible_population,ballots_counted FROM electoral_votes,states,voters where electoral_votes.state_id=states.id AND voters.state_id=states.id AND states.name='Wyoming' AND electoral_votes.census_year=2010")
    checkErr(err)
    for powerful_rows.Next() {
      err := powerful_rows.Scan(&state_report.Vs.Name,&state_report.Vs.Votes,&state_report.Vs.VotingEligiblePopulation,&state_report.Vs.BallotsCounted)
      checkErr(err)
      state_report.Vs.PercentAsPowerful = (float32(state_report.Vs.VotingEligiblePopulation)/float32(state_report.Vs.Votes)) / (float32(state_report.Voter.VotingEligiblePopulation) / float32(state_report.ElectoralVotes[0].Votes))
      state_report.Vs.Turnout = float32(state_report.Vs.BallotsCounted) / float32(state_report.Vs.VotingEligiblePopulation)
    }

    // Append the state report to the array
    state_reports = append(state_reports, state_report)      	  

  }

  json.NewEncoder(w).Encode(state_reports)

}

func checkErr(err error) {
     if err != nil {
     	panic(err)
	}
}