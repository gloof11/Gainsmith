# Gainsmith

This is a workout tracker that I made for myself personally. It's made with VueJS on the frontend and Pocketbase in the backend. This project was a culmination of learning VueJS and making responsive webapps, and I'd say I'm pretty proud of it :)

# Installation 

- Clone the repository
- Copy the `.env.example` file, and rename it to `.env`. Set the `VITE_APP_POCKETBASE_URL=` variable to point to your Gainsmith Instance IP (Example: "localhost" if you are running on your local machine. Or use the IP of the VPS you will be hosting it on). 
- Start the docker compose file with the following command:
`docker compose up -d --build`
- Open the logs for the container and find the default admin credentials for PocketBase. PocketBase generates a new token for superuser on first start up. It will look similar to this: 

```
$ docker logs -f gainsmith-pub-gainsmith-1           
2026/03/08 11:08:56 [notice] 14#14: using the "epoll" event method
2026/03/08 11:08:56 [notice] 14#14: nginx/1.29.5
...

(!) Launch the URL below in the browser if it hasn't been open already to create your first superuser account:
http://<gainsmith IP>/_/#/pbinstal/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJwYmNfMzE0MjYzNTgyMyIsImV4cCI6MTc3Mjk2OTkzNiwiaWQiOiI2aTNsYWI1M29rYzJhdWIiLCJyZWZyZXNoYWJsZSI6ZmFsc2UsInR5cGUiOiJhdXRoIn0.oYFrgq8m3dRK7fS3FkM6Y4AcYRVVnNv3z0EQleX9StI
```

Enjoy!