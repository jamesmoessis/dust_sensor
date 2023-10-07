# dust_sensor

Arduino Mega connected to a dust sensor which triggers output if adjustable dust level is exceeded.

## How it works

### On a hardware level

* A potentiometer is to adjust acceptable dust level threshold. The Arduino's ADC reads the output of the Potentiometer, and the program uses it as a coefficient. If you want to control the coefficient remotely, the `SCALE` variable can be changed. 
* The dust sensor is power cycled at the start of the program. This is done via a PNP BJT (active low).
* A MOSFET (Active HIGH) is activated when averaged measurements exceed desired threshold level. This shorts an external input.
* A green and red LED are used to indicate threshold being exceeded (or not exceeded).
* The dust levels are communicated by the Honeywell HPMA115S0 via UART protocol.
* The Arduino Dust RX pin is pulled down toward ground via a 2k4 resistor. The TX is pulled up via a 2k4 resistor. This is to accomodate the difference in the communication voltage between the Arduino and the HPMA115S0. Arduino communicates via 5V and the HPMA via 3.3V. 
* This all will be indicated via the schematic.

### On a software level

* The third party hpma115S0.h Arduino library is used to communicate with the sensor. https://github.com/felixgalindo/HPMA115S0
* The main code is in the Arduino sketch dust_sensor.ino
* Treshold dust level level can be scaled via changing `SCALE`. Which must be defined as a `float`. 
* The `setup()` function establishes Serial communication to the HPMA115S0 and the Arduino Serial monitor. It runs initialization protocol to the dust sensor. It also initializes the trigger pin (to the BJT) as an output. It also power cycles the dust sensor.
* The dust measurements are stored in `circle[]`. This circularly loops through, and will be up to date with the 300 most recent measurements.
* The main `loop()` function in order:
    1. Reads dust sensor. Assigns to `circle[i]`.
    2. Reads ADC, converts reading into `threshold`, a percentage of the maximum possible value.
    3. Sums all 300 measurements in `circle[]` and takes `average`.
    4. If `average` exceeds `SCALE * threshold`, trigger output to PNP BJT. (Active LOW)
    5. Go back to step 1.

## Monitoring
Integration has been added with MegunoLink monitoring software.
* The Arduino will output values through the USB and MegunoLink (Once set up properly on a port) will graph and log these values.
* MegunoLink will generate a timestamped csv file and also create a graph against time. Every 300 particle measurements, the threshold and average level is sent. This is generally about every 1.5 minutes.
* The MegunoLink file can be found in the repo as *.mlx file. This file will need to be edited from computer to computer.
	
	
## Remote Control

Edit threshold by scaling it. Remotely edit the SCALE variable and then upload the new code.
