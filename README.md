# dust_sensor

Arduino connected to a dust sensor which triggers output if adjustable dust level is exceeded.

Currently, no code has been written. Code will be written once the hardware has been finalised.

This repo is set up with a webhook. It POST requests the above Jenkins server automatically, which triggers a build which compiles and runs the test code in this repo! It's pretty cool to set up my first Continuous Integration project!
Note, it will not build on commits to .md files.


