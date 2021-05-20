package goarea

import "math"

//Pi é um proporção numérica definida pela relação entre
//o perímetro de um circunferência e seu diâmetro
const Pi = 3.1416

func Circ(raid float64) float64 {
	return math.Pow(raid, 2) * Pi
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
