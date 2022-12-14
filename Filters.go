/*
	В данном файле находятся функции для работы с сигналами
	Все функции принимают на вход сигнал, время его длительности, шаг дискретизации
	и возвращают сигнал, время его длительности, шаг дискретизации

	Функция SeparateSignalByAmplitude работает быстрее, но намного менее точно
	Если вам нужна точность, то используйте функцию SeparateSignalByCount

	Обе функции могут лишь примерно определить амплитуду, фазы сдвинуты. Частоты точные,
	но для этого лучше подавать сигнал длительностью в 1 период и как можно точнее
	определить число частот
*/

package main

import (
	"fmt"
	"math/cmplx"
)

// SeparateSignalByAmplitude
//
//	min_amplitude - минимальная амплитуда, которая будет учитываться
//	лучше всего брать около 1 наименьшей амплитуды
//
// На данный момент time не используется
func SeparateSignalByAmplitude(signal []complex128, time float64, step float64, min_amp float64) (frequencies []float64, amplitudes []float64, phases []float64) {
	fmt.Println("Separating signal by amplitudes...")
	FFT(&signal)
	for i := 1; i < len(signal)/2; i++ {
		if (cmplx.Abs(signal[i]) > cmplx.Abs(signal[i-1])) && (cmplx.Abs(signal[i]) > cmplx.Abs(signal[i+1]) && (2*cmplx.Abs(signal[i])/float64(len(signal)) > min_amp)) {
			frequencies = append(frequencies, float64(i)/step/float64(len(signal)))
			amplitudes = append(amplitudes, 2*cmplx.Abs(signal[i])/float64(len(signal)))
			phases = append(phases, cmplx.Phase(signal[i]))
		}
	}
	return frequencies, amplitudes, phases
}

// SeparateSignalByCount
// Фильтр с отсечкой по количеству частот count
// На данный момент time не используется
func SeparateSignalByCount(signal []complex128, time float64, step float64, count int) (frequencies []float64, amplitudes []float64, phases []float64) {
	fmt.Println("Separating signal by count...")
	FFT(&signal)
	for min_amp := 0.01; ; min_amp += 0.01 {
		frequencies = nil
		amplitudes = nil
		phases = nil
		for i := 1; i < len(signal)/2; i++ {
			if (cmplx.Abs(signal[i]) > cmplx.Abs(signal[i-1])) && (cmplx.Abs(signal[i]) > cmplx.Abs(signal[i+1]) && (2*cmplx.Abs(signal[i])/float64(len(signal)) > min_amp)) {
				frequencies = append(frequencies, float64(i)/step/float64(len(signal)))
				amplitudes = append(amplitudes, 2*cmplx.Abs(signal[i])/float64(len(signal)))
				phases = append(phases, cmplx.Phase(signal[i]))
			}
		}
		if len(frequencies) <= count {
			break
		}
	}
	if len(frequencies) == 0 {
		fmt.Println("----------Can't separate signal by count----------")
	}
	return frequencies, amplitudes, phases
}

// Функция для добавления частот
func AddFrequenciesToSignal(signal []complex128, new_frequencies []float64, new_amplitudes []float64, new_phases []float64, time float64, step float64) []complex128 {
	if len(new_frequencies) != len(new_amplitudes) || len(new_frequencies) != len(new_phases) {
		panic("AddFrequency: len(new_frequencies) != len(new_amplitudes) || len(new_frequencies) != len(new_phases)")
	}
	fmt.Println("Adding frequencies...")
	f, a, p := SeparateSignalByAmplitude(signal, time, step, 0.5)
	for i := 0; i < len(new_frequencies); i++ {
		f = append(f, new_frequencies[i])
		a = append(a, new_amplitudes[i])
		p = append(p, new_phases[i])
	}
	return GenerateComplexSignal(f, a, p, time, step)
}
