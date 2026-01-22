# FastAPI Digital Library

A RESTful API built using FastAPI to manage a digital library.

## Features
- CRUD operations for books
- Strict data validation using Pydantic
- Custom middleware for logging and request timing
- Swagger/OpenAPI documentation

## Setup
```bash
pip install -r requirements.txt
uvicorn backend.main:app --reload