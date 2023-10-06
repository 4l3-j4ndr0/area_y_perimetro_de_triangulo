package main

import (
	"fmt"
	"math"
	"github.com/sqrthree/toFixed"
)
func main() {
	
	var (
		cateto1 float64
		cateto2 float64
	)
	const decimales = 2
	fmt.Println(`Ingrese la longitud de los dos catetos de un triangulo rectangulo.
	Estos valores deben ser positivos y enteros.`)
	fmt.Println("Cateto 1")
		fmt.Scan(&cateto1)
		if cateto1 > 0 {
			fmt.Println("Cateto 2")
			fmt.Scan(&cateto2)
		}
	
if cateto1 > 0 && cateto2 > 0 {
	cateto1Cuadrado := math.Pow(float64(cateto1),2)
	cateto2Cuadrado := math.Pow(float64(cateto2),2)
	hipotenusa := math.Sqrt(cateto1Cuadrado + cateto2Cuadrado) 
	fmt.Println("CATETO 1 AL CUADRADO :", cateto1Cuadrado)
	fmt.Println("CATETO 2 AL CUADRADO :", cateto2Cuadrado)
	fmt.Printf("HIPOTENUSA : %g \n", hipotenusa)
	fmt.Println("AREA :", (cateto1 * cateto2)/2 )
	fmt.Println("PERIMETRO :", (cateto1 + cateto2 + hipotenusa))
	fmt.Printf("AREA 2 DECIMALES : %.2f \n",((cateto1 * cateto2)/2)) 
	fmt.Printf("AREA 2 DECIMALES USANDO CONSTANTE : %.*f \n",decimales,((cateto1 * cateto2)/2)) 
	fmt.Printf("PERIMETRO 2 DECIMALES : %.2f \n", (cateto1 + cateto2 + hipotenusa))
	fmt.Printf("PERIMETRO 2 DECIMALES USANDO CONSTANTE : %.*f \n", decimales,(cateto1 + cateto2 + hipotenusa))
	fmt.Println("PERIMETRO 2 DECIMALES ( OTRA VIA : ToFixed() ) :", toFixed.ToFixed((cateto1 + cateto2 + hipotenusa),decimales))
}else {
	fmt.Println("Todos los datos deben ser enteros y mayor que cero")
}
	
}
