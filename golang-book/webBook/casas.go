package main
import "fmt"

//Define a new type
type hogar struct {
  tipo string
  direccion string
  tamano int
}

//Compare the type of home that 2 poeple had
//Then return the biggest, the difference and the type of the bigger home

func Bigger (h1, h2 hogar) (hogar, int) {
  if h1.tamano > h2.tamano {
    return h1, h1.tamano - h2.tamano
  }
  return h2, h2.tamano - h1.tamano
}


//Function to capture data
func (h *hogar) capturarDatos (home hogar) (hogar)   {

  var tipo, direccion string
  var tamano int

  fmt.Println("Ingrese el tipo de la casa\n")
  fmt.Scanf("%v", &tipo)
  fmt.Println("Ingrese la direcion de la casa\n")
  fmt.Scanf("%v", &direccion)
  fmt.Println("Ingrese el tamaño de la casa\n")
  fmt.Scanf("%v", &tamano)
/* Another way to initialize the struct
-------------------------------------------
  home.tipo = tipo
  home.direccion = direccion
  home.tamaño = tamano
------------------------------------------
*/
  home = hogar{tipo:tipo, direccion:direccion, tamano:tamano}

  return home
  }

func main () {

var h1, h2 hogar

//Capture the data of the first house
  fmt.Println("Ingrese los datos de la primera casa a comparar\n")
    //After the capture rewrite te global variables
  //And set this as the data of the first house
  h1.capturarDatos(h1)

  fmt.Println("Ingrese los datos de la segunda casa a comparar\n")
  h2.capturarDatos(h2)

  biggerHome, biggerHomeDiff := Bigger(h1, h2)
  fmt.Println("La casa mas grande, es la casa localizada en la direcion \n", biggerHome.direccion, "\n", "Y del tipo \n",  biggerHome.tipo)
  fmt.Println("\nY ademas ganó por una diferencia de \n", biggerHomeDiff)
}
