# Build the Frontend
FROM node:25-alpine3.22 AS frontend-build 

WORKDIR /app/frontend

COPY frontend/package.json frontend/package-lock.json ./ 
RUN npm ci
COPY frontend/ ./ 
RUN npm run build .

#Build the Backend 
FROM golang:1.26.2-alpine3.23 AS backend-build
RUN apk add --no-cache tzdata
RUN apk add --no-cache make 

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
COPY --from=frontend-build /app/frontend/dist ./ui/dist/
RUN make build 

# Create the runtime
FROM alpine:3.22.4
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=backend-build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=backend-build /app/bin/gainsmith /app/gainsmith

EXPOSE 8090

ENTRYPOINT [ "/app/gainsmith" ]
CMD ["serve", "--http=0.0.0.0:8090"]
