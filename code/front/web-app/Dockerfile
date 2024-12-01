ARG NODE_VERSION="18.18.1"
ARG ALPINE_VERSION="3.18"
ARG NGINX_VERSION="1.16.0"

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /app

COPY . .

RUN npm
CMD ["npm", "start", "--host", "0.0.0.0"]
