#include "Arduino.h"
#include "./HPMA115S0/src/HPMA115S0.h"

#define BAUDRATE 9600;
// These are used to give names to the pins used:
const int analogInPin = A0;  // Potentiometer Analog ping
const int analogOutPin = 9; 
int sensorValue = 0;        // value read from the pot
int outputValue = 0;        // value output to the PWM

int read_adc(); //definition

//Create an instance of the HPMA115S0 library
HPMA115S0 honeywell(Serial1);

void setup() {
  // initialize serial communications at 9600 bps:
  Serial.begin(BAUDRATE); //begin adc comms
  Serial1.begin(BAUDRATE);//begin honeywell comms
  while (!Serial1) {
    ; // wait for honeywell to connect
  }
  Serial.println("Starting...");
  honeywell.Init();
  honeywell.StartParticleMeasurement();
  Serial.println("Setup func complete!");
}


void loop() {
    // threshold dust level out of 100
    float threshold;
    threshold = read_adc() * 100.0f/255.0f;


}



void read_adc(){
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
