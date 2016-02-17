# Installs our Rancher Compose command line runner application
FROM        golang:1.4.3-cross
MAINTAINER  Thane Thomson <connect@thanethomson.com>

# Install Rancher Compose
ENV         RANCHER_COMPOSE_VERSION 0.7.1
RUN         cd /tmp && \
            git clone https://github.com/rancher/rancher-compose.git && \
            cd rancher-compose && \
            git checkout tags/v${RANCHER_COMPOSE_VERSION} && \
            chmod +x scripts/build && \
            ./scripts/build && \
            cp bin/rancher-compose /bin/ && \
            cd / && \
            rm -rf /tmp/*

ADD         drone-rancher-compose /bin/
ENTRYPOINT  ["/bin/drone-rancher-compose"]
