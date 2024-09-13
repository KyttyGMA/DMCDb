# Modelo de Datos en DualMemoryDB

## Introducción
DualMemoryDB emplea un modelo de datos de **dos capas** para gestionar grandes volúmenes de información de manera eficiente. Este modelo está compuesto por:

- **Capa Primaria**: Almacena metadatos en memoria (RAM) usando **hashmaps** para un acceso rápido.
- **Capa Secundaria**: Contiene los datos completos en almacenamiento persistente, utilizando **metadata** para la organización y el acceso. La persistencia de los datos en esta capa es opcional y configurable.

## Estructura de Datos en la Capa Primaria
La capa primaria utiliza **hashmaps** para almacenar metadatos clave que proporcionan una vista rápida de los datos. Los metadatos incluyen:

- **ID**: Identificador único del dato.
- **Tamaño**: Tamaño del dato completo.
- **Fecha de creación/modificación**: Tiempos relevantes para el dato.
- **Hash**: Hash del dato para validación de integridad.
- **Etiquetas**: Categorías o etiquetas asociadas.

## Estructura de Datos en la Capa Secundaria
La capa secundaria se encarga de almacenar los **datos completos** de manera persistente. Esta capa usa **metadata** para organizar y acceder a los datos almacenados. La persistencia en esta capa es opcional y puede ser configurada según las necesidades del sistema.

## Ventajas de la Proyección de Metadatos
1. **Acceso rápido**: Los metadatos en la capa primaria permiten consultas rápidas sin cargar los datos completos.
2. **Búsqueda eficiente**: La capa primaria facilita búsquedas basadas en los metadatos (por ejemplo, por etiquetas o fechas).
3. **Validación de datos**: El hash en los metadatos permite verificar la integridad del dato completo cuando se accede a la capa secundaria.

## Consideraciones de Persistencia
La persistencia en la capa secundaria es opcional. Dependiendo de los requisitos del sistema, puedes configurar el almacenamiento de datos completos en disco o en otro medio persistente. La capa primaria siempre mantendrá metadatos en memoria para garantizar un acceso rápido.
