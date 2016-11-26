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

type Vs struct {
  Name  string                    `json:"name"`
  Votes int                       `json:"votes"`
  VotingEligiblePopulation int    `json:"voting_eligible_population"`
  BallotsCounted int              `json:"ballots_counted"`
  PercentAsPowerful float32       `json:"percent_as_powerful"`
  Turnout                   float32 `json:"turnout"`
}
