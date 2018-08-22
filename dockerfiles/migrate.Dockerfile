FROM chapsuk/miga:v0.5.0
COPY ./migrations                ./migrations
COPY ./dockerfiles/dev/seeds     ./seeds/development