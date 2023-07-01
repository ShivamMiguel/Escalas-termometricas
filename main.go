package main

import "fmt"

func kelvinToCelsius(kelvin float64) float64 {
    return kelvin - 273.15
}

func main() {
    pontoEbulicaoKelvin := 400.15
    pontoEbulicaoCelsius := kelvinToCelsius(pontoEbulicaoKelvin)

    fmt.Printf("O ponto de ebulição da água em Kelvin é %.2f\n", pontoEbulicaoKelvin)
    fmt.Printf("O ponto de ebulição da água em graus Celsius é %.2f\n", pontoEbulicaoCelsius)
}