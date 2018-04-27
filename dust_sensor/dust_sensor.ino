/*
 * @file dust_sensor.ino
 * @author James Moessis
 * @date April 2018
 * 
 * See Program notes at bottom.
 * See README.md for more info.
 */

#include "Arduino.h"
#include <hpma115S0.h>
#include <TimeLib.h>

#define BAUDRATE 9600
#define BJT 2 // High when threshold exceeded - PNP
#define RED_LED 3 // Low when threshold exceeded (LED)
#define POWER_DUST 4

// GLOBALS
  time_t t;
  
  //Controlled by Potentiometer
  const int analogInPin = A0;  // Potentiometer Analog ping
  const int analogOutPin = 9; 
  short int sensorValue = 0;        // value read from the pot
  short int outputValue = 0;        // value output to the PWM
  float threshold;
  
  // Controlled by Dust Sensor
  double sum = 0; // note: maxes at 65,535.
  float circle[300]; // Dust measurements store
  size_t circle_size = sizeof(circle)/sizeof(circle[0]);
  unsigned int i; // index for circle


//Create an instance of the HPMA115S0 library
HPMA115S0 honeywell(Serial1);

void setup() {
  
  //t starts at 0 so dont need to record time
  pinMode(POWER_DUST, OUTPUT);
  pinMode(BJT, OUTPUT);
  pinMode(RED_LED, OUTPUT);
  
  //power cycle dust sensor
  digitalWrite(POWER_DUST, HIGH); 
  delay(3000);
  digitalWrite(POWER_DUST, LOW); 
  
  Serial.begin(BAUDRATE);
  Serial.println("Hello Computer.");
  
  // initialize all measurements as 0
  for( i = 0; i < circle_size; i++){
    circle[i] = 0;
  }
  i = 0; // i is global, reset to 0
  
  Serial1.begin(BAUDRATE); //begin honeywell comms
  do {
    Serial.println("Starting Serial1...");
    delay(5000); // wait 5s for honeywell to connect
  } while (!Serial1);

  //power control will be on pin 3, active low
  
  honeywell.Init();
  honeywell.StartParticleMeasurement();
  Serial.println("Setup func complete!");
}


void loop() {
    
    // Read Honeywell dust sensor
    unsigned int pm2_5, pm10;
    if (honeywell.ReadParticleMeasurement(&pm2_5, &pm10)) {
      Serial.println("PM 2.5: " + String(pm2_5) + " ug/m3" );
      Serial.println("PM 10: " + String(pm10) + " ug/m3" );
      
      // convert dust measurement to value out of 100
      circle[i] = pm10 * ( (float)100 / (float)1000 ); //problem here - circle not of type float
      Serial.println("circle[i] = " + String(circle[i]));
      Serial.println("i = " + String(i));
      delay(10);
    }
    
    // Increment iterator
    // If reached end of array, go back to beginning
    if(i >= circle_size - 1) {
      i = 0;
      Serial.println("Resetting i!");
      time_t lap_time = now() - t;
      Serial.print("lap_time = ");
      Serial.println(lap_time);
      t = now();
    }
    else {
      i++;
    }      

    //read analog
    //threshold dust level out of 100
    threshold = read_adc() * ( (float)100 / (float)1023 );
    Serial.println("Threshold = " + String(threshold));

    // Sum all recent measurements
    int k = 0;
    sum = 0;
    for(k = 0; k < circle_size; k++) {
      sum = sum + circle[k];
      //sum = sum + (float)99.9; //test maxed values
    }

    // Average all recent measurements
    float average; 
    average = ( (float)sum ) / ( (float)circle_size ); 
    Serial.println("Sum = " + String(sum));
    Serial.println("Avg = " + String(average));     
    
    
    // Active LOW output to PNP BJT
    // Red Light on
    if (average >= threshold) {
      digitalWrite(BJT, LOW);
      digitalWrite(RED_LED, HIGH);
      delay(1);
      Serial.println("Threshold Exceeded");
    }
    // Green light on
    else {
      digitalWrite(BJT, HIGH);
      digitalWrite(RED_LED, LOW);
      delay(1);
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
  Serial.print("Potentiometer output = ");
  Serial.println(outputValue);

  // wait 2 milliseconds before the next loop for the analog-to-digital
  // converter to settle after the last reading:
  delay(2);
  return sensorValue;
}

/***************************************************
 * Notes
 * Currently the use of memory is quite inefficient.
 * Floats are not necessarily needed, but were ok
 * in the current case.
 * 
 * To scale this program to do more, delays should
 * be altered, and memory efficiency should be tuned.
 * Additionally, the variables should be made local.
*****************************************************/


