## RetestPrototype

Tutorial for creating a retest feature to be added to a CI pipeline. 

Once this app is running, any comment on a pull request will send a webhook to this web server. If the comment contains the phrase `/retest`, this app will then comment `Retest Initiated` under that comment.

This is the first step to implementing a retest feature into a continuous testing pipeline. Instead of re-commenting Retest Initiated, you could use that part of the code to re-run all of the tests for a given branch in a pull request.

The port 8080 is exposed to the internet using ngrok. The web server is then setup to listen on 8080.
