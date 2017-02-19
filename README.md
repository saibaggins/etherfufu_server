# etherfufu_server

A simple Golang based REST service.

## Development

```
docker build -t etherfufu-server .
```

Once the build is created run the following command with the appropriate values for ENV variables :

```
 docker run -it --rm -p 3000:3000 \
 -e ACTIVE_ENV=staging \
 -e AWS_ACCESS_KEY_ID=something \
 -e AWS_SECRET_ACCESS_KEY=something_else \
 etherfufu-server
```

Add `-d` as an option for running in daemon mode.

## Deployment

Generate archive from your current source for deployment.

```
git pull origin master
git reset --hard HEAD
git archive --format=zip HEAD > /some/path/etherfufu-server_(git rev-parse HEAD).zip

```

Login to https://console.aws.amazon.com/elasticbeanstalk/ and upload the archive generated for deployment.