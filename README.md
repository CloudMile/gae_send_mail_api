# API Server for Send Mail on GAE
Sample API server for send mail on GAE

## Enable GAE service
[GCP console](https://console.cloud.google.com/)

## Setup Config

## How to Deploy
Change your `from mail` into `main/app.yaml`
```sh
$ vim ./main/app.yaml
```
chang `<YOUR_GAE_MAIL_SENDER>` you need

And then check this [url](https://cloud.google.com/appengine/docs/standard/python/getting-started/deploying-the-application) to learn how to deploy

## How to Use
You need to using POST Form
```sh
$ curl -X POST \
-F "to=to.mail@mile.cloud" \
-F "cc=cc1.mail@mile.cloud,cc2.mail@mile.cloud" \
-F "bcc=bcc1.mail@mile.cloud,bcc2.mail@mile.cloud" \
-F "subject=Send mail from GAE" \
-F "data=@./favicon.png" \
-F "body=upload file" \
https://[HTTPS_GAE_URL]/send
```
