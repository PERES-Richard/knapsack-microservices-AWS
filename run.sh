#!/usr/bin/env zsh
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo "${YELLOW}Starting Docker container and binding ports to localhost...${NC}"
docker-compose up -d

echo "${YELLOW}Generating a new KnapSac witha  size of 100 and 12 random items...${NC}"
curl -X POST -s "localhost:8080/generate?bagSize=100&numberOfItems=12" -o data.json > /dev/null
cat data.json

echo "${YELLOW}Solve this KnaSac using a naive algorithm...${NC}"
curl -X POST "localhost:8081/solve" -d @data.json
