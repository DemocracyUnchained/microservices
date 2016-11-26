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
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "strconv"
  "github.com/gorilla/mux"
)

type State struct {
	Id			int			`json:"id"`
	Name  		string		`json:"name"`
	Joined		string		`json:"joined"`
	Is_state	bool		`json:"is_state"`
}

type States []State

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

