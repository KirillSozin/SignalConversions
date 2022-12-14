package main

import (
	"fmt"
	"math"
)

func main() {
	time := 0.2
	step := time / (math.Pow(2.0, 14.0))
	// с помощью функции GenerateComplexSignal генерируем сложный сигнал
	// создаем массив частот, амплитуд и фаз
	frequencies := []float64{5, 20}
	amplitudes := []float64{5, 5}
	phases := []float64{-1, 1}

	// генерируем сложный сигнал
	complexSignal := GenerateComplexSignal(frequencies, amplitudes, phases, time, step)

	// строим график сигнала
	PlotSignal(complexSignal, time, "сomplex signal", "complexSignal.png")

	// добавляем шум
	noisySignal := AddNoiseToSignal(complexSignal, 1, time, step)
	// строим график шума
	PlotSignal(noisySignal, time, "noise", "noise.png")

	// фильтруем шум
	f, a, p := SeparateSignalByCount(noisySignal, time, step, 3)

	fmt.Println("frequencies:", f)
	// строим график фильтрации
	PlotSignal(GenerateComplexSignal(f, a, p, time, step), time, "filtered signal", "filteredSignal.png")
}
