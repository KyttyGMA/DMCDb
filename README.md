# Dual Memory Contextual Database (DMCDB)

## Description
The **Dual Memory Contextual Database (DMCDB)** project is an experimental in-memory database designed to optimize data access through a primary and secondary memory structure. It uses an innovative approach based on a **one-dimensional topological space** that folds onto itself, allowing access to different "layers" of data according to the context of the request. This project is a Proof of Concept (POC).

## Key Features
- **Dual Memory**: Data is organized across two tiers of memory: primary (RAM) and secondary (persistent storage). The primary memory uses hashmaps to store metadata for fast access, while the secondary storage uses metadata for organizing and accessing complete data. Persistence in the secondary layer is optional and configurable.
- **Contextual Access**: Depending on the type of query, different layers of information are accessed based on the context.
- **Resource Optimization**: By folding the data space, the use of primary memory is minimized and access to relevant information is optimized.
- **Caching System**: Frequently queried data is stored in primary memory for quick access.
- **Standalone**: The database is developed as a standalone system without dependencies on external storage systems.
- **Proof of Concept (POC)**: This project is a proof of concept and may not include all production-ready features.

## Technologies Used
- **Language**: Go
- **In-Memory Storage**: Utilizes Go's built-in data structures for primary and secondary memory.
- **Access Algorithms**: Topological space folding algorithm for contextual access.

## Data Model

### Primary Layer
- **Data Structure**: Uses hashmaps to store metadata, which includes:
  - **ID**: Unique identifier of the data.
  - **Size**: Size of the complete data.
  - **Creation/Modification Date**: Timestamps for the data.
  - **Hash**: Hash for data integrity validation.
  - **Tags**: Associated categories or labels.

### Secondary Layer
- **Data Structure**: Stores complete data with optional persistence. This layer uses metadata to organize and access the data. The persistence of data in this layer can be configured based on the requirements.

## Advantages
- **Fast Access**: Metadata in the primary layer allows for rapid queries without needing to load the full data.
- **Efficient Searching**: The primary layer supports efficient searches based on metadata (e.g., tags, dates).
- **Data Integrity**: Hashes in metadata enable verification of the complete data's integrity when accessed from the secondary layer.


## License
This project is licensed under the [MIT License](LICENSE).
