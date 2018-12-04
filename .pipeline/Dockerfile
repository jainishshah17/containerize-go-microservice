FROM gcr.io/cloud-builders/docker:18.06.1

ENV JFROG_CLI_VERSION=1.22.1
ENV HELM_VERSION=v2.10.0

RUN apt-get update && apt-get install -qqy wget curl \
  && wget https://dl.bintray.com/jfrog/jfrog-cli-go/${JFROG_CLI_VERSION}/jfrog-cli-linux-amd64/jfrog \
  && chmod +x jfrog \
  && mv jfrog /usr/bin/ \
  && curl -SL https://storage.googleapis.com/kubernetes-helm/helm-${HELM_VERSION}-linux-amd64.tar.gz -o helm.tar.gz \
  && tar zxvf helm.tar.gz \
  && mv linux-amd64/helm /usr/bin/ \
  && rm -rf linux-amd64 \
  && apt-get remove -qqy --purge wget \
  && rm /var/lib/apt/lists/*_*

RUN helm init -c

ENTRYPOINT ["jfrog"]