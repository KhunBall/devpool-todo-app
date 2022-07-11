# ### STAGE 1: Build Angular ###
# FROM node:16.16.0-alpine AS builder
# WORKDIR /app
# COPY package.json ./
# RUN npm install
# COPY . .
# RUN npm run build

# ### STAGE 2: Run Angular ###
# FROM nginx:1.23.0-alpine
# COPY nginx.conf /etc/nginx/conf.d/default.conf
# COPY --from=builder /app/dist /usr/share/nginx/html
# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]

### STAGE 3: Build Go ###
FROM golang:1.18-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
COPY *.go ./
RUN go build -o /docker-goapi

### STAGE 4: Run Go ###
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /docker-goapi /docker-goapi
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/docker-goapi"]

# ### STAGE 3: Build Go ###
# FROM golang:1.18-alpine AS build
# WORKDIR /app
# COPY . .
# RUN go mod download
# COPY *.go ./
# RUN go build -o /docker-goapi