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

type ZipCode struct {
  Zip     string  `json:"zip"`
  County  string  `json:"county"`
  City    string  `json:"city"`
}

type ZipCodes []ZipCode

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
