#Append File Golang

> Funcion que añade contenido a un archivo y en caso no exista el archivo lo crea y añade el contenido.

- instalación

```shell
go get github.com/Leonardo-Antonio/append-file
```

- ejemplo
```go
err := AppendFile("log.txt", "[GET][201] | localhost:8080/api/v1/users")
if err != nil {
	panic(err)
}
```