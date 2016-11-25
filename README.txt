			Democracy Unchained - Microservices Architecture

Copyright 2016 Democracy Unchained -- http://DemocracyUnchained.org

Licensed under the GNU Affero General Public License. See LICENSE.txt for details.

Note: other non-source-code works from Democracy Unchained are typically licensed differently
(typically Creative Commons with Attribution).  Please refer to the relevant licensing details
in those projects.

REST Services

All services provide a JSON response package.

/states/
Returns an array of all the states, along with an ID number that may be used to refer to each one.

/states/stateName
Returns information about the state requested.  The stateName may be the name of the state (e.g., "California") or an ID corresponding to the state (as returned by /states/).  This includes expanded information such as population trends and electoral college information about the state in question.

/zipcodes/zipCode
This returns general information about zipCode:
	zip:		the zipCode requested
	city:		name of the city for that zipCode
	county:		name of the county
