int readPot(void);

// These are used to give names to the pins:
const int analogInPin = A0;  // Analog input pin that the potentiometer is attached to
const int analogOutPin = 9; // Analog output pin that the LED is attached to

int potValue = 0;        // value read from the pot
int outputPot = 0;       // value output to the PWM (analog out)


void setup() {
  // initialize serial communications at 9600 bps:
  Serial.begin(9600);

}

void loop() {
  // put your main code here, to run repeatedly:
  readPot();
}











int readPot(void){
  // This function is from Arduino Library
  // http://www.arduino.cc/en/Tutorial/AnalogInOutSerial
  
  // read the analog in value:
  potValue = analogRead(analogInPin);
  // map it to the range of the analog out:
  outputPot = map(outputPot, 0, 1023, 0, 255);
  // change the analog out value:
  analogWrite(analogOutPin, outputPot);

  // print the results to the Serial Monitor:
  Serial.print("sensor = ");
  Serial.print(potValue);
  Serial.print("\t output = ");
  Serial.println(outputPot);

  // wait 2 milliseconds before the next loop for the analog-to-digital
  // converter to settle after the last reading:
  delay(500);
  return 0;
}

