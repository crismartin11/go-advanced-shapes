# go-advanced-shapes

# GO Avanzado - Consignas del ejercicio
En base al ejercicio anterior se deben realizar las siguientes actualizaciones:
•	Implementar GinGonic con 2 rutas, uno para la lectura y creación del archivo y otra la inserción en la tabla, la lambda debe ser proxy-request. 
•	Implementar el patrón Adapter para la comunicación con dynamodb
•	El calculo de Area de las figuras se deberá realizar de manera concurrente (channels y goroutines)
•	Las consultas a DynamoDB se deben hacer con PartiQL(solo para la Query)
•	El Read a DynamoDB deberá estar paginado
•	Se tiene que testear con Ginkgo y Testify(mocks) hasta alcanzar al menos el 70% del coverage


# Descripción
Proyecto desarrollado en Golang. El desarrollo consiste en una lambda para la creación de figuras y generación de archivo txt con listado de figuras.

En el caso de la generación de archivo txt recibe como parámetro de entrada el campo "tipo" que debe ser ELLIPSE, RECTANGLE o TRIANGLE. Con ese parámetro se consulta a la tabla devShapes para obtener las figuras del tipo indicado. Luego genera un archivo txt que sube al folder SHAPES/ del S3.

En el caso de la creación recibe como parámetros de entrada el campo "tipo" que debe ser ELLIPSE, RECTANGLE o TRIANGLE, el campo id (tipo string del "1" al "12"), y los campos a y b de tipo floar64.

Tanto para el proceso de generación de archivo txt como para creación de figura se genera un endpoint al que se puede acceder desde Postman. No requieren pasar credenciales.

# Despliegue en uala-global.labssupport-dev
A continuación se encuentra la sección de Build and Deploy, aún así esto ya se encuentra desplegado en la cuenta uala-global-labssuport-dev para pruebas. Se puede probar desde Postman con los request de ejemplo más abajo.

endpoints:
  POST - https://z2f0loob18.execute-api.us-east-1.amazonaws.com/dev/read
  POST - https://z2f0loob18.execute-api.us-east-1.amazonaws.com/dev/create
functions:
  api: go-advanced-shapes-dev-api


Ejemplo de request para el endpoint /read:
{
  "tipo": "TRIANGLE"
}

Ejemplo de request para el endpoint /create:
{
  "tipo": "TRIANGLE",
  "a": 3.9,
  "b": 5.9,
  "id": "6"
}


# Build and deploy
Tanto el build como el deploy pueden ejecutarse por medio del Makefile incluido en el pryecto (método no disponible en windows). Requiere de la instalación de Serverless Framework para el despliegue. 


# Comentarios
 - Paginación con páginas de 2 registros, para ver comportamiento (logs)
 - Hay logs en exceso en algunas partes solo para ver el comportamiento de, por ejemplo, uso de channels. En otros lugares faltaron logs, por ejemplo en algunos de los servicios.
 - Las constants son configuraciones que deberían ir como variables de ambiente.
 - Dejé algunas lineas comentadas concientemente, para destacar el cambio. También alguno métodos deprecados, y notas en varios lugares.
 - El uso de goroutines y channels no se aplicó directamente sobre el cálculo del área sino a la generación del detalle de cada figura, ya que el método detail() contien el cálculo del área (ver models/shape.go)
 - Para implementar el patrón adapter se agregó un servicio (service_for_adapter.go) que ejecuta las consultas a dynamo con PartiQL y usa paginación (aprovecho a aplicar esos dos requisitos del ejercicio), algo que debería estar en el repositorio pero lo estructuré así para simular un servicio que requiere adaptación.
