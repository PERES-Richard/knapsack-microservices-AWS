FROM sanicframework/sanic:3.8-latest

WORKDIR /naive-solver

COPY . .

RUN pip install -r requirements.txt

EXPOSE 8081

CMD ["python", "server.py"]
