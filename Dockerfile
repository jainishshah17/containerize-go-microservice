FROM k8s-art.jfrog.team/golang:1.11.2-alpine

ENV PORT 8080

# Install dependencies
RUN apk update && apk add --update gcc git musl-dev curl

# Notice the backslash in \$latest, ref: https://github.com/jfrog/jfrog-cli-go/issues/96
RUN curl -Lo /usr/bin/jfrog https://api.bintray.com/content/jfrog/jfrog-cli-go/\$latest/jfrog-cli-linux-386/jfrog?bt_package=jfrog-cli-linux-386 \
   && chmod a+x /usr/bin/jfrog

# Set workspace
WORKDIR /src/jainishshah17/containerize-go-microservice/

# Copy microservice
COPY ./ /src/jainishshah17/containerize-go-microservice/

# Build microservices
RUN cd /src/jainishshah17/containerize-go-microservice/ && go install

CMD ["/go/bin/containerize-go-microservice"]