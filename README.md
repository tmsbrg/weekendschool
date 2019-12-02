# Weekendschool Hacking Challenge

Made for school children to learn about IT security. Maybe a bit too hard for them. Still a fun hacking challenge.

Run with docker:

```
docker pull tmsbrg/weekend:1.0
docker run --name weekend -p 127.0.0.1:5000:5000 tmsbrg/weekend:1.0
```

The run command will run the hacking challenge web server on http://localhost:5000/.

To kill the server:
```
docker kill weekend
```

Note that docker commands usually have to be run by root.


Can also be manually installed. Check out nginx-config.test, nginx-config.prod and deploy.sh.
Basically required go to go build backend, and an nginx to host the frontend and forward connections to the backend.
