# Sistema de Arquitectura de la Base de Datos de Memoria Dual

## Introducción

El sistema de base de datos que estamos desarrollando se basa en un enfoque **contextual de memoria dual**, organizando los datos en dos capas en el mismo espacio de **RAM**. La **capa primaria** usa **hashmaps** para ofrecer acceso rápido a datos frecuentemente utilizados, mientras que la **capa secundaria** utiliza metadatos para almacenar los datos completos de manera estructurada. Esta arquitectura busca maximizar el rendimiento y la escalabilidad, al mismo tiempo que mantiene una estructura eficiente para el acceso a los datos.

## Componentes Principales

1. **Capa Primaria (RAM)**
   - La capa primaria almacena los datos más frecuentemente utilizados en **RAM**.
   - Utiliza **hashmaps** para proporcionar acceso en tiempo constante **O(1)**, facilitando consultas rápidas y efectivas.
   - Contiene una versión reducida o parcial de los datos completos que residen en la capa secundaria, optimizando el uso de memoria.

2. **Capa Secundaria**
   - La capa secundaria también reside en **RAM**, pero organiza los datos de manera que los metadatos permiten acceder a información completa.
   - Esta capa garantiza que la información proyectada en la capa primaria se pueda recuperar en su totalidad, utilizando un diseño de almacenamiento eficiente.
   - Aunque el acceso es más lento en comparación con la capa primaria, asegura la persistencia y capacidad de manejar grandes volúmenes de información.

3. **Contexto de Acceso a Datos**
   - El **contexto** define qué datos se cargan en la **capa primaria** y cómo se accede a ellos.
   - Dependiendo de la solicitud, el sistema consulta la capa primaria para un subconjunto de datos o la capa secundaria para la información completa.
   - El contexto puede ser determinado por la frecuencia de acceso, la importancia de los datos o solicitudes específicas de los usuarios.

## Flujo de Datos

El flujo de datos en el sistema sigue estos pasos:

1. **Escritura de Datos**
   - Nuevos datos se almacenan en la **capa secundaria** en su forma completa y detallada.
   - Un subconjunto de estos datos, o una versión simplificada, se guarda en la **capa primaria** según el contexto o la frecuencia de acceso.

2. **Lectura de Datos**
   - Las solicitudes de datos se dirigen primero a la **capa primaria (RAM)**. Si los datos están disponibles, se devuelve la respuesta rápidamente.
   - Si los datos no están en la capa primaria o se necesita una versión más detallada, se consulta la **capa secundaria**.

3. **Sincronización de Capas**
   - Las capas deben **sincronizarse** cuando se actualizan los datos. Las actualizaciones en la capa secundaria pueden reflejarse en la capa primaria si los datos son de uso frecuente.
   - La capa primaria puede descargarse de la memoria si los datos no son utilizados frecuentemente, reduciendo la carga de la RAM.

## Estrategias de Gestión de Memoria

Para gestionar eficientemente la capa primaria, se pueden aplicar estrategias como:

- **LRU (Least Recently Used)**: Elimina datos no accedidos recientemente para liberar espacio en la capa primaria.
- **TTL (Time to Live)**: Establece un tiempo de vida para los datos en la capa primaria. Los datos expirados se eliminan de la RAM.

## Interacción entre Capas

La interacción entre las capas es clave para equilibrar **rendimiento** y **persistencia**:

- **Lecturas frecuentes** se sirven desde la **capa primaria** para reducir la latencia.
- **Escrituras y lecturas esporádicas** utilizan la **capa secundaria**, asegurando la persistencia y la integridad de los datos.
- Las capas mantienen comunicación constante, reflejando cambios en los datos de manera eficiente.

## Escalabilidad y Optimización

Esta arquitectura permite escalar eficientemente en sistemas con grandes volúmenes de datos. La **capa primaria** se enfoca en los datos más relevantes o frecuentemente solicitados, mientras que la **capa secundaria** maneja grandes volúmenes sin afectar el rendimiento. El uso de **hashmaps**, junto con estrategias de gestión de memoria, optimiza el uso de **RAM** y **almacenamiento en disco**, ideal para entornos de alto rendimiento.

## Conclusión

La arquitectura de esta base de datos contextual de memoria dual proporciona una solución eficiente para manejar grandes volúmenes de datos en tiempo real, combinando **RAM** y **almacenamiento persistente**. El uso de **hashmaps** en la capa primaria asegura consultas rápidas, mientras que la capa secundaria garantiza la persistencia de datos a largo plazo. Esta solución está diseñada para maximizar rendimiento y escalabilidad, ofreciendo robustez para aplicaciones de alta capacidad de respuesta.
