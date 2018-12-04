FROM gcr.io/cloud-builders/docker:18.06.1

ENV JFROG_CLI_VERSION=1.22.1

RUN apt-get update && apt-get install -qqy wget curl \
  && wget https://dl.bintray.com/jfrog/jfrog-cli-go/${JFROG_CLI_VERSION}/jfrog-cli-linux-amd64/jfrog \
  && chmod +x jfrog \
  && mv jfrog /usr/bin/ \
  && apt-get remove -qqy --purge wget \
  && rm /var/lib/apt/lists/*_*

ENTRYPOINT ["jfrog"]