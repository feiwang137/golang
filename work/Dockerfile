FROM golang:1.18-alpine
COPY bin/httpServer /tmp/httpServer
RUN mv /tmp/httpServer /httpServer
EXPOSE 80
USER nobody
CMD /httpServer
