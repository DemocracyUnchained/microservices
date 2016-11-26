			Democracy Unchained - Microservices Architecture

Copyright 2016 Democracy Unchained -- http://DemocracyUnchained.org

Licensed under the GNU Affero General Public License. See LICENSE.txt for details.

Note: other non-source-code works from Democracy Unchained are typically licensed differently
(typically Creative Commons with Attribution).  Please refer to the relevant licensing details
in those projects.

INTRODUCTION

This is the Democracy Unchained open source project for our democracy microservices platform. The purpose of
these microservices is to provide a consolidated view into the data pertaining to US election processes.
Examples include: electoral college structure; voting patterns; ad spending & campaign appearances; population
demographics; voter eligibility; etc. These microservices may be used to explore data related to US
democratic processes and institutions, and enables the creation of other data-driven applications.

The easiest way to access our microservices is to simply use them, here: http://api.democracyunchained.org:8080

However, we are always looking for more people who would like to contribute to our open source project. We hope
to add more data types and analysis over time, and you can be a part of it!  For more information, visit:
http://DemocracyUnchained.org

SET UP

For doing local development on our microservices architecture (or to create your own local fork to experiment with
our software) you'll need to follow these steps:

  1. Install MySQL 5 or later
  2. Create a database to contain the Democracy Unchained dataset
  3. Obtain the Democracy Unchained dataset from here: [TODO]
  4. Import the Democracy Unchained dataset to the database you just created
  5. Install GoLang 1.6 or later on your computer
  6. Create a directory for this project
  7. Do a "go build" in this directory; resolve any dependencies
  8. Copy config.toml to /etc/config.toml and edit your database and server configuration options
  9. Run as follows:  ./democracy
  10. Have fun hacking!

REST Services - Basic Documentation

All services provide a JSON response package.

/states/
Returns an array of all the states, along with an ID number that may be used to refer to each one.

/states/stateName
Returns information about the state requested.  The stateName may be the name of the state (e.g., "California") or an ID corresponding to the state (as returned by /states/).  This includes expanded information such as population trends and electoral college information about the state in question.

/voters/stateId
Tells you information about the voters in the state for each election we know about (eligible voters, population of voting age, and how many were ineligible felons)
http://35.163.215.7:8080/voters/6

/zipcodes/zipCode
This returns general information about zipCode:
	zip:		the zipCode requested
	city:		name of the city for that zipCode
	county:		name of the county

