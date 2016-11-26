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
  "fmt"
  "github.com/gorilla/mux"
)

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

