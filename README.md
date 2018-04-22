# dust_sensor

Arduino Mega connected to a dust sensor which triggers output if adjustable dust level is exceeded.

## How it works

### On a hardware level

* A potentiometer is used to adjust acceptable dust level threshold. The Arduino's ADC reads the output of the Potentiometer, and the program uses it as a coefficient.
* A BJT is activated when averaged measurements exceed desired threshold level. This shorts an external input.
* The BJT is in parallel with a LED to indicate when levels have been exceeded.
* The dust levels are communicated by the Honeywell HPMA115S0 via UART protocol.
* The Arduino Dust RX pin is pulled down toward ground via a 2k4 resistor. The TX is pulled up via a 2k4 resistor. This is to accomodate the difference in the communication voltage between the Arduino and the HPMA115S0. Arduino communicates via 5V and the HPMA via 3.3V. This all will be indicated via the schematic.

### On a software level

* The third party hpma115S0.h Arduino library is used to communicate with the sensor. https://github.com/felixgalindo/HPMA115S0
* The `setup()` function establishes Serial communication to the HPMA115S0 and the Arduino Serial monitor. It runs initialization protocol to the dust sensor. It also initializes the trigger pin (to the BJT) as an output.
* The dust measurements are stored in `circle[]`. This circularly loops through, and will be up to date with the 300 most recent measurements.
* The main `loop()` function in order:
1. Reads dust sensor. Assigns to `circle[i]`.
2. Reads ADC, converts reading into `threshold`, a percentage of the maximum possible value.
3. Sums up all 300 measurements in `circle[]` and takes `average`.
4. If `average` exceeds `threshold`, trigger output to BJT.
5. Go back to step 1.

## Remote Control

This repo activates a WebHook on a commit to master. On a commit to master, a Jenkins job compiles and runs the test code in this repo!
Note, it will not build on commits to .md files.


