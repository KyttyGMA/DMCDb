# Flujo de Operaciones en la Base de Datos de Memoria Dual

## Introducción

Este documento describe el flujo de operaciones en la base de datos contextual de memoria dual, que organiza los datos en dos capas dentro del mismo espacio de **RAM**: la **capa primaria** (RAM) y la **capa secundaria** (RAM, con metadatos para gestión avanzada). Se detallan las inserciones, consultas, actualizaciones y eliminaciones, así como el manejo de fallos y la persistencia opcional en sistemas de IoT.

## Inserciones

1. **Inserción en la Capa Secundaria**:
   - Los nuevos datos se insertan primero en la **capa secundaria**. Esta capa organiza los datos completos y detallados utilizando metadatos para estructurar la información.
   - Aunque la capa secundaria reside en RAM, la arquitectura permite un diseño lógico que facilita la persistencia opcional mediante técnicas como almacenamiento en disco si se requiere.

2. **Sincronización con la Capa Primaria**:
   - Tras la inserción en la capa secundaria, los datos relevantes se copian o se resumen en la **capa primaria** (RAM) si se consideran importantes para consultas rápidas.
   - La capa primaria se actualiza para incluir estos nuevos datos, asegurando que la información frecuentemente accesible esté disponible en memoria para acceso rápido.

## Consultas

1. **Consulta en la Capa Primaria**:
   - El sistema primero busca los datos en la **capa primaria** (RAM).
   - Si los datos están presentes, se devuelve la información solicitada inmediatamente.

2. **Consulta en la Capa Secundaria**:
   - Si los datos no están en la capa primaria, se consulta la **capa secundaria** (RAM) para recuperar la información completa.
   - Los datos obtenidos de la capa secundaria pueden actualizar la capa primaria para futuras consultas si se consideran relevantes para el acceso rápido.

3. **Optimización del Caché**:
   - Se pueden emplear técnicas como **Least Recently Used (LRU)** para gestionar los datos en la capa primaria, manteniendo en RAM solo los datos más relevantes o frecuentemente solicitados.

## Actualizaciones

1. **Actualización en la Capa Secundaria**:
   - Los cambios se aplican primero en la **capa secundaria**, asegurando que la información persistente sea precisa y actualizada.

2. **Sincronización con la Capa Primaria**:
   - Tras actualizar la capa secundaria, el sistema refleja estos cambios en la **capa primaria** para asegurar que los datos recientes estén disponibles para consultas rápidas.

## Eliminaciones

1. **Eliminación en la Capa Secundaria**:
   - La eliminación de datos se realiza en la **capa secundaria** de manera persistente, eliminando los datos del almacenamiento en disco si se utiliza almacenamiento adicional.

2. **Eliminación en la Capa Primaria**:
   - Después de eliminar los datos de la capa secundaria, también se deben eliminar de la **capa primaria**.
   - Los datos eliminados se eliminan de la RAM o se marcan para su eliminación si aún están en uso o en caché.

## Manejo de Fallos

1. **Fallos en la Capa Primaria (RAM)**:
   - En caso de fallos en la capa primaria (por ejemplo, pérdida de datos en RAM), los datos se pueden recuperar desde la **capa secundaria**.
   - El sistema debe contar con mecanismos para reconstruir o recargar los datos de la capa secundaria a la capa primaria para restaurar la funcionalidad completa.

2. **Resiliencia y Redundancia**:
   - El diseño de la base de datos debe incluir estrategias de resiliencia y redundancia para minimizar el impacto de fallos en la capa primaria.
   - La capa secundaria debe estar respaldada y protegida para garantizar la integridad y disponibilidad continua de los datos.

## Conclusión

El flujo de operaciones en la base de datos contextual de memoria dual asegura una gestión eficiente de inserciones, consultas, actualizaciones y eliminaciones entre las capas primaria y secundaria. La sincronización adecuada y el manejo de fallos garantizan un rendimiento óptimo y alta disponibilidad de datos. Este enfoque proporciona una base de datos robusta y escalable, adecuada para aplicaciones que requieren acceso rápido y fiable a grandes volúmenes de datos, con opciones de persistencia adaptadas a entornos de IoT.
