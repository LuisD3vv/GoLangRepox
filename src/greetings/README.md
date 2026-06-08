# Saludos en go

Este paquete proporciona una simple manera de obtener saludos personalizados en go

## Instalacion

Ejecutar el siguiente comando para instalar el paquete:

```bash
go get -u https://github.com/LuisD3vv/GoLangRepox/blob/main/src/greetings/greetings.go
```

## Uso

Aqui hay un ejemplo de uso

```go
package main

import (
    "fmt"
    "https://github.com/LuisD3vv/GoLangRepox/blob/main/src/greetings/greetings.go"
)

func main() {

    // Asi pasamos el parametro al modulo Greetings
    message,err := greeting.Hello("Luis")
    
    // Manejamos errores
    if err != nil {
        fmt.Println("Ocurrio un error",err)
        return
    }
    // imprimimos el mensaje
    fmt.Println(message)

}
```

### De esta forma podemos crear nuestro modulos propios e importarlos, descargarlos y usarlos de manera local como los ya incluidos en el lenguaje
