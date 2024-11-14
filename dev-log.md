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
- no relevant information is returned by attack end point

want:
- end point that returns all active games (just id) DONE
- ephemeral active battles (store in db until battle is over, maybe)
