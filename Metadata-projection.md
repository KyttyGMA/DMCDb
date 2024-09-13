# Proyección de Metadatos en DualMemoryDB

## Introducción

En **DualMemoryDB**, la proyección de datos en la capa primaria se basa en el uso de **metadatos**. Este enfoque permite una rápida referencia a información esencial sin necesidad de cargar los datos completos desde el almacenamiento secundario, optimizando así el acceso y la gestión de los datos en la base de datos.

## Estructura de los Metadatos

Los metadatos en la capa primaria incluyen la siguiente información clave:

- **ID**: Identificador único del dato, generalmente un hash, que permite una referencia rápida.
- **Tamaño**: Tamaño del dato completo en bytes, útil para la gestión de recursos y optimización.
- **Fecha de Creación/Modificación**: Registra cuándo fue creado o modificado el dato, facilitando la gestión de versiones y auditoría.
- **Hash de Verificación**: Hash del dato completo utilizado para validar su integridad y detectar posibles corrupciones.
- **Etiquetas**: Conjunto de etiquetas que facilita la búsqueda eficiente y la clasificación de los datos en la capa primaria.

## Sincronización entre Capas

El flujo de datos y metadatos entre capas funciona de la siguiente manera:

1. **Inserción de Datos**:
   - Los datos completos se almacenan en la **capa secundaria (almacenamiento persistente)**.
   - Los metadatos correspondientes se guardan en la **capa primaria (RAM)** para un acceso rápido.

2. **Consulta de Datos**:
   - Al consultar un dato, el sistema primero accede a los metadatos en la capa primaria para determinar la disponibilidad y características del dato.
   - Si los datos completos son necesarios, se accede a la **capa secundaria** para recuperar la información completa utilizando la información proporcionada por los metadatos.

## Ventajas de la Proyección de Metadatos

1. **Acceso Rápido**: Los metadatos en la capa primaria permiten realizar consultas rápidas sin necesidad de cargar los datos completos desde el almacenamiento secundario.
2. **Búsqueda Eficiente**: La capa primaria facilita la búsqueda basada en metadatos (por ejemplo, etiquetas o fechas), optimizando el rendimiento de las consultas.
3. **Validación de Datos**: El hash de verificación en los metadatos permite validar la integridad del dato completo cuando se accede a la capa secundaria, asegurando que la información no ha sido alterada.

