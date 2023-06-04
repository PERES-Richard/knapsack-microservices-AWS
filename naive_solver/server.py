import time
from os import getenv

from sanic import Sanic
from sanic_ext import Extend

from sanic.response import json, text
from sanic.log import logger

from uuid6 import uuid7

from solver.solver import bestPathFinder

PORT=int(getenv("PORT"))

app = Sanic(name='naive_solver')
app.config.CORS_ORIGINS = "*"
Extend(app)


def checkArgs(request):
    if request.json["bagSize"] is None or request.json["items"] is None:
        return False
    return True


@app.post("/naive-solver/solve")
async def solve_handler(request):
    if not checkArgs(request):
        return json({"Error": "Parameters are invalid.", "bagSize": request.json["bagSize"], "items": request.json["items"]}, status=400)

    # Creates a UUID for the incoming req
    uid = uuid7()

    logger.info("New solve request id "+ str(uid) +" for a bagSize of : " + str(request.json["bagSize"]) + " and " + str(len(
        request.json["items"])) + " items.")

    # Record time for benchmark analysis
    start = time.time()

    # Start solving
    max, items = bestPathFinder(request.json["bagSize"], request.json["items"])

    end = time.time()
    logger.info("Request id " + str(uid) + " solved in " + str(end - start) + " seconds.")

    return json({
        "max_value": max,
        "items": items,
        "duration": end - start
    })

@app.get("/naive-solver/health")
async def healthcheck_handler(request):
    return text("naive solver is healthy")

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=PORT, fast=True)
