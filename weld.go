package main

import (
    "machine"
    "time"
)

func main() {
    // Назначение пинов на выход
    pins := []machine.Pin{
        machine.D11,
        machine.D10,
        machine.D9,
        machine.D8,
        machine.D7,
        machine.D6,
        machine.D5,
        machine.D4,
        machine.led
    }
    
 

    for {

for _, pin := range pins {
        pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
        // Установка пина в LOW
        pin.High()
    }

    // Задержка в 1 секунду
    time.Sleep(time.Second)


for _, pin := range pins {
        pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
        // Установка пина в LOW
        pin.Low()
    }
    }
}
