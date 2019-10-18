# build stage
FROM node:lts-alpine as build-stage
ENV ALPINE_MIRROR "http://mirrors.aliyun.com/alpine"
RUN echo "${ALPINE_MIRROR}/latest-stable/main" > /etc/apk/repositories;
RUN echo "${ALPINE_MIRROR}/latest-stable/community" >> /etc/apk/repositories;

RUN apk update && \
    apk add git python python-dev build-base
WORKDIR /app
COPY package*.json ./
RUN npm install --registry=https://registry.npm.taobao.org
COPY . .
RUN npm run build:prod

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
# EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
