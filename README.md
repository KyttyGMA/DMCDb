# Dual Memory Contextual Database (DMCDB)

## Descripción
El proyecto **Dual Memory Contextual Database (DMCDB)** es una base de datos experimental diseñada para optimizar el acceso a datos mediante una estructura de memoria primaria y secundaria. Utiliza un enfoque innovador basado en un **espacio topológico unidimensional** que se pliega sobre sí mismo, permitiendo acceder a diferentes "capas" de los datos según el contexto de la solicitud.

## Características Principales
- **Memoria Dual**: Los datos se almacenan en dos niveles de memoria, primaria (RAM) y secundaria (disco).
- **Acceso Contextual**: Según el tipo de consulta, se accede a diferentes capas de información.
- **Optimización de Recursos**: Al plegar el espacio de datos, se minimiza el uso de memoria primaria y se optimiza el acceso a la información relevante.
- **Sistema de Caché**: Datos frecuentemente consultados son almacenados en memoria primaria para acceso rápido.

## Tecnologías Utilizadas
- **Lenguaje**: Go (o el lenguaje que prefieras)
- **Sistema de Almacenamiento**: Redis para memoria primaria y Cassandra para almacenamiento persistente (opcional).
- **Algoritmos de Acceso**: Algoritmo de plegado de espacio topológico para el acceso contextual.

## Instalación
1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu-usuario/dmcdb.git
