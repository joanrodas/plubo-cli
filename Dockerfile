FROM scratch
COPY plubo-cli /usr/bin/plubo-cli
ENTRYPOINT ["/usr/bin/plubo-cli"]