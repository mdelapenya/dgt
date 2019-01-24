# DGT Distintivos Ambientales de las Matrículas Españolas

Esta pequeña aplicación de ejemplo consulta el servicio web, ejem, de la DGT que permite obtener el tipo de distintivo ambiental asociado a una matrícula.

## Plates API

### Get sticker for a plate

```http
GET /plates/:plate
```

##### Request

```bash
curl "http://localhost:8080/plates/0000bbb"
curl "https://dgt.wedeploy.io/plates/0000bbb"
```

##### Response `200 OK`

```json
{"result":"Etiqueta Ambiental C"}
```

> Escribe la matrícula sin guiones ni espacios (0000XXX / XX0000XX / C0000XXX)