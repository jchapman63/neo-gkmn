# Developer Logs

I quickly realized that I work in spread out chunks.  Some times its weeks before I pick a project back up.
Life is busy out of college and as a full time engineer.  This project is fun and I think deserves to be fully
fleshed, deployed, and played with.  To do so, I will need to work with me, not against me.  Here is a log of changes
as I make them.

### 11-13-24
Finished up a refactor of the gkmn service. I split out the handler, endpoints, and util functions into files of their
own.

behavior notes:
- non-existent battle id returns empty mon list
with that id (okay, but useless). 404 type of response is better to use from connect
- no relevant information is returned by attack end point. THIS SHOULD NOT BE THE CASE. FIXED.

want:
- end point that returns all active games (just id) DONE
- ephemeral active battles (store in db until battle is over, maybe)

### 11-23-24
Fixed:
- attack endpoint now returns relevant info
- attack endpoint is accurate

want:
- connect errors on bad requests (e.g. 404 for nonexistent objs)
- timeout logic for move requests that take too long to complete.

### 11-19-42
refactor service dir to just have handler and connect dir to have gen proto code.

### 11-24-24
Goal:
~~ - Refactor channel logic into a battle service, handler should not deal with concurrency ~~
- implement priority queue decision making when a player takes a turn - WIP
- TODO: Fix pqeue typing error.  The "Item" type does not match BattleMon.

change of direction.  I think I need to focus on better abstractions and data availability/locality.  For example,
when a battle happens, I should fetch everything I need for that battle out of the database and store that in the active battle.

I refactored the battle struct so that it makes use of maps of all data by id of the item.  This lets me quickly access monsters, move, etc

DO NOT START ANYTHING NEW UNTIL YOU FIX PQUEUE

### 11-29-24
- Work on PQUEUE
    - battle tests with mockery
- Think about battle client

### 12-1-24
- Used mockery to create Querier mocks
- TODO: Finish up stat and movemap tests and then work on PQUEUE bug

### 12-11-24
- Finished movemap and stat tests
- TODO: Implement battle tests with mockery
- TODO: Implement client

## Client / Server Architecture
I want to write a client that makes direct use of my connectrpc generated code.  This would be just like how I generated a
TypeScript client for bodata except the code is already generated. Here are my rough ideas:
- The client is called by running neo-gkmn's binary with a client arg
- This starts a client instance, which can become a member of a battle or create a battle
- These actions are done by leveraging the server, another instance of the noe-gkmn binary with a server arg,
which leverages the connect client to create and send rpc requests to the server

The client app here would be a way to play games and connect to the server with no UI, just by sending client commands.
If I want a web app, I would be reimplementing the same logic, polling the API for game state. 

### 12-25-24

## Ebitengine

I started looking into web assembly recently and was interested in the idea of "close to" native compiling go in the web browser.
I saw that the common use case was games in the browser and thought it would be nice to write the graphics of neo-gkmn in golang.
I did a quick search and found Ebitengine on reddit.  Here are their docs: https://ebitengine.org/en/documents/cheatsheet.html

Today, I am going to copy over their hello world program and play around with the code in neo-gkmn to see what happens.  I think
I will be able to implement something similar to what I described on 12-11-24 and have a client host the Ebitengine logic.

