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
  "net/http"
  "fmt"
  "log"
  "github.com/gorilla/mux"
  "strconv"
)

type MyServer struct {
    r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }
    // Lets Gorilla work
    s.r.ServeHTTP(rw, req)
}

func InitHTTP() {

  router := mux.NewRouter().StrictSlash(true)

  // /states
  router.HandleFunc("/states", StatesIndex)
  router.HandleFunc("/states/{stateName}", StatesShow)

  // /voters
  router.HandleFunc("/voters", VotersIndex)
  router.HandleFunc("/voters/{stateId}", VotersShow)

  // /zipcodes
  router.HandleFunc("/zipcodes", ZipCodeIndex)
  router.HandleFunc("/zipcodes/{zipCode}", ZipCodeShow)

  http.Handle("/", &MyServer{router})

  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(Config.Server.Port), nil))

}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

