
## Flujo de Operaciones en la Base de Datos de Memoria Dual ##

### Introducción

Este documento describe el flujo de operaciones en la base de datos contextual de memoria dual, detallando cómo se manejan las inserciones, consultas, actualizaciones y eliminaciones en el sistema, así como el manejo de fallos. La base de datos está estructurada en dos capas: la capa primaria (RAM) y la capa secundaria (almacenamiento persistente en disco).

### Inserciones

1. **Inserción en la Capa Secundaria**:
   - Cuando se insertan nuevos datos en el sistema, se guardan primero en la **capa secundaria** (almacenamiento en disco). Esto asegura que los datos se mantengan de forma persistente y no se pierdan en caso de fallos del sistema.
   - La capa secundaria almacena los datos en su forma completa y detallada.

2. **Sincronización con la Capa Primaria**:
   - Tras la inserción en la capa secundaria, los datos relevantes se copian o se resumen en la **capa primaria** (RAM) si se consideran importantes para las consultas rápidas.
   - La capa primaria se actualiza para incluir estos nuevos datos, asegurando que los datos más frecuentemente accesibles estén disponibles en memoria para un acceso rápido.

### Consultas

1. **Consulta en la Capa Primaria**:
   - Cuando se realiza una consulta, el sistema primero busca los datos en la **capa primaria** (RAM).
   - Si los datos están presentes en la capa primaria, se devuelve la información solicitada inmediatamente.

2. **Consulta en la Capa Secundaria**:
   - Si los datos no se encuentran en la capa primaria, el sistema consulta la **capa secundaria** (almacenamiento en disco) para recuperar la información completa.
   - Una vez obtenidos los datos de la capa secundaria, se pueden actualizar o cargar en la capa primaria para futuras consultas, si se consideran relevantes para el acceso rápido.

3. **Optimización del Caché**:
   - El sistema puede emplear técnicas como **Least Recently Used (LRU)** para gestionar los datos en la capa primaria, manteniendo en RAM solo los datos más relevantes o frecuentemente solicitados.

### Actualizaciones

1. **Actualización en la Capa Secundaria**:
   - Cuando se actualizan los datos, los cambios se aplican primero en la **capa secundaria**. Esto garantiza que la información persistente sea precisa y esté actualizada.

2. **Sincronización con la Capa Primaria**:
   - Después de actualizar la capa secundaria, el sistema puede reflejar estos cambios en la **capa primaria**.
   - La capa primaria se actualiza para reflejar las modificaciones, asegurando que los datos recientes estén disponibles para consultas rápidas.

### Eliminaciones

1. **Eliminación en la Capa Secundaria**:
   - Cuando se eliminan datos, la eliminación se realiza en la **capa secundaria**. Los datos se eliminan de manera persistente del almacenamiento en disco.

2. **Eliminación en la Capa Primaria**:
   - Una vez que los datos son eliminados de la capa secundaria, también se deben eliminar de la **capa primaria**.
   - Los datos eliminados se eliminan de la RAM, o se marcan para su eliminación si aún están en uso o en caché.

### Manejo de Fallos

1. **Fallos en la Capa Primaria (RAM)**:
   - Si ocurre un fallo en la capa primaria (por ejemplo, pérdida de datos en RAM), los datos pueden ser recuperados desde la **capa secundaria**.
   - El sistema debe tener mecanismos para reconstruir o recargar los datos de la capa secundaria a la capa primaria para restaurar la funcionalidad completa.

2. **Resiliencia y Redundancia**:
   - El diseño de la base de datos debe incluir estrategias de resiliencia y redundancia para minimizar el impacto de fallos en la capa primaria.
   - Los datos en la capa secundaria deben estar respaldados y protegidos para garantizar la integridad y la disponibilidad continua en caso de fallos.

### Conclusión

El flujo de operaciones en la base de datos contextual de memoria dual asegura que las inserciones, consultas, actualizaciones y eliminaciones se manejen de manera eficiente entre las capas primaria y secundaria. La sincronización y el manejo adecuado de fallos garantizan un rendimiento óptimo y una alta disponibilidad de datos. Este enfoque permite una base de datos robusta y escalable, adecuada para aplicaciones que requieren un acceso rápido y fiable a grandes volúmenes de datos.
