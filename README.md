# Test MS

El siguiente repositorio es parte de una prueba técnica, se cuentan con 2 carpetas en las cuales esta debidamente implementado cada servicio.

beer-api es una API Rest que puede ser consumida por cualquier usuario.
currency es un microservicio que se alimenta de fuente de datos externas y almacena cierta información que obtiene de esta misma.


## Problema

Bender es fanÃ¡tico de las cervezas y quiere tener un registro de todas las cervezas que prueba y como calcular el precio que necesita para comprar una caja de algÃºn tipo especifico de cervezas. Para esto necesita una API REST con esta informaciÃ³n que posteriormente compartirÃ¡ con sus amigos.

### DescripciÃ³n

Se solicita crear un API REST basÃ¡ndonos en la definiciÃ³n que se encuentra en el archivo **openapi.yaml**.

#### Funcionalidad

- GET /Beers: Lista todas las cervezas que se encuentran en el sistema.
- POST /Beers: Permite ingresar una nueva cerveza.
- GET /beers/{beerID}: Lista un detalle de una cerveza especifica.
- GET /beets/{beerID}/boxprice: Entrega el valor que cuesta una caja especÃ­fica de cerveza dependiendo de los parÃ¡metros ingresados, esto quiere decir que multiplique el precio por la cantidad una vez se homologara la moneda a lo que se ingreso por parÃ¡metro.
	- Quantity: Cantidad de cervezas a comprar (valor por defecto 6).
	- Currency: Tipo de moneda con la que desea pagar, para este caso se recomienda que utilice esta API <https://currencylayer.com/>

### Requisitos

- Puede usar alguno de los siguientes lenguajes Java, NodeJS, Go o Python. Aunque valoramos el uso de GO.
- Usar Docker y Docker Compose para los diferentes servicios.
- Se puede usar librarÃ­as externas y frameworks
- Requisito un 70% de cobertura de cÃ³digo
- Completa libertad para agregar nuevas funcionalidades.

### Entrega

- Enviar el link del repositorio donde se realiza este ejercicio.