# Simple CRUD

Everty thing start from simple thing. this project is Restful API for CRUD operation with Echo as web framework

## Running 

To run, run the following command

```bash
  go run main.go
```

## Endpoint 

```http
  POST /contents
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`    | `json` | **Required**.     |
| `body`    | `json` | **Required**. 

```http
  GET /contents
```

```http
  GET /contents/:id
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`    | `param` | **Required**.     |


```http
  PUT /contents/:id
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`    | `param` | **Required**.     |


```http
  DELETE /contents/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`    | `param` | **Required**.     |


