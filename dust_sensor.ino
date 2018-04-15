/*
 * @file dust_sensor.ino
 * @author James Moessis
 * @date April 2018
 */

#include "Arduino.h"
#include <hpma115S0.h>

#define BAUDRATE 9600

// These are used to give names to the pins used:
const int analogInPin = A0;  // Potentiometer Analog ping
const int analogOutPin = 9; 
int sensorValue = 0;        // value read from the pot
int outputValue = 0;        // value output to the PWM
float threshold;

//int read_adc(); //definition

//Create an instance of the HPMA115S0 library
HPMA115S0 honeywell(Serial1);

void setup() {
  // initialize serial communications at 9600 bps:
  Serial.begin(BAUDRATE); // begin comms with USB (to Arduino Serial Monitor)
  Serial.println("Hello Computer.");
  Serial1.begin(BAUDRATE); //begin honeywell comms
  do {
    delay(5000); // wait for honeywell to connect
    Serial.println("Starting Serial1...");
  } while (!Serial1);
  honeywell.Init();
  honeywell.StartParticleMeasurement();
  Serial.println("Setup func complete!");
}


void loop() {
    // threshold dust level out of 100
    threshold = read_adc() * 100.0f/255.0f;
    
    //honeywell read
    unsigned int pm2_5, pm10;
    if (honeywell.ReadParticleMeasurement(&pm2_5, &pm10)) {
    Serial.println("PM 2.5: " + String(pm2_5) + " ug/m3" );
    Serial.println("PM 10: " + String(pm10) + " ug/m3" );
  }
  
}

//adapted from example Arduino sketch
int read_adc(){
  // read the analog in value:
  sensorValue = analogRead(analogInPin);
  // map it to the range of the analog out:
  outputValue = map(sensorValue, 0, 1023, 0, 255);
  // change the analog out value:
  analogWrite(analogOutPin, outputValue);

  // print the results to the Serial Monitor:
  Serial.print("sensor = ");
  Serial.print(sensorValue);
  Serial.print("\t output = ");
  Serial.println(outputValue);

  // wait 2 milliseconds before the next loop for the analog-to-digital
  // converter to settle after the last reading:
  delay(2);
  return outputValue;
}

