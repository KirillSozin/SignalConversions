package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

// максимальная ошибка
const max_mistake = 0.05

func BenchmarkMultiply(b *testing.B) {
	// создаем и заполняем массив случайными числами
	var A []complex128
	for i := 0; i < 2000000; i++ {
		// случайные числа от 0 до 100
		A = append(A, complex(rand.Float64()*100, rand.Float64()*100))
	}
	var B []complex128
	for i := 0; i < 2000000; i++ {
		// случайные числа от 0 до 100
		B = append(B, complex(rand.Float64()*100, rand.Float64()*100))
	}
	for i := 0; i < b.N; i++ {
		Multiply(A, B)
	}
}

func BenchmarkSeparateSignal(b *testing.B) {
	// создаем случайный сигнал из не более b.N частот
	var freqs []float64
	var amps []float64
	var phases []float64
	max_freq := 0.1
	min_amp := 0.1
	for i := 0; i < b.N; i++ {
		// стараемся чтобы частоты не были слишком близко друг к другу
		freqs = append(freqs, rand.Float64()*100)
		amps = append(amps, rand.Float64()*100)
		phases = append(phases, rand.Float64()*2*math.Pi)

		if freqs[i] < max_freq {
			max_freq = freqs[i]
		}
		if amps[i] < min_amp {
			min_amp = amps[i]
		}
	}
	time := 1.0 / max_freq
	step := time / math.Pow(2, 10)
	signal := GenerateComplexSignal(freqs, amps, phases, time, step)
	f, a, p := SeparateSignalByAmplitude(signal, time, step, 0.6*min_amp)
	// проверяем насколько сигналы совпадают
	freq_err := 0
	amp_err := 0
	phase_err := 0
	for i := 0; i < len(freqs); i++ {
		// ищем есть ли такая же частота в полученном сигнале
		ok := -1
		for j := 0; j < len(f); j++ {
			if math.Abs(freqs[i]-f[j]) < math.Abs(freqs[i])*max_mistake {
				ok = j
				break
			}
		}
		if ok == -1 {
			freq_err++
			amp_err++
			phase_err++
			continue
		} else {
			if math.Abs(amps[i]-a[ok]) > math.Abs(amps[i])*max_mistake {
				amp_err++
			}
			if math.Abs(phases[i]-p[ok]) > math.Abs(phases[i])*max_mistake {
				phase_err++
			}
		}
	}
	fmt.Println("Count of frequencies: ", len(freqs))
	fmt.Println("Percentage of frequency errors:", float64(freq_err)/float64(len(freqs))*100, "%")
	fmt.Println("Percentage of amplitude errors:", float64(amp_err)/float64(len(freqs))*100, "%")
	fmt.Println("Percentage of phase errors:", float64(phase_err)/float64(len(freqs))*100, "%")
}

func BenchmarkSeparateSignalByCount(b *testing.B) {
	// создаем случайный сигнал из не более b.N частот
	var freqs []float64
	var amps []float64
	var phases []float64
	max_freq := 0.1
	for i := 0; i < b.N; i++ {
		// стараемся чтобы частоты не были слишком близко друг к другу
		freqs = append(freqs, rand.Float64()*100)
		amps = append(amps, rand.Float64()*100)
		phases = append(phases, rand.Float64()*2*math.Pi)

		if freqs[i] < max_freq {
			max_freq = freqs[i]
		}
	}
	time := 1.0 / max_freq
	step := time / math.Pow(2, 14)
	signal := GenerateComplexSignal(freqs, amps, phases, time, step)
	f, a, p := SeparateSignalByCount(signal, time, step, len(freqs))
	// проверяем насколько сигналы совпадают
	freq_err := 0
	amp_err := 0
	phase_err := 0
	for i := 0; i < len(freqs); i++ {
		// ищем есть ли такая же частота в полученном сигнале
		ok := -1
		for j := 0; j < len(f); j++ {
			if math.Abs(freqs[i]-f[j]) < math.Abs(freqs[i])*max_mistake {
				ok = j
				break
			}
		}
		if ok == -1 {
			freq_err++
			amp_err++
			phase_err++
			continue
		} else {
			if math.Abs(amps[i]-a[ok]) > math.Abs(amps[i])*max_mistake {
				amp_err++
			}
			if math.Abs(phases[i]-p[ok]) > math.Abs(phases[i])*max_mistake {
				phase_err++
			}
		}
	}
	fmt.Println("Count of frequencies: ", len(freqs))
	fmt.Println("Percentage of frequency errors:", float64(freq_err)/float64(len(freqs))*100, "%")
	fmt.Println("Percentage of amplitude errors:", float64(amp_err)/float64(len(freqs))*100, "%")
	fmt.Println("Percentage of phase errors:", float64(phase_err)/float64(len(freqs))*100, "%")
}
