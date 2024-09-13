# Dual Memory Contextual Database (DMCDB)

## Description
The **Dual Memory Contextual Database (DMCDB)** project is an experimental in-memory database designed to optimize data access through a primary and secondary memory structure. It uses an innovative approach based on a **one-dimensional topological space** that folds onto itself, allowing access to different "layers" of data according to the context of the request. This project is a Proof of Concept (POC).

## Key Features
- **Dual Memory**:Data is organized across two tiers of memory: primary and secondary, both residing in RAM.
- **Contextual Access**: Depending on the type of query, different layers of information are accessed.
- **Resource Optimization**: By folding the data space, the use of primary memory is minimized and access to relevant information is optimized.
- **Caching System**: Frequently queried data is stored in primary memory for fast access.
- **Standalone**: The database is developed as a standalone system without dependencies on external storage systems. 
- **Proof of Concept (POC)**: This project is a proof of concept and may not include all production-ready features.

## Technologies Used
- **Language**: Go
- **In-Memory Storage**: Utilizes Go's built-in data structures for primary and secondary memory.
- **Access Algorithms**: Topological space folding algorithm for contextual access.

