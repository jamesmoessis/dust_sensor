# dust_sensor

Arduino Mega connected to a dust sensor which triggers output if adjustable dust level is exceeded.

## How it works

### On a hardware level

* The dust sensor is power cycled at the start of the program. This is done via a PNP BJT (active low).
* A MOSFET (Active HIGH) is activated when averaged measurements exceed desired threshold level. This shorts an external input.
* A green and red LED are used to indicate threshold being exceeded (or not exceeded).
* The dust levels are communicated by the Honeywell HPMA115S0 via UART protocol.
* The Arduino Dust RX pin is pulled down toward ground via a 2k4 resistor. The TX is pulled up via a 2k4 resistor. This is to accomodate the difference in the communication voltage between the Arduino and the HPMA115S0. Arduino communicates via 5V and the HPMA via 3.3V. 
* This is indicated via the schematic [DESIGN_SCHEMATIC.pdf](./DESIGN_SCHEMATIC.pdf).

### On a software level

* The third party hpma115S0.h Arduino library is used to communicate with the sensor. https://github.com/felixgalindo/HPMA115S0
* The main code is in the Arduino sketch dust_sensor.ino
* The measurements are taken in bursts, separated by one minute break. The sensor is powered off in between bursts to preserve its lifetime. 

## Monitoring

todo
	
## Remote Control

todo 

## Web User Interface

* `frontend/` contains a basic react application created with `npx create-react-app`. 
* The repository does not contain the module files to run the react app. Run the command `npm install` while in the directory to install necessary dependencies which are listed in `package.json`.
* Once it's installed, running `npm start` in the command line will start the application and open it in a localhost browser window. 
* More information on scripts is provided in the README file provided by react in `frontend/`