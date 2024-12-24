ARG NODE_VERSION="18.18.1"
ARG ALPINE_VERSION="3.18"
ARG NGINX_VERSION="1.16.0"

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build
RUN ls

FROM nginx:${NGINX_VERSION}-alpine AS runner

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
