from fastapi import Request
import time
import logging

logging.basicConfig(level=logging.DEBUG)

async def logging_midleware(request: Request, call_next):
    start_time = time.time()

    UserAgent = request.headers.get("User-Agent")
    logging.info(f"[LOG] Request received from: {UserAgent}")

    response = await call_next(request)
    
    process_time = time.time() - start_time
    response.headers["X-Process-Time"] = str(process_time)

    return response