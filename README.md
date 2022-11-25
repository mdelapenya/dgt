# DGT Distintivos Ambientales de las Matrículas Españolas

Esta pequeña aplicación de ejemplo consulta el servicio web, ejem, de la DGT que permite obtener el tipo de distintivo ambiental asociado a una matrícula.

Es posible ejecutarlo en local en dos modos:
1. Como línea de comandos.
2. Levantando un servidor web, exponiendo ciertos API endpoints.

## Plates CLI

La CLI ofrece el comando `scrap`, incluyendo la posibilidad de personalizar la respuesta con algunos flags:

```shell
Scraps all car plates retrieving their ECO sticker, starting in 0000BBB

Usage:
  dgt scrap [flags]

Flags:
  -F, --from string    Plate where to scrap from
  -h, --help           help for scrap
  -p, --persist        If the result will be persisted in a data store
  -P, --plate string   Plate to scrap
```

### Flags

| Flag      | Short version | Type    | Default | Requerido | Descripción                                                 |
| --------- | --------------| ------- | --------| --------- | ----------------------------------------------------------- |
| --from    | -F            | string  |         | no        | Una matrícula válida desde la que empezar el procesado      |
| --until   | -U            | string  |         | no        | Una matrícula válida en la que terminar el procesado      |
| --persist | -p            | boolean | no      | no        | Si es necesario persistir el resultado en un almacenamiento |
| --plate   | -P            | string  |         | no        | Si tiene valor, únicamente se procesará esa matrícula       |

## Docker
Es posible ejecutar la herramienta como una imagen Docker, previa construcción de la misma:

```shell
$ docker build -t mdelapenya/dgt:latest .
```

#### Comprobar una matrícula
```shell
$ docker run --rm  mdelapenya/dgt:latest scrap --plate 9334LSL
```

#### Escanear todas las mátriculas
```shell
$ docker run --rm  mdelapenya/dgt:latest scrap
```

#### Escanear todas las mátriculas desde una dada
Ésto es útil para saber por qué matrícula vamos:
```shell
$ docker run --rm  mdelapenya/dgt:latest scrap --from 9334LSL
```

#### Escanear todas las mátriculas desde una hasta otra
Ésto es útil para saber por qué matrícula vamos:
```shell
$ docker run --rm  mdelapenya/dgt:latest scrap --from 0000LSL --until 1000LSL
```

## Docker Compose
Es posible ejecutar la herramienta como un stack de Docker Compose, incluyendo el servidor web con el API de matrículas así como una base de datos MySQL para la persistencia de los datos.

```shell
$ docker compose up --build
```

De esta manera es posible levantar el stack, y utilizar además la CLI para consultar las matrículas, aunque para ello es necesario indicar la contraseña y la ubicación de la base de datos:

```shell
MYSQL_ROOT_PASSWORD=passw0rd MYSQL_SERVER=localhost go run main.go scrap -p
```

## Plates API

### Get sticker for a plate

```http
GET /plates/:plate
```

##### Request

```bash
curl "http://localhost:8080/plates/0000bbb"
```

##### Response `200 OK`

```json
{"result":"Etiqueta Ambiental C"}
```

> Escribe la matrícula sin guiones ni espacios (0000XXX)
