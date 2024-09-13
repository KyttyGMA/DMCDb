
# Modelo de Datos en DualMemoryDB

## Introducción al Modelo de Datos
DualMemoryDB utiliza un modelo de datos de **dos capas** para optimizar el acceso y la gestión de grandes volúmenes de información. El modelo se compone de una **capa primaria** que almacena metadatos en memoria (RAM) y una **capa secundaria** que contiene los datos completos en almacenamiento persistente.

## Estructura de Datos en la Capa Primaria
La capa primaria utiliza un **Hashmap** para almacenar metadatos que proporcionan una vista rápida de los datos sin necesidad de cargar la información completa. Los metadatos incluyen:

- **ID**: Identificador único del dato.
- **Tamaño**: Tamaño del dato completo.
- **Fecha de creación/modificación**: Tiempos relevantes para el dato.
- **Hash**: Hash del dato para validación de integridad.
- **Etiquetas**: Categorías o etiquetas asociadas.

### Ejemplo de Estructura en Go
```go
// Estructura para los metadatos
type Metadata struct {
    ID          string
    Size        int
    CreatedAt   string
    ModifiedAt  string
    Hash        string
    Tags        []string
}
```

## Estructura de Datos en la Capa Secundaria
La capa secundaria se encarga de almacenar los **datos completos**. En esta capa, se guardan los datos como archivos o en una base de datos simple, utilizando el ID para acceder a ellos.

### Ejemplo de Almacenamiento en Go
```go
// Estructura de la base de datos
type DualMemoryDB struct {
    Primary   map[string]Metadata   // Capa primaria 
    Secondary map[string]string     // Capa secundaria 
}
```

## Ejemplos Prácticos

### Insertar Datos
```go
func (db *DualMemoryDB) Insert(key string, value string) {
    hashedKey := hash(key)

    // Crear metadatos
    metadata := Metadata{
        ID:         hashedKey,
        Size:       len(value),
        CreatedAt:  time.Now().Format(time.RFC3339),
        ModifiedAt: time.Now().Format(time.RFC3339),
        Hash:       hash(value),
        Tags:       []string{"general"},
    }

    // Insertar en la capa primaria
    db.Primary[hashedKey] = metadata

    // Insertar el dato completo en la capa secundaria
    db.Secondary[hashedKey] = value
}
```

### Recuperar Datos
```go
func (db *DualMemoryDB) GetFullData(key string) string {
    hashedKey := hash(key)
    return db.Secondary[hashedKey]
}
```

### Recuperar Metadatos
```go
func (db *DualMemoryDB) GetMetadata(key string) Metadata {
    hashedKey := hash(key)
    return db.Primary[hashedKey]
}
```

