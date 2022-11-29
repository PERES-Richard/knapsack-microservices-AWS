# Security best pratices from Snyk https://snyk.io/blog/best-practices-containerizing-python-docker/
# syntax=docker/dockerfile:1

FROM python:3.10-slim AS build

# With default ('root') user
RUN apt-get update
RUN apt-get install -y --no-install-recommends \
    build-essential gcc
RUN pip install --upgrade pip

WORKDIR /usr/app
RUN python -m venv /usr/app/venv
ENV PATH="/usr/app/venv/bin:$PATH"

COPY requirements.txt .
RUN pip install -r requirements.txt


FROM python:3.10-slim

# Use a custom non root user
RUN groupadd -g 999 python && \
    useradd -r -u 999 -g python python

RUN mkdir /usr/app && chown python:python /usr/app

WORKDIR /usr/app
COPY --chown=python:python --from=build /usr/app/venv ./venv
COPY --chown=python:python . .

USER 999

ENV PATH="/usr/app/venv/bin:$PATH"

EXPOSE 8081

CMD [ "python3", "server.py" ]
#CMD [ "gunicorn", "--worker-class uvicorn.workers.UvicornWorker ", "--bind", "0.0.0.0:8081", "server:app" ] # doesnt work since Sanic is not a wsgi
HEALTHCHECK --interval=2s --timeout=3s --start-period=2s --retries=3 CMD curl -f http://localhost:8081/health
