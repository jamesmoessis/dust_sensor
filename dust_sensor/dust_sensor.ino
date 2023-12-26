/*
 * @file dust_sensor.ino
 * @author James Moessis
 * @date April 2018
 * @date December 2023
 *
 * See README.md for more info.
 */

#include "Arduino.h"
#include <SPI.h>
#include <Ethernet.h>
#include <hpma115S0.h>  // https://github.com/felixgalindo/HPMA115S0
#include <TimeLib.h>    // https://github.com/PaulStoffregen/Time

#define BAUDRATE 9600
#define FET 2         // Trigger threshold. MOSFET. Actives red LED
#define GREEN_LED 3   // High when threshold note exceeded
#define POWER_DUST 4  // pnp bjt powering HPMA
#define BURST_SIZE 50
#define DELAY_BETWEEN_BURSTS_MS 60000
#define FAILURE_TOLERANCE 0.1

typedef struct DynamicSettings {
  bool on;
  int threshold;  // between 0 and 1000
} dynamic_settings;

typedef struct Measurements {
  int average;
  int maximum;
  int minimum;
  int threshold;
  int failure_count;
  time_t lap_time;
} measurements;

void reset_sensor();
int read_adc();
void set_settings();
void init_ethernet();
void start_sensor();
void stop_sensor();
void emit_data(measurements);

// GLOBALS
time_t t;

//Controlled by Potentiometer
const int analogInPin = A0;  // Potentiometer Analog ping (pin 54)
const int analogOutPin = 9;
short int sensorValue = 0;  // value read from the pot
short int outputValue = 0;  // value output to the PWM

//Create an instance of the HPMA115S0 class
HPMA115S0 honeywell(Serial1);
bool sensor_on = false;

// Enter a MAC address for your controller below.
// Newer Ethernet shields have a MAC address printed on a sticker on the shield
byte mac[] = { 0xDE, 0xAD, 0xBE, 0xEF, 0xFE, 0xED };

// name address for CloudFront that sits in front of API (using DNS)
char server[] = "d1d1khgtxr0hea.cloudfront.net";  

// Set the static IP address to use if the DHCP fails to assign
IPAddress ip(192, 168, 0, 177);
IPAddress myDns(192, 168, 0, 1);

// Initialize the Ethernet client library
// with the IP address and port of the server
// that you want to connect to (port 80 is default for HTTP):
EthernetClient client;

// Variables to measure the speed
unsigned long beginMicros, endMicros;
unsigned long byteCount = 0;
bool printWebData = true;  // set to false for better speed measurement

// Cloudfront distribution which turns HTTP into HTTPS to access backend API GW / Lambda
String hostname = "d1d1khgtxr0hea.cloudfront.net";
String port = "80";

dynamic_settings settings;

void setup() {
  // initialize pins
  pinMode(POWER_DUST, OUTPUT);
  pinMode(FET, OUTPUT);
  pinMode(GREEN_LED, OUTPUT);

  digitalWrite(FET, LOW);
  digitalWrite(GREEN_LED, HIGH);

  Serial.begin(BAUDRATE);
  Serial.println("Hello Computer.");

  Serial1.begin(BAUDRATE);  //begin honeywell comms
  do {
    Serial.println("Starting Serial1...");
    delay(5000);  // wait 5s for honeywell to connect
  } while (!Serial1);

  init_ethernet();

  settings.on = true;
  settings.threshold = 1000;

  Serial.println("Setup func complete!");
}


void loop() {
  set_settings();

  Serial.println("IsOn: " + String(settings.on));
  Serial.println("Threshold: " + String(settings.threshold));

  if (settings.on) {
    start_sensor();

    // Record burst of measurements
    int dust_levels[BURST_SIZE];
    int failure_count = 0;
    unsigned int pm2_5, pm10;
    for (int k = 0; k < BURST_SIZE; k++) {
      dust_levels[k] = 0;
      // Read Honeywell dust sensor
      if (honeywell.ReadParticleMeasurement(&pm2_5, &pm10)) {
        Serial.println("PM 2.5: " + String(pm2_5) + " ug/m3");
        Serial.println("PM 10: " + String(pm10) + " ug/m3");

        // pm10 is between 0 and 1000
        dust_levels[k] = pm10;
        Serial.println("circle[" + String(k) + "] = " + String(dust_levels[k]));
      } else {
        dust_levels[k] = -1;  // indicates a reading which should not be counted
        failure_count++;
      }
      delay(200);
    }

    //stop_sensor();

    for (int i = 0; i < BURST_SIZE; i++) {
      Serial.print(String(dust_levels[i]));
    }
    Serial.println();

    if (failure_count > FAILURE_TOLERANCE * BURST_SIZE) {
      Serial.println("Warning: dust sensor readings failed " + String(failure_count) + " times in the current burst.");
    }

    measurements m;
    int sum = 0;
    m.maximum = 0;
    m.minimum = 10000;
    m.threshold = settings.threshold;
    m.failure_count = failure_count;
    for (int k = 0; k < BURST_SIZE; k++) {
      int val = dust_levels[k];
      if (val != -1) {
        sum += val;
        if (val > m.maximum) {
          m.maximum = val;
        }
        if (val < m.minimum) {
          m.minimum = val;
        }
      }
    }
    m.average = sum / BURST_SIZE;

    if (m.average >= settings.threshold) {
      digitalWrite(GREEN_LED, LOW);  // Red Light on
      digitalWrite(FET, HIGH);       // enable alarm
      Serial.println("Threshold Exceeded");
    } else {
      digitalWrite(GREEN_LED, HIGH);  // Green light on
      digitalWrite(FET, LOW);         // disable alarm
    }

    m.lap_time = now() - t;
    t = now();  // time in sec since program started
    
    emit_data(m);
  } else {
    //stop_sensor();
  }

  delay(DELAY_BETWEEN_BURSTS_MS);
}

void set_settings() {
  if (client.connect(server, 80)) {
    Serial.print("connected to ");
    Serial.println(client.remoteIP());
    // Make a HTTP request:
    client.println("GET /api/settings HTTP/1.1");
    client.println("Host: " + String(server));
    client.println("User-Agent: Arduino");
    client.println("Connection: close");
    client.println();
  } else {
    // if you didn't get a connection to the server:
    Serial.println("connection failed");
    return;
  }

  bool success = false;
  String line = "";
  String first_line = "";
  while(client.connected()) {
    line = client.readStringUntil('\n');
    line.trim();
    if (success == 0) {
      first_line = line;
      if (line.endsWith("200 OK")) {
        success = true;
      }
    }
  }

  client.stop();

  if (!success) {
    Serial.println("Response failed. First line: " + first_line);
    return;
  } else {
    Serial.println("Last line (body): " + line);
  }

  // EXTREMELY janky JSON parsing
  // Relies on response being all on one line and no whitespace 
  // (which we can ensure because we have written the backend)
  // todo make this more clean or use a library if enough memory
  int start = line.indexOf("isOn\":") + 6;
  int end = line.indexOf(',');
  bool isOn = line.substring(start, end).equals("true");

  start = line.indexOf("threshold\":") + 11;
  end = line.indexOf("}"); // we know threshold is last value
  int threshold = line.substring(start, end).toInt();

  settings.on = isOn;
  settings.threshold = threshold;
}

void emit_data(measurements m) {
  if (client.connect(server, 80)) {
    Serial.print("connected to ");
    Serial.println(client.remoteIP());
    String query_string = "average=" + String(m.average) + 
      "&failurecount=" + String(m.failure_count) + 
      "&laptime=" + String(m.lap_time) + 
      "&minimum=" + String(m.minimum) + 
      "&maximum=" + String(m.maximum) +
      "&threshold=" + String(m.threshold); 

    
    // Make a HTTP request:
    Serial.println("query string: " + query_string);
    client.println("POST /api/measurements?" + query_string + " HTTP/1.1");
    client.println("Host: " + String(server));
    client.println("User-Agent: Arduino");
    client.println("Connection: close");
    client.println();
  } else {
    // if you didn't get a connection to the server:
    Serial.println("connection failed");
    return;
  }

  bool success = false;
  String line = "";
  String first_line = "";
  while(client.connected()) {
    line = client.readStringUntil('\n');
    line.trim();
    if (success == 0) {
      first_line = line;
      if (line.endsWith("200 OK")) {
        success = true;
      }
    }
  }
  if (!success) {
    Serial.println("Failed to report measurements: " + first_line);
  }
  Serial.println("Average: " + String(m.average));
  Serial.println("Max: " + String(m.maximum));
  Serial.println("Min: " + String(m.minimum));
  Serial.println("Threshold: " + String(m.threshold));
  Serial.println("Failure count: " + String(m.failure_count));
  Serial.println("Lap time: " + String(m.lap_time) + "s");
}

//adapted from example Arduino sketch
int read_adc() {
  // read the analog in value:
  sensorValue = analogRead(analogInPin);
  // map it to the range of the analog out:
  outputValue = map(sensorValue, 0, 1023, 0, 255);
  // change the analog out value:
  analogWrite(analogOutPin, outputValue);

  // wait 2 milliseconds before the next loop for the analog-to-digital
  // converter to settle after the last reading:
  delay(2);
  return sensorValue;
}

void reset_sensor() {
  Serial.println("Resetting Honeywell Sensor...");
  stop_sensor();
  delay(3000);
  start_sensor();
}

void stop_sensor() {
  if (sensor_on) {
    Serial.println("Powering off Honeywell Sensor...");
    honeywell.StopParticleMeasurement();
    delay(100);
    digitalWrite(POWER_DUST, HIGH);
    sensor_on = false;
  }
}

void start_sensor() {
  if (!sensor_on) {
    Serial.println("Powering on Honeywell Sensor...");
    digitalWrite(POWER_DUST, LOW);
    delay(100);
    honeywell.Init();
    sensor_on = true;
  }
}

void init_ethernet() {
  Serial.println("Initialize Ethernet with DHCP:");
  if (Ethernet.begin(mac) == 0) {
    Serial.println("Failed to configure Ethernet using DHCP");
    // Check for Ethernet hardware present
    if (Ethernet.hardwareStatus() == EthernetNoHardware) {
      Serial.println("Ethernet shield was not found.  Sorry, can't run without hardware. :(");
      while (true) {
        delay(1);  // do nothing, no point running without Ethernet hardware
      }
    }
    if (Ethernet.linkStatus() == LinkOFF) {
      Serial.println("Ethernet cable is not connected.");
    }
    // try to configure using IP address instead of DHCP:
    Ethernet.begin(mac, ip, myDns);
  } else {
    Serial.print("  DHCP assigned IP ");
    Serial.println(Ethernet.localIP());
  }
}
