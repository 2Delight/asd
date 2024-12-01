ARG NODE_VERSION="18.18.1"
ARG ALPINE_VERSION="3.18"

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /app

COPY . .

CMD ["npm", "start", "--host", "0.0.0.0"]
