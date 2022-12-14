package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func GenerateSignal(frequency float64, amplitude float64, phase float64, time float64, step float64) []complex128 {
	var signal []complex128
	fmt.Println("Generating signal...")
	for i := 0; i < int(time/step); i++ {
		// progress bar (100 штук)
		if i%(int(time/step)/100) == 0 {
			fmt.Print("#")
		}
		// считаем сигнал
		signal = append(signal, complex(amplitude*math.Sin(2*math.Pi*frequency*float64(i)*step+phase), 0))
	}
	fmt.Println()
	return signal
}

func GenerateComplexSignal(frequencies []float64, amplitudes []float64, phases []float64, time float64, step float64) []complex128 {
	var signal []complex128
	fmt.Println("Generating complex signal...")
	for i := 0; i < int(time/step); i++ {
		// progress bar (100 штук)
		if i%(int(time/step)/100) == 0 {
			fmt.Print("#")
		}
		var value complex128
		for j := 0; j < len(frequencies); j++ {
			value += complex(amplitudes[j]*math.Sin(2*math.Pi*frequencies[j]*float64(i)*step+phases[j]), 0)
		}
		signal = append(signal, value)
	}
	fmt.Println()
	return signal
}

func ReadSignalFromFile(filename string) []complex128 {
	var signal []complex128
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return signal
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		var re, im float64
		fmt.Sscanf(text, "%f %f", &re, &im)
		signal = append(signal, complex(re, im))
	}
	return signal
}

func WriteSignalToFile(signal []complex128, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	for i := 0; i < len(signal); i++ {
		fmt.Fprintln(file, real(signal[i]), imag(signal[i]))
	}
}

// GenerateNoise генерирует шум с заданной амплитудой
func GenerateNoise(amplitude float64, time float64, step float64) []complex128 {
	var noise []complex128
	for i := 0; i < int(time/step); i++ {
		noise = append(noise, complex(amplitude*rand.NormFloat64(), 0))
	}
	return noise
}

// AddNoise служебная функция
// Не используйте ее если хотите добавить шум к сигналу
// Для этих целей вам нужен AddNoiseToSignal
// AddNoise складывает два сигнала
func AddNoise(signal []complex128, noise []complex128) []complex128 {
	var result []complex128
	for i := 0; i < len(signal); i++ {
		result = append(result, signal[i]+noise[i])
	}
	return result
}

// AddNoiseToSignal добавляет шум с заданной амплитудой к сигналу
func AddNoiseToSignal(signal []complex128, amplitude float64, time float64, step float64) []complex128 {
	return AddNoise(signal, GenerateNoise(amplitude, time, step))
}
