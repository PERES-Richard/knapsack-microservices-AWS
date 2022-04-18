import time

from sanic import Sanic
from sanic.response import json, text
from sanic.log import logger

from solver.solver import bestPathFinder

app = Sanic(name='naive_solver')


@app.route("/solve")
async def solve_handler(request):
    logger.info("New solve request for a bagSize of : " + str(request.json["bagSize"]) + " and " + str(len(
        request.json["items"])) + " items.")
    start = time.time()
    max, items = bestPathFinder(request.json["bagSize"], request.json["items"])
    end = time.time()
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
