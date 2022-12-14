package main

import (
	"fmt"
	"math"
	"os"
)

// ElectricalCircuit
// класс для задания электрической цепи
// из функции переменного напряжения, резистора,
// конденсатора, индуктивности
type ElectricalCircuit struct {
	PowerSource float64
	Resistor    float64
	Capacitor   float64
	Inductance  float64
}

// Input: ElectricalCircuit.go
// Ввод данных о цепи через консоль
func (circuit *ElectricalCircuit) Input() {
	fmt.Println("Введите данные о цепи:")
	fmt.Println("Источник питания:")
	fmt.Scan(&circuit.PowerSource)
	fmt.Println("Резистор:")
	fmt.Scan(&circuit.Resistor)
	fmt.Println("Конденсатор:")
	fmt.Scan(&circuit.Capacitor)
	fmt.Println("Индуктивность:")
	fmt.Scan(&circuit.Inductance)
}

// Output: ElectricalCircuit.go
// Вывод данных о цепи в консоль
func (circuit *ElectricalCircuit) Output() {
	fmt.Println("Источник питания:", circuit.PowerSource)
	fmt.Println("Резистор:", circuit.Resistor)
	fmt.Println("Конденсатор:", circuit.Capacitor)
	fmt.Println("Индуктивность:", circuit.Inductance)
}

// InputFromFile: ElectricalCircuit.go
// Ввод данных о цепи из файла
// Формат файла:
// <вольтаж>
// <сопротивление>
// <ёмкость>
// <индуктивность>
func (circuit *ElectricalCircuit) InputFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fscan(file, &circuit.PowerSource)
	fmt.Fscan(file, &circuit.Resistor)
	fmt.Fscan(file, &circuit.Capacitor)
	fmt.Fscan(file, &circuit.Inductance)
}

// OutputToFile: ElectricalCircuit.go
// Вывод данных о цепи в файл
// Формат файла:
// Источник питания: <значение>
// Резистор: <значение>
// Конденсатор: <значение>
// Индуктивность: <значение>
func (circuit *ElectricalCircuit) OutputToFileForHuman(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintln(file, "Источник питания:", circuit.PowerSource)
	fmt.Fprintln(file, "Резистор:", circuit.Resistor)
	fmt.Fprintln(file, "Конденсатор:", circuit.Capacitor)
	fmt.Fprintln(file, "Индуктивность:", circuit.Inductance)
}

// OutputToFile: ElectricalCircuit.go
// Вывод данных о цепи в файл
// Формат файла:
// <вольтаж>
// <сопротивление>
// <ёмкость>
// <индуктивность>
func (circuit *ElectricalCircuit) OutputToFileForProgram(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintln(file, circuit.PowerSource)
	fmt.Fprintln(file, circuit.Resistor)
	fmt.Fprintln(file, circuit.Capacitor)
	fmt.Fprintln(file, circuit.Inductance)
}

func (circuit *ElectricalCircuit) CalculateAmplitude() float64 {
	return circuit.PowerSource / circuit.Resistor
}

func (circuit *ElectricalCircuit) CalculatePhase() float64 {
	return circuit.Capacitor / circuit.Inductance
}

func (circuit *ElectricalCircuit) CalculateFrequency() float64 {
	return 1 / circuit.Inductance
}

// GenerateSignal: ElectricalCircuit.go
// Генерация сигнала длительностью time с шагом step
// f(t) = A * sin(2 * pi * freq * t + phi)
func (circuit *ElectricalCircuit) GenerateSignal(filename string, time float64, step float64) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	amplitude := circuit.CalculateAmplitude()
	phase := circuit.CalculatePhase()
	frequency := circuit.CalculateFrequency()
	for t := 0.0; t <= time; t += step {
		fmt.Fprintln(file, amplitude*math.Sin(frequency*t+phase))
	}
}

// GenerateSignalForVariableVoltage: ElectricalCircuit.go
// Генерация сигнала длительностью time с шагом step
// f(t) = A(t) * sin(2 * pi * freq * t + phi)
// Собственное значение напряжения игнорируется
func (circuit *ElectricalCircuit) GenerateSignalForVariableVoltage(filename string, time float64, step float64, voltage func(float64) float64) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	phase := circuit.CalculatePhase()
	frequency := circuit.CalculateFrequency()
	for t := 0.0; t <= time; t += step {
		amplitude := voltage(t) / circuit.Resistor
		fmt.Fprintln(file, amplitude*math.Sin(frequency*t+phase))
	}
}
