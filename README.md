# dust_sensor

Arduino connected to a dust sensor which triggers output if adjustable dust level is exceeded.

Currently, no code has been written. Code will be written once the hardware has been finalised.

Connect to this Jenkins server at http://ec2-54-153-133-215.ap-southeast-2.compute.amazonaws.com:8080/

This repo is set up with a webhook. It POST requests the above jenkins server automatically, which triggers a build which compiles and runs the test code in this repo!

ssh into Jenkins instance using:
`ssh ubuntu@ec2-54-153-133-215.ap-southeast-2.compute.amazonaws.com`


