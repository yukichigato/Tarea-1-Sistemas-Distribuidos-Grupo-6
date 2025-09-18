***
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
    "books_id_list": [1, 4, 2, ...],
    "deposit": 100,
    "late_fee": 1 // 1: aplicar penalización, 0: no aplicar
}
```

**Response:**
```json
{
    "message": "saldo actualizado con exito"
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
_Cantidad de veces que se adquiere el libro_
```json
{
    "quantity": 3
}
```

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
    "status": "Disponible",
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
_No requiere body_

**Response:**
```json
{
    "message": "estado del prestamo actualizado con exito"
}
```

### Endpoint `/loans`
**Método:** POST

**Body request:**
```json
{
    "user_id": 1,
    "book_id": 2,
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
```json
{
    "user_id": 3,
    "book_id": 6,
    "sale_date": "2025-09-18T12:34:56Z"
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
    "book_id": 2,
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
    "password": "contraseña",
}
```

**Response:**
```json
{
    "message": "login exitoso"
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
_No requiere body_

**Response:**   
_Descuenta en 1 el stock del libro de id entregada_
```json
{
    "message": "inventario actualizado correctamente"
}
```