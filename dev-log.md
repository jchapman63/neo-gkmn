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
- Refactor channel logic into a battle service, handler should not deal with concurrency

change of direction.  I think I need to focus on better abstractions and data availability/locality.  For example,
when a battle happens, I should fetch everything I need for that battle out of the database and store that in the active battle.

I refactored the battle struct so that it makes use of maps of all data by id of the item.  This lets me quickly access monsters, move, etc
