FROM alpine:latest

RUN mkdir /app

COPY profileApp /app

CMD [ "/app/profileApp" ]