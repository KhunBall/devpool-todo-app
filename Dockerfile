### STAGE 1: Build Angular ###
FROM node:16.16.0-alpine AS builder
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

### STAGE 3: Build Go ###
FROM golang:1.18-alpine AS build
WORKDIR /app
COPY goapi/go.mod ./
COPY goapi/go.sum ./
COPY goapi/*.go ./
COPY goapi/db ./db
COPY goapi/users ./users
RUN go mod download
COPY *.go ./
RUN go build -o /goapi

## STAGE Final: Run Angular ###
# FROM alpine:latest
FROM alpine:3.16.0
RUN apk update && apk add --no-cache supervisor nginx
COPY supervisord.conf /etc/supervisord.conf
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=build /goapi /app/
EXPOSE 80
EXPOSE 8080
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]


# ## STAGE 2: Run Angular ###
# FROM nginx:1.23.0-alpine
# COPY nginx.conf /etc/nginx/conf.d/default.conf
# COPY --from=builder /app/dist /usr/share/nginx/html
# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]


# ### STAGE 4: Run Go ###
# FROM gcr.io/distroless/base-debian10
# WORKDIR /
# COPY --from=build /docker-goapi /docker-goapi
# EXPOSE 8080
# USER nonroot:nonroot
# ENTRYPOINT ["/docker-goapi"]


