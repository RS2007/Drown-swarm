#include <WiFi.h>

const char *ssid = "SSID-of-wifi(say DLink-DIR-600M)";
const char *password = "Password-of-wifi";

void setup() {
  Serial.begin(9600);
  while (!Serial) {
    Serial.println("Waiting for serial connection to the PC to be made");
  }

  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print("."); // Whilst connecting
  }

  Serial.println("");
  Serial.println("Wi-Fi connected");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  if (WiFi.status() == WL_CONNECTED) {
    WiFiClient client;

    if (client.connect("<ip-address-of-machine-running-server>", 1234)) {
      Serial.println("Connected to server");

      while (client.connected()) {
        if (client.available()) {
          char c = client.read();
          if (c == NULL) {
            Serial.print("No message");
          } else {
            Serial.print(c);
          }
        }
      }
      client.stop();
      Serial.println();
      Serial.println("Server disconnected");
    } else {
      Serial.println("Failed to connect");
    }
    delay(5000); // Delay before making the next connection attempt
  }
}
