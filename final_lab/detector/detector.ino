int reading = 0;

const int pin = A0;
const int threshold = 69;

void setup() {
  Serial.begin(9600);
}

void loop() {
  reading = analogRead(pin);

  // If reading passes the threshold,
  // it is a hit. Sometimes duplicates
  // occur, but these are deduplicated
  // by the client because the Arduino
  // has limited resources.
  if (reading >= threshold) {
    Serial.println(micros()); // Write timestamp to serial port
  }

  delay(50); // Don't overwhelm the serial port buffer
}
