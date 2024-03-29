# workshop rgb led

This workshop cover the usage of The KY-016 rgb full color led module with the Raspberry Pi.  

# Wiring

![rgb](doc/img/rgb.jpg)

|Sensor             | Raspberry Pi                 |
|-------------------|---------------------------------|
|-                  | GND pin                         |
|Blue               | GPIO 23                         |
|Green              | GPIO 24                         |
|Red                | GPIO 25                         |

![wiring](doc/img/gpio.png)

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to Raspberry Pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download relay binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_rgb_led/master/bin/arm6/rgb
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_rgb_led/master/bin/arm7/rgb
```
## make it executable
```
chmod +x rgb
```

## 4- execute binary
```
./rgb
```

You can call one of those color **green**, **red**, **orange** to be written in lower case

## How to stop the program
To quit or stop the program click on **Ctrl+C**

# Ressources
https://www.piddlerintheroot.com/rgb-moisture-sensor/

