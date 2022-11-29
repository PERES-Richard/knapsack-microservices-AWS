import time

from sanic import Sanic
from sanic.response import json, text
from sanic.log import logger

from uuid6 import uuid7

from solver.solver import bestPathFinder

app = Sanic(name='naive_solver')


@app.route("/solve", ["POST"])
async def solve_handler(request):
    # Creates a UUID for the incominq req
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

@app.route("/health")
async def solve_handler(request):
    return text("naive solver is healthy")

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8081)
