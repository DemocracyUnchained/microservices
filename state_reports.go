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

type StateReport struct {
  State	 	                        `json:"state"`
  ElectoralVotes	                `json:"electoral_votes"`
  Population		                `json:"population_recent"`
  Populations	                    `json:"populations"`     
  Voter                             `json:"voters"`
  Vs                                `json:"vs"`
}

type StateReports []StateReport
