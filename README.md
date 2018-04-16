# dust_sensor

Arduino Mega connected to a dust sensor which triggers output if adjustable dust level is exceeded.

This repo activates a WebHook on a commit to master. It POST requests the above Jenkins server automatically, which triggers a build which compiles and runs the test code in this repo! It's pretty cool to set up my first Continuous Integration project!
Note, it will not build on commits to .md files.


