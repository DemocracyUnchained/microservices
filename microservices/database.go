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
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error

func InitDB() {

    // Create an sql.DB and check for errors
    conn_fmt_str := Config.DB.Username+":%v@tcp("+Config.DB.Host+":%v)/"+Config.DB.Database
    conn_str := fmt.Sprintf(conn_fmt_str, Config.DB.Password, Config.DB.Port)
    fmt.Println(fmt.Sprintf(conn_fmt_str, "********", Config.DB.Port))
    db, err = sql.Open("mysql", conn_str)

    if err != nil {
       panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
       panic(err.Error())
    }

}

func checkErr(err error) {
    if err != nil {
      panic(err)
    }
}


