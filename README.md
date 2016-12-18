# Democracy Unchained - Microservices Architecture

Copyright 2016 Democracy Unchained -- http://DemocracyUnchained.org

Licensed under the GNU Affero General Public License. See LICENSE.txt for details.

Note: other non-source-code works from Democracy Unchained are typically licensed differently (typically Creative Commons with Attribution).  Please refer to the relevant licensing details in those projects.

## INTRODUCTION

This is the Democracy Unchained open source project for our democracy microservices platform. The purpose of these microservices is to provide a consolidated view into the data pertaining to US election processes.
Examples include:

1. electoral college structure
2. voting patterns
3. ad spending & campaign appearances
4. population demographics
5. voter eligibility

These microservices may be used to explore data related to US democratic processes and institutions, and enables the creation of other data-driven applications.

The easiest way to access our microservices is to simply use them, here: http://api.democracyunchained.org:8080

However, we are always looking for more people who would like to contribute to our open source project. We hope to add more data types and analysis over time, and you can be a part of it!  For more information, visit:
http://DemocracyUnchained.org

## SET UP

For doing local development on our microservices architecture (or to create your own local fork to experiment with our software) you'll need to follow these steps:

1. Install Docker.
  * On Linux, you'll need to install Docker Compose separately. See the instructions [here][https://docs.docker.com/compose/install/].
  * On Windows/Mac, use [Docker Toolbox][https://www.docker.com/products/docker-toolbox] to install.
2. Clone the microservices project:
  * `git clone https://github.com/DemocracyUnchained/microservices.git`
3. Change into the docker directory:
  * `cd microservices/docker`
4. Initialize your database:
  * `make init_db`
  ** If your local machine is already running MySQL listening on port 3306, this step will fail. To fix the failure, either shutdown you local MySQL instance *or* change the port used by your local MySQL instance or the Docker MySQL container.
5. Start the project containers:
  * `make start`
6. Point your browser to the development instance here:
  * <http://localhost:8089/states>
6. Profit! And have fun hacking!

## REST Services - Basic Documentation

All services provide a JSON response package.

`/states/`

Returns an array of all the states, along with an ID number that may be used to refer to each one.

`/states/stateName`

Returns information about the state requested.  The stateName may be the name of the state (e.g., "California") or an ID corresponding to the state (as returned by /states/).  This includes expanded information such as population trends and electoral college information about the state in question.

`/voters/stateId`

Tells you information about the voters in the state for each election we know about (eligible voters, population of voting age, and how many were ineligible felons)

`/zipcodes/zipCode`

This returns general information about zipCode:

* zip: the zipCode requested
* city:	name of the city for that zipCode
* county: name of the county
