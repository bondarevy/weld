package main

import (
	"machine"
	"time"
)

// Инициализируем светодиод (D13)
var led = machine.LED

// Инициализируем пин для кнопки c прерыванием (D2)!
var button = machine.D3

//переменная нажата ли кнопка?
var buttonPressed bool = false

// Выбираем пины которые использует сварка для подачи мощности объединяем их в общий массив (слайс) с названием "powerPins"
var powerPins = []machine.Pin{
	machine.D2,
	machine.D4,
	machine.D5,
	machine.D6,
	machine.D7,
	machine.D8,
	machine.D9,
	machine.D10,
}

func buttonInterruptCallback() {

	//инвертируем значение переменное - если true - делаем false и наоборот
	if buttonPressed {
		buttonPressed = false
	}
}

// В функции init() обычно задаются все начальные параметры переменных - она выполнясется до функции main().
func init() {

	//настраиваем светодиод:
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Настраиваем пин для кнопки как вход и включаем внутреннюю подтяжку к питанию
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	//button.Configure(machine.PinConfig{Mode: buttonMode})

	// Настраиваем прерывание на пине кнопки
	//button.SetInterrupt(buttonPinChange, buttonInterruptCallback)

	// Конфигурирум все пины из слайса "powerPins" как исходящие "output" и подаем на них ток "High" по умолчанию (так устроена конструкция сварки).
	for _, powerPin := range powerPins {
		powerPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
		powerPin.High()
	}

}

// Эта функция задает мощность с которой варит сварка от 1 до 8 (количество транзисторов) - weldWithPowerLevel(1) = минимальная мощность, weldWithPowerLevel(8) = максимальная мощность.
func weldWithPowerLevel(powerLevel int) {

	// Проверяем что переданное значение powerLevel не превышает количество пинов (в нашем случае 8шт) и не меньше 0 .
	if powerLevel > len(powerPins) && powerLevel >= 0 {
		//печатаем в консоль ошибку если выбрано недопустимое количество мощности:
		println("Error: powerLevel (%d) is higher than ammount of transistors (%d).", powerLevel, len(powerPins))
	} else {
		// Eсли все в порядке (мощность равна или меньше количества транзисторов) переходим к следующему циклу выполнения:

		// Включаем сварку на указанную мощность (задействуем только необходимые транзисторы)
		for pinIndex, pin := range powerPins {
			if pinIndex < powerLevel {
				//включем мощность для выбраных пинов (пускаем мимо транзистора)
				pin.Low()
			} else {
				//отключаем мощность (пускаем через транзистор на резисторы)
				pin.High()
			}
		}

		// ждем 1 секунду (1000милисекунд):
		time.Sleep(1000 * time.Millisecond)

		// отключаем сварку
		for _, pin := range powerPins {
			//отключаем мощность (пускаем через транзистор на резисторы)
			pin.High()
		}

		//конец проверки powerLevel
	}

	//конец функции
}

func main() {

	// Запускаем бесконечный цикл
	for {
		
		// Проверяем если нажата кнопка:
		if button.Get() {
			//кнопка нажата
			//выполним сварку с мощностью = 2 пропущенными транзисторами:
			weldWithPowerLevel(2)

			//включаем светодиод для визуальной проверки
			led.High()
		} else {
			//кнопка не нажата
			//отключаем светодиод
			led.Low()
		}

	}
}
