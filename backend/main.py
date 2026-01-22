from fastapi import FastAPI
from routes import router
from midleware import logging_midleware

app = FastAPI(
    title="Digital Library API",
    description="REST API for managing a digital library",
    version="1.0.0"
)

app.middleware("http")(logging_midleware)
app.include_router(router)