# Tarea-1-Sistemas-Distribuidos-Grupo-6
### Integrantes
- Andrés Araya Díaz | 202287004-4
- Sebastian Lopez | 202173020-6
- Baltazar Portilla 202173112-1

### Explicaciones sobre el uso de la IA en el proyecto
- IA fue utilizada para reorganizar el cliente desde un solo archivo `main.go` a varios por razones de mantenibilidad.
- IA fue utilizado para la creación de comentarios genéricos en las funciones que solo se encargan de hacer requests a la API.

### Justificaciones de endpoints extra
- La creación del endpoint `POST /login` fue preferido en lugar de utilizar `GET /users?email=xxxx&password=xxxx` por razones de diseño de APIs REST, semánticas de HTTP, seguridad, no utilizar malas prácticas y por razones prácticas.
- La creación del endpoint `GET /inventory/id` y `PATCH /inventory/id` fueron preferidas para un manejo más cómodo del inventario y no tener que hacer maniobras extras para la actualización de este.

### Ejecucion 
-vm2: 
- Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api/db/poblar$
- go run poblar.go (revisar que no exista mi_base.db en la carpeta)
-cd Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm2/api
-go run main.go (revisar que este disponible el puerto :8080)
- vm1:
- Tarea-1-Sistemas-Distribuidos-Grupo-6/tarea/vm1/client
- go run . 
## Uso del cliente

Entrar a la carpeta del cliente, en `Desktop` y ejecutar el siguiente comando para levantar el cliente y utilizarlo.
```sh
go run .
```

## Uso de la API

### Endpoint `/users`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
  "users": [
    {
        "id": 1,
        "first_name": "Baltazar",
        "last_name": "Portilla",
        "email": "baltazar.portilla@email.com",
        "password": "",
        "usm_pesos": -5
    },
    {
        "id": 2,
        "first_name": "Anakin",
        "last_name": "Skywalker",
        "email": "notvader@email.com",
        "password": "",
        "usm_pesos": 700
    },
    ...
  ]
}
```

### Endpoint `/users/id`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
    "id": 1,
    "first_name": "Baltazar",
    "last_name": "Portilla",
    "email": "baltazar.portilla@email.com",
    "password": "",
    "usm_pesos": -5
}
```

### Endpoint `/users/id`
**Método:** PATCH

**Body request:**   
_Campos opcionales, enviar solo lo que se desea actualizar_
```json
{
    "first_name": "Darth",
    "last_name": "Vader",
    "email": "vader@email.com",
    "usm_pesos": 10000
}
```

**Response:**
```json
{
    "message": "usuario actualizado con exito"
}
```

### Endpoint `/users`
**Método:** POST

**Body request:**
```json
{
    "first_name": "Nombre",
    "last_name": "Apellido",
    "email": "correo@email.com",
    "password": "contraseña"
}
```

**Response:**
```json
{
    "message": "usuario registrado con exito"
}
```

***

### Endpoint `/books`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
  "books": [
    {
      "id": 1,
      "book_name": "El principito",
      "book_category": "Infantil",
      "transaction_type": "Venta",
      "price": 10,
      "status": "Disponible",
      "popularity_score": 4
    },
    {
        "id": 4,
        "book_name": "Don Quijote de la Mancha",
        "book_category": "Novela",
        "transaction_type": "Arriendo",
        "price": 3,
        "status": "Disponible",
        "popularity_score": 12
    },
    ...
  ]
}
```

### Endpoint `/books/id`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
    "id": 1,
    "book_name": "El principito",
    "book_category": "Infantil",
    "transaction_type": "Venta",
    "price": 10,
    "status": "Disponible",
    "popularity_score": 4
}
```

### Endpoint `/books/id`
**Método:** PATCH

**Body request:**   
_Campos opcionales, enviar solo lo que se desea actualizar_
```json
{
    "book_name": "Nombre",
    "book_category": "Categoria",
    "transaction_type": "Venta/Arriendo",
    "price": 300,
    "status": "Disponible",
    "popularity_score":1
}

**Response:**
```json
{
    "message": "libro actualizado con exito"
}
```

### Endpoint `/books`
**Método:** POST

**Body request:**
```json
{
    "book_name": "Nombre",
    "book_category": "Categoria",
    "transaction_type": "Venta/Arriendo",
    "price": 4,
    "status": "Disponible"
}
```

**Response:**
```json
{
    "message": "libro registrado con exito"
}
```

***

### Endpoint `/loans`
**Método:** GET

**Body request:**  
_No requiere body_

**Query params opcionales:**    
* _id → Filtra préstamos por user_id. Si se omite, devuelve todos los préstamos_

**Response:**
```json
{
  "loans": [
    {
        "id": 1,
        "user_id": 1,
        "book_id": 2,
        "start_date": "2025-09-17T12:34:56Z",
        "return_date": "2025-09-24T12:34:56Z",
        "status": "Pendiente"
    },
    {
        "id": 2,
        "user_id": 2,
        "book_id": 3,
        "start_date": "2025-09-10T12:34:56Z",
        "return_date": "2025-09-17T12:34:56Z",
        "status": "Finalizado"
    },
    ...
  ]
}
```

### Endpoint `/loans/id`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
    "id": 1,
    "user_id": 1,
    "book_id": 2,
    "start_date": "2025-09-17T12:34:56Z",
    "return_date": "2025-09-24T12:34:56Z",
    "status": "Pendiente"
}
```

### Endpoint `/loans/id`
**Método:** PATCH

**Body request:**   
_Campos opcionales, enviar solo lo que se desea actualizar_
```json
{
    "user_id": 3,
    "book_id": 5,
    "start_date": "2025-09-20T12:34:56Z",
    "return_date": "2025-09-27T12:34:56Z",
    "status": "Finalizado"
}
```

**Response:**
```json
{
    "message": "prestamo actualizado con exito"
}
```

### Endpoint `/loans`
**Método:** POST

**Body request:**
```json
{
    "user_id": 1,
    "book_id": 2
}
```

**Response:**
```json
{
    "message": "prestamo registrado con exito"
}
```

***

### Endpoint `/sales`
**Método:** GET

**Body request:**  
_No requiere body_

**Query params opcionales:**    
* _id → Filtra ventas por user_id. Si se omite, devuelve todas las ventas

**Response:**
```json
{
  "sales": [
    {
        "id": 1,
        "user_id": 1,
        "book_id": 2,
        "sale_date": "2025-08-28T12:34:56Z"

    },
    {
        "id": 2,
        "user_id": 2,
        "book_id": 3,
        "sale_date": "2025-09-09T12:34:56Z"
    },
    ...
  ]
}
```

### Endpoint `/sales/id`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
    "id": 1,
    "user_id": 1,
    "book_id": 2,
    "sale_date": "2025-08-28T12:34:56Z"
}
```

### Endpoint `/sales/id`
**Método:** PATCH

**Body request:**   
_Campos opcionales, enviar solo lo que se desea actualizar_
```json
{
    "user_id": 2,
    "book_id": 6,
    "sale_date": "2025-08-27T12:34:56Z"
}
```

**Response:**
```json
{
    "message": "venta actualizada con exito"
}
```

### Endpoint `/sales`
**Método:** POST

**Body request:**
```json
{
    "user_id": 1,
    "book_id": 2
}
```

**Response:**
```json
{
    "message": "venta registrada con exito"
}
```

***
### Endpoint `/transactions`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**
```json
{
  "transactions": [
    {
        "transaction_id": 1,
        "book_id": 1,
        "book_name": "El principito",
        "transaction_type": "Venta",
        "transaction_date": "2025-09-17T12:34:56Z",
        "book_price": 4

    },
    {
        "transaction_id": 1,
        "book_id": 4,
        "book_name": "Don Quijote de la Mancha",
        "transaction_type": "Arriendo",
        "transaction_date": "2025-09-03T12:34:56Z",
        "book_price": 3
    },
    ...
  ]
}
```

***

### Endpoint `/login`
**Método:** POST

**Body request:**
```json
{
    "email": "correo@email.com",
    "password": "contraseña"
}
```

**Response:**
```json
{
    "id": 1,
    "first_name": "Cristobal",
    "last_name": "Alvarez",
    "email": "correo@email.com",
    "usm_pesos": 53
}
```

***

### Endpoint `/inventory/id`
**Método:** GET

**Body request:**  
_No requiere body_

**Response:**   
_Devuelve el stock del libro de id entregada_
```json
{
    "available_quantity": 3
}
```

### Endpoint `/inventory/id`
**Método:** PATCH

**Body request:**   
_Campos opcionales, enviar solo lo que se desea actualizar_
```json
{
    "available_quantity": 3
}
```

**Response:**
```json
{
    "message": "inventario actualizada con exito"
}
```
