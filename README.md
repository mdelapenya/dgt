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
  -P, --plate string   Plate to scrap. It will ignore the 'persist' flag
  -U, --until string   Plate where to scrap until (included)
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
Es posible ejecutar la herramienta como un stack de Docker Compose, incluyendo el servidor web con el API de matrículas así como una base de datos MySQL para la persistencia de los datos. Además, se levantarán tantos contenedores como letras llevamos en las matrículas (de la B a la M), que ejecutarán en su arranque la CLI para escanear todas las matrículas que comiencen por dicha letra, terminando en la última matrícula de la serie.

```shell

```shell
$ docker compose up --build
```

Accediendo a la base de datos es posible consultar las matrículas procesadas y sus distintivos:

```shell
$ docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt
# retornar todas las matrículas
mysql> select p.plate, s.description, s.emoji from plates p, stickers s where s.sticker_id=p.sticker_id;
# agregar matrículas por distintivo
mysql> select s.description, s.emoji, count(1) as count from plates p, stickers s where s.sticker_id=p.sticker_id group by s.sticker_id;
```

### Dump de la base de datos

Si necesitas extraer la base de datos, ejecuta el siguiente script:

```shell
./scripts/backup.sh
```

Un fichero SQL se creará en el directorio `database/sql` con el nombre `001_dgt.sql`, de modo que si recreas la ejecución desde cero, la base de datos aparecerá pre-cargada.

Importante decir que tendrás que actualizar las matrículas a utilizar en el fichero `docker-compose.yml`, para que no se solapen con las que ya están en la base de datos. Lee el siguiente apartado.

### Errores de memoria

Si por alguna razón se queda sin memoria el equipo en el que se ejecuta, puedes ejecutar el siguiente script:
  
```shell
./scripts/oom.sh
```

Generará un fichero `plates.txt` en el directorio actual, con la última matrícula procesada, que puedes usar para reanudar el proceso.

Básicamente, el script hace lo siguiente:

1. para todos los servicios menos la base de datos:

```shell 
$ docker stop dgt-go-j-1 dgt-go-l-1 dgt-go-f-1 dgt-go-h-1 dgt-go-c-1 dgt-go-k-1 dgt-go-m-1 dgt-go-b-1 dgt-go-g-1 dgt-go-d-1
```

2. detecta en qué punto se quedaron cada uno de los servicios:

```shell
# ejecuta una consulta en servicio de la base de datos, obteniendo el valor de la última matrícula procesada, en este caso para la letra M
$ docker exec -it dgt-db-1 mysql -u root -ppassw0rd --database=dgt -e "select plate from plates where plate_id=(select plate_id from plates where plate like '____M__' order by plate_id desc limit 1);" >> ./plates.txt
```

3. Coge esas matrículas y cambia el valor del flag `--from` para el servicio de cada letra en el fichero `docker-compose.yml` para que sea la siguiente a la última procesada.
4. Arranca los servicios de nuevo con `docker compose up`, para que coja los nuevos valores de matrícula desde los que empezar a procesar.

```shell
./scripts/restart.sh
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
