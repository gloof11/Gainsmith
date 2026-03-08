# Build the Frontend
FROM node:latest AS frontend-build 

WORKDIR /app

COPY ./package.json /app/package.json
RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build-only

# Initialize Nginx
FROM nginx AS nginx-build
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Copy the app build to nginx
COPY --from=frontend-build /app/dist /usr/share/nginx/html 

EXPOSE 80

ARG PB_VERSION=0.36.2
RUN apt update && apt install unzip
# download and unzip PocketBase
ADD https://github.com/pocketbase/pocketbase/releases/download/v${PB_VERSION}/pocketbase_${PB_VERSION}_linux_amd64.zip /tmp/pb.zip
RUN unzip /tmp/pb.zip -d /pb/

COPY ./entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "./entrypoint.sh" ]
