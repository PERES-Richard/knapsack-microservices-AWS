# pull latest julia image
FROM --platform=linux/amd64 julia:latest

# create dedicated user
RUN useradd --create-home --shell /bin/bash genie

# set up the app
RUN mkdir /home/genie/app
COPY . /home/genie/app
WORKDIR /home/genie/app

# C compiler for PackageCompiler
RUN apt-get update && apt-get install -y g++

# configure permissions
RUN chown -R genie:genie /home/

RUN chmod +x bin/repl
RUN chmod +x bin/server
RUN chmod +x bin/runtask

# switch user
USER genie

# instantiate Julia packages
RUN julia -e "using Pkg; Pkg.activate(\".\"); Pkg.instantiate(); Pkg.precompile(); "

# Compile app
RUN julia --project compiled/make.jl

# ports
EXPOSE 3000
EXPOSE 3000

# set up app environment
ENV JULIA_DEPOT_PATH "/home/genie/.julia"
ENV GENIE_ENV "prod"
ENV GENIE_HOST "0.0.0.0"
ENV PORT "3000"
ENV WSPORT "3000"
ENV EARLYBIND "true"

# run app
CMD ["bin/server-prod"]

# or maybe include a Julia file
# CMD julia -e 'using Pkg; Pkg.activate("."); include("IrisClustering.jl"); '
