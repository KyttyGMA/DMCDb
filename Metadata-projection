# Proyección de Metadatos en DualMemoryDB

## Introducción
En la **DualMemoryDB**, la proyección de datos en la capa primaria se basa en el uso de **metadatos**. Esto permite una rápida referencia a información esencial de los datos sin necesidad de cargar los datos completos desde el almacenamiento secundario.

## Estructura de los Metadatos
Los metadatos incluyen la siguiente información clave:

- **ID**: Un identificador único del dato, generalmente un hash.
- **Tamaño**: El tamaño del dato completo en bytes.
- **Fecha de creación/modificación**: Registra cuándo fue creado o modificado el dato.
- **Hash de verificación**: Un hash del dato completo que permite validar su integridad.
- **Etiquetas**: Un conjunto de etiquetas que facilitan la búsqueda eficiente.

## Sincronización entre Capas
Cuando se inserta un dato en la base de datos, los metadatos se guardan en la **capa primaria (RAM)**, mientras que los datos completos se almacenan en la **capa secundaria (almacenamiento)**. Al consultar un dato, primero se accede a los metadatos en la capa primaria y, si es necesario, se accede a los datos completos en la capa secundaria.

## Ventajas de la Proyección de Metadatos
1. **Acceso rápido**: Los metadatos en la capa primaria permiten consultas rápidas sin cargar los datos completos.
2. **Búsqueda eficiente**: Se pueden realizar búsquedas en la capa primaria basadas en los metadatos (por ejemplo, por etiquetas o fechas).
3. **Validación de datos**: El uso de un hash en los metadatos permite validar la integridad del dato completo cuando se accede a la capa secundaria.
