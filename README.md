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

# Screenshots
<img width="1903" height="965" alt="Screenshot 2026-03-10 at 9 11 30" src="https://github.com/user-attachments/assets/b58e86dd-1c03-4e73-b12a-b9c4e1335be3" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/10a5b57b-6ac1-44fc-b3a7-f92cfbd9c134" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/ab9db75a-38cc-4e9c-8858-92bf1f347d9d" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/1789667a-43e9-4436-a75d-e776288bef03" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/d202d5d0-6ecc-44a2-bcdd-ba670b75eb86" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/c935e865-0a83-48d0-beed-9493fee1ae49" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/f8b54f9e-35bd-44a0-9ba2-3992d88c8c72" />

<img width="1903" height="965" alt="image" src="https://github.com/user-attachments/assets/89872279-78d8-4d0d-8398-d350a27bf715" />
