package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	_ "gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	_ "log"
	"math"
	"math/cmplx"
	"sync"
)

const MinFreq = 1e-3

// base_FFT - функция для вычисления БПФ
// Здесь реализован алгоритм БПФ с помощью рекурсии
// Если хотите применить БПФ к своему массиву, то вызывайте функцию FFT
// В base_FFT нет проверки на степень двойки и нормализации после вычисления
func base_FFT(x *[]complex128) {
	N := len(*x)
	if N <= 1 {
		return
	}

	// разбиваем на четные и нечетные
	even := make([]complex128, N/2)
	odd := make([]complex128, N/2)
	for i := 0; i < N/2; i++ {
		even[i] = (*x)[2*i]
		odd[i] = (*x)[2*i+1]
	}

	// рекурсивно вызываем FFT для четных и нечетных
	base_FFT(&even)
	base_FFT(&odd)

	// собираем результат
	for k := 0; k < N/2; k++ {
		t := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(N)) * odd[k]
		(*x)[k] = even[k] + t
		(*x)[k+N/2] = even[k] - t
	}
}

// FFT - функция для вычисления БПФ
// Осторожно! Меняет исходный массив
func FFT(x *[]complex128) {
	// дополняем массив до степени двойки
	n := 1
	for n <= len(*x) {
		n *= 2
	}

	for i := len(*x); i < n; i++ {
		*x = append(*x, complex(0, 0))
	}

	base_FFT(x)

	for i := 0; i < len(*x); i++ {
		if math.Abs(real((*x)[i])) < MinFreq {
			(*x)[i] = complex(0, imag((*x)[i]))
		}
		if math.Abs(imag((*x)[i])) < MinFreq {
			(*x)[i] = complex(real((*x)[i]), 0)
		}
	}
}

// base_IFFT - функция для вычисления обратного БПФ
// Здесь реализован алгоритм БПФ с помощью рекурсии
// Если хотите применить обратное БПФ к своему массиву, то вызывайте функцию IFFT
// В base_IFFT нет проверки на степень двойки и нормализации после вычисления
func base_IFFT(x *[]complex128) {
	N := len(*x)
	if N <= 1 {
		return
	}

	// разбиваем на четные и нечетные
	even := make([]complex128, N/2)
	odd := make([]complex128, N/2)
	for i := 0; i < N/2; i++ {
		even[i] = (*x)[2*i]
		odd[i] = (*x)[2*i+1]
	}

	// рекурсивно вызываем FFT для четных и нечетных
	base_IFFT(&even)
	base_IFFT(&odd)

	// собираем результат
	for k := 0; k < N/2; k++ {
		t := cmplx.Rect(1, 2*math.Pi*float64(k)/float64(N)) * odd[k]
		(*x)[k] = even[k] + t
		(*x)[k+N/2] = even[k] - t
	}
}

// IFFT - функция для вычисления обратного БПФ
// Осторожно! Меняет исходный массив
func IFFT(x *[]complex128) {
	// дополняем массив до степени двойки
	n := 1
	for n < len(*x) {
		n *= 2
	}

	for i := len(*x); i < n; i++ {
		*x = append(*x, complex(0, 0))
	}

	base_IFFT(x)
	N := len(*x)
	for i := 0; i < N; i++ {
		(*x)[i] /= complex(float64(N), 0)
	}

	for i := 0; i < len(*x); i++ {
		if math.Abs(real((*x)[i])) < MinFreq {
			(*x)[i] = complex(0, imag((*x)[i]))
		}
		if math.Abs(imag((*x)[i])) < MinFreq {
			(*x)[i] = complex(real((*x)[i]), 0)
		}
	}
}

// Multiply - функция для умножения двух полиномов
// Не меняет исходные массивы.
// Возвращает массив, содержащий результат умножения
func Multiply(A []complex128, B []complex128) []complex128 {
	if len(A) > len(B) {
		for i := len(B); i < len(A); i++ {
			B = append(B, 0)
		}
	}
	if len(A) < len(B) {
		for i := len(A); i < len(B); i++ {
			A = append(A, 0)
		}
	}

	// вычисляем БПФ для обоих массивов в две рутины и ждем их завершения
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		FFT(&A)
	}()
	go func() {
		defer wg.Done()
		FFT(&B)
	}()
	var complex_result []complex128

	for i := 0; i < len(A); i++ {
		complex_result = append(complex_result, A[i]*B[i])
	}

	IFFT(&complex_result)

	return complex_result
}

// PlotSignal - функция для построения графика сигнала
// time - время сигнала (в секундах)
func PlotSignal(signal []complex128, time float64, title string, filename string) {
	fmt.Println("Plotting " + title + "...")
	p := plot.New()

	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	x := make(plotter.XYs, len(signal))

	for i := 0; i < len(signal); i++ {
		// progress bar (100)
		if i%(len(signal)/100) == 0 {
			fmt.Print("#")
		}
		x[i].X = time / float64(len(signal)) * float64(i)
		x[i].Y = real(signal[i])
	}
	fmt.Println()

	err := plotutil.AddLinePoints(p, title, x)

	if err != nil {
		panic(err)
	}

	// сохраняем график в файл
	fmt.Println("Saving " + filename + "...")
	if err := p.Save(8*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
	fmt.Println(filename + " saved!")
}
