# go-ecommerce-proyect

Este proyecto es una aplicación de comercio electrónico desarrollada en Go que utiliza una arquitectura hexagonal (también conocida como arquitectura de puertos y adaptadores). Proporciona una estructura modular y desacoplada que facilita la escalabilidad, el mantenimiento y la prueba de la aplicación.

## Arquitectura Hexagonal

La arquitectura hexagonal se basa en los principios de separación de preocupaciones y dependencias inversas. En esta arquitectura, la lógica de negocio central se encuentra en el núcleo de la aplicación y está aislada de las capas externas. Las interacciones con el mundo exterior, como la persistencia de datos, la interfaz de usuario y servicios externos, se implementan a través de puertos y adaptadores.

La estructura general de la arquitectura hexagonal incluye las siguientes capas:

- **Domain Layer**: Esta capa contiene la lógica de negocio central de la aplicación, incluyendo las entidades, los objetos de valor y los casos de uso. No tiene dependencias de ninguna otra capa y es independiente de la infraestructura o tecnología utilizada.

- **Application Layer**: Esta capa implementa los casos de uso y orquesta la lógica de negocio utilizando los servicios definidos en el dominio. También se encarga de la validación de entrada y la coordinación de las operaciones.

- **Infrastructure Layer**: Esta capa proporciona las implementaciones concretas de los puertos definidos en el dominio. Incluye la persistencia de datos, servicios externos, interfaces de usuario y cualquier otra infraestructura necesaria para ejecutar la aplicación.

- **Adapters**: Estos adaptadores conectan los componentes externos a los puertos definidos en el dominio. Se encargan de adaptar los datos y las operaciones según las necesidades de cada componente externo. Por ejemplo, pueden implementar la comunicación con una base de datos, una API externa o una interfaz de usuario específica.

## Estructura del proyecto

El proyecto está organizado en los siguientes directorios:

- `controllers`: Contiene los controladores que manejan las solicitudes HTTP y coordinan las operaciones en la capa de aplicación.

- `database`: Contiene la configuración y utilidades relacionadas con la base de datos de la aplicación.

- `middleware`: Contiene middleware que se utiliza para interceptar y procesar las solicitudes HTTP antes de que lleguen a los controladores.

- `models`: Contiene las definiciones de los modelos de datos utilizados en la aplicación.

- `routes`: Contiene las definiciones de las rutas de la API y su mapeo a los controladores correspondientes.

- `test`: Contiene pruebas unitarias para los diferentes componentes de la aplicación.

- `tokens`: Contiene la lógica relacionada con la generación y gestión de tokens utilizados para la autenticación y autorización.

Además de estos directorios, encontrarás otros archivos importantes como `main.go` que es el punto de entrada de la aplicación, `docker-compose.yaml` para facilitar la configuración del entorno de desarrollo y `README.md` para proporcionar información adicional sobre el proyecto.

## Requisitos y ejecución

Para ejecutar este proyecto, se requiere tener instalado Go en tu entorno de desarrollo. Puedes ejecutar la aplicación utilizando el comando `go run main.go`. Asegúrate de cumplir con los requisitos y seguir las instrucciones de configuración descritas en el proyecto.

## Contribución

Si deseas contribuir a este proyecto, puedes hacerlo abriendo un problema o enviando una solicitud de extracción. Agradecemos tus ideas, sugerencias y mejoras para hacer crecer esta aplicación de comercio electrónico basada en la arquitectura hexagonal.

¡Esperamos que disfrutes trabajando con go-ecommerce-proyect!

```bash
# You can start the project with below commands
docker-compose up -d
go run main.go
```

- **SIGNUP FUNCTION API CALL (POST REQUEST)**

http://localhost:8000/users/signup

```json
{
  "first_name": "Nombre",
  "last_name": "Apellido",
  "email": "correo@example.com",
  "password": "contraseña",
  "phone": "+1234567890"
}

```

Response :"Successfully Signed Up!!"

- **LOGIN FUNCTION API CALL (POST REQUEST)**

  http://localhost:8000/users/login

```json
{
  "email": "correo@example.com",
  "password": "contraseña"
}

```

response will be like this

```json
{
  "_id": "***********************",
  "first_name": "Nombre",
  "last_name": "Apellido",
  "password": "$2a$14$UIYjkTfnFnhg4qhIfhtYnuK9qsBQifPKgu/WPZAYBaaN17j0eTQZa",
  "email": "correo@example.com",
  "phone": "+1234567890",
  "token": "eyJc0Bwcm90b25vbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiam9zZXBoIiwiTGFzdF9OYW1lIjoiaGVybWlzIiwiVWlkIjoiNjE2MTRmNTM5ZjI5YmU5NDJiZDlkZjhlIiwiZXhwIjoxNjMzODUzNjUxfQ.NbcpVtPLJJqRF44OLwoanynoejsjdJb5_v2qB41SmB8",
  "Refresh_Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnLCJVaWQiOiIiLCJleHAiOjE2MzQzNzIwNTF9.ocpU8-0gCJsejmCeeEiL8DXhFcZsW7Z3OCN34HgIf2c",
  "created_at": "2022-04-09T08:14:11Z",
  "updtaed_at": "2022-04-09T08:14:11Z",
  "user_id": "61614f539f29be942bd9df8e",
  "usercart": [],
  "address": [],
  "orders": []
}
****
```

- **Admin add Product Function (POST REQUEST)**

  http://localhost:8000/admin/addproduct

```json
{
  "product_name": "Alienware x15",
  "price": 2500,
  "rating": 10,
  "image": "alienware.jpg"
}

```

Response : "Successfully added our Product Admin!!"

- **View all the Products in db GET REQUEST**

  http://localhost:8000/users/productview

Response

```json
[
  {
    "Product_ID": "6153ff8edef2c3c0a02ae39a",
    "product_name": "alienwarex15",
    "price": 1500,
    "rating": 10,
    "image": "alienware.jpg"
  },
  {
    "Product_ID": "616152679f29be942bd9df8f",
    "product_name": "giner ale",
    "price": 900,
    "rating": 5,
    "image": "gin.jpg"
  },
  {
    "Product_ID": "616152ee9f29be942bd9df90",
    "product_name": "iphone 13",
    "price": 1700,
    "rating": 4,
    "image": "ipho.jpg"
  },
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "whiskey",
    "price": 100,
    "rating": 7,
    "image": "whis.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "acer predator",
    "price": 3000,
    "rating": 10,
    "image": "acer.jpg"
  }
]

```

- **Search Product by regex function (GET REQUEST)**

defines the word search sorting
http://localhost:8000/users/search?name=al

response:

```json
[
  {
    "Product_ID": "616152fa9f29be942bd9df91",
    "product_name": "Alienware x15",
    "price": 1500,
    "rating": 10,
    "image": "1.jpg"
  },
  {
    "Product_ID": "616153039f29be942bd9df92",
    "product_name": "ginger Ale",
    "price": 300,
    "rating": 10,
    "image": "1.jpg"
  }
]
```

- **Adding the Products to the Cart (GET REQUEST)**

  http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

  Corresponding mongodb query

- **Removing Item From the Cart (GET REQUEST)**

  http://localhost:8000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx

- **Listing the item in the users cart (GET REQUEST) and total price**

  http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx

- **Addding the Address (POST REQUEST)**

  http://localhost:8000/addadress?id=user_id**\*\***\***\*\***

  The Address array is limited to two values home and work address more than two address is not acceptable

```json
{
  "house_name": "white house",
  "street_name": "white street",
  "city_name": "washington",
  "pin_code": "332423432"
}
```

- **Editing the Home Address(PUT REQUEST)**

  http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

- **Editing the Work Address(PUT REQUEST)**

  http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

- **Delete Addresses(GET REQUEST)**

  http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

  delete both addresses

- **Cart Checkout Function and placing the order(GET REQUEST)**

  After placing the order the items have to be deleted from cart functonality added

  http://localhost:8000?id=xxuser_idxxx

- **Instantly Buying the Products(GET REQUEST)**
  http://localhost:8000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
