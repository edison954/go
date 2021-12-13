# go

Nace para aprovechar los cores, nucleos

creado por:

Ken Thompson (B, C, Unix, UTF-8)
Rob Pike (UNIX, UTF-8)
Robert Griesemer (HotSpot, JVM)

golang.org

Playground

play.golang.org

pakages:

https://pkg.go.dev/std


Hola mundo en go

fmt.Println("Hola mundo")


```
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	nb, err := fmt.Println("Hello, Edison")
	//fmt.Println(nb, err)
	_, _ = fmt.Println(nb, err)
}
```

Operador declaracion corta

...Tipo

```
//...interface{}    (todos los tipos en Go implementan la interfaz vacia)

	fmt.Println("Hola", 42, true)


```


verbos


```
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var a int
var b string = "Programa"
var c bool

func main() {
	fmt.Println("Hello", a, b, c)
	fmt.Printf("El valor de la variable a es: %v\n", a)
	fmt.Printf("El valor de la variable a es: %v\n", b)
	fmt.Printf("El valor de la variable a es: %v\n", c)

	s1 := fmt.Sprint("El", b, "dice", " Hola mundo")
	fmt.Println(s1)
}
```


hola edison:

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Edison Arley")
}

```
go run main.go


go mux  --->> modulo para crear rutas

https://github.com/gorilla/mux

go get -u github.com/gorilla/mux

go mod init example/restapiconcept  


https://github.com/githubnemo/CompileDaemon   --> para mentener el servidor ejecutando

~/go/bin/CompileDaemon

~/go/bin/CompileDaemon -command="./restapiconcept"  			--> mantener la compilacionj

sudo nano /etc/paths.