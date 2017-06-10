#include "ESP8266WiFi.h";
#include <DHT.h>;


// WiFi parameters
const char* ssid = "l4rs.net";
const char* password = "hangover1993ismykey";

const char* host = "192.168.55.21";
const int httpPort = 8080;

#define DHTPIN 5 // == PIN D1 on NodeMCU EP8266
#define DHTTYPE DHT22   // DHT 22  (AM2302)


 //// Initialize DHT sensor
DHT dht(DHTPIN, DHTTYPE);

float hum;
float temp;

void setup()
{
    Serial.begin(115200);
    dht.begin();

    Serial.println();
    Serial.println();
    Serial.print("Connecting to ");
    Serial.println(ssid);
    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
      Serial.print(".");
    }

    Serial.println("");
    Serial.println("WiFi connected");
    Serial.println("IP address: ");
    Serial.println(WiFi.localIP());
}

void loop()
{

      Serial.print("Connecting to [host] / [port] ");
      Serial.println(host);
      Serial.println(httpPort);
  
      WiFiClient client;
      if (!client.connect(host, httpPort)) {
        Serial.println("connection failed");
        return;
     }
  
    hum = dht.readHumidity();
    temp= dht.readTemperature();

    Serial.print("temp: ");
    Serial.println(temp);

    Serial.print("hum: ");
    Serial.println(hum);

    Serial.println("sending data");     
    client.print(String("PUT /data/test123/" + String(temp) + "/" + String(hum) + " HTTP/1.1\r\n" + "Host: " + host + "\r\n" + "Connection: close\r\n\r\n"));
    Serial.println("done sending... waiting for answer");
    delay(10);
    
    while(client.available()){
      String line = client.readStringUntil('\r');
      Serial.print(line);
    }
  
    Serial.println();
    Serial.println("closing connection");

    // 1000 milliseconds == 1 second
    delay(3000); //15min == 900000
}
