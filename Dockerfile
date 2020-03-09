FROM alpine
ADD weather-srv /weather-srv
ENTRYPOINT [ "/weather-srv" ]
