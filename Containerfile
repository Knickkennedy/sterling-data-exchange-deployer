FROM docker.io/golang:1.24.3

LABEL author="Knicholas Kennedy"
LABEL product="sde-deployer"

ENV PIP_ROOT_USER_ACTION=ignore

USER 0

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y python3 && \
    apt-get install -y python3-pip && \
    apt-get install -y ansible && \
    apt-get install -y python3-kubernetes

VOLUME [ "/data" ]

RUN mkdir -p /sde-deployer && \
    mkdir -p /data

COPY . /sde-deployer/

WORKDIR /sde-deployer

COPY --from=golang:1.24.3 /usr/local/go /usr/local/go

RUN echo 'export PS1="\u@\w: "' >> /etc/bash.bashrc

ENV USER_UID=1001

ENV GOCACHE=/tmp/go-build

ENV HELM_VERSION="v3.16.3"

RUN curl -L https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar zx && \
    mv linux-amd64/helm /usr/local/bin/helm && \
    chmod +x /usr/local/bin/helm && \
    rm -rf linux-amd64

RUN curl -s -L -o openshift-client-linux.tar.gz https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/stable/openshift-client-linux.tar.gz && \
    tar -xvf openshift-client-linux.tar.gz -C /usr/local/bin/ && \
    rm -rf openshift-client-linux.tar.gz

RUN chown -R ${USER_ID}:0 /data && \
    chown -R ${USER_ID}:0 /sde-deployer && \
    chmod -R ug+rwx /sde-deployer/*.go

ENTRYPOINT [ "go", "run", "/sde-deployer/entrypoint.go" ]
CMD []
