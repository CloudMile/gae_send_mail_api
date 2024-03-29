# Mail Server on GAE

## Abstract

Because of GCE couldn't be a mail smtp server[1].
For workaround, setup a mail server with API mode on GAE.

## Setup GAE

Create a new project on GCP
1. Choose App Engine.
2. Choose language for development, we using `Go`.
3. Choose deploy location[2].
4. You can skip the tutorial.
5. Refresh the page, and you can see the GAE main page.

## Setup Sender
On GAE → Settings → Email senders → ADD
Note, there are some restrictions for sender

- Project Owner
- Make project ID to be a domain name; for example, the project is `hello-world-2018`
  you can use `mail@hello-world-2018.appspotmail.com`, this mail DON'T need to add into Email senders
- More info[3]


## Get Mail Server Project Source Code

Open Cloud Shell on GCP console
check $GOPATH
```shell
$ echo $GOPATH
/home/<GCP_UESR>/gopath:/google/gopath
```

if $GOPATH not exist, create a new one
```shell
$ mkdir -p ~/gopath
```

get source code
```shell
$ go get -u github.com/CloudMile/gae_send_mail_api
```

this is a warning, skip it
```shell
package github.com/CloudMile/gae_send_mail_api: no Go files in /home/<GCP_UESR>/gopath/src/github.com/CloudMile/gae_send_mail_api
```

get http controller lib
```shell
$ go get -u github.com/gorilla/mux
```

cd to project
```shell
$ cd ~/gopath/src/github.com/CloudMile/gae_send_mail_api/
```

set up app.yaml
```shell
$ vim main/app.yaml
```
you can set up CUSTOM_TOKEN for enable auth check
if you enable, you need to add `curl -H 'Custom-Token: <YOUR_TOEKN>' `

deploy
```shell
$ make deploy PROJECT_ID='<YOUR_PROJECT_ID>' FROM='mail@<YOUR_PROJECT_ID>.appspotmail.com'
```

## Test or Use
cURL with multipart/form-data
```shell
$ curl -X POST \
-F "to=to.mail@any-mail.com" \
-F "cc=cc1.mail@any-mail.com,cc2.mail@any-mail.com" \
-F "bcc=bcc1.mail@any-mail.com,bcc2.mail@any-mail.com" \
-F "subject=Send mail from GAE" \
-F "data=@./favicon.png" \
-F "body=upload file" \
"https://mail-dot-<YOUR_PROJECT_ID>.<REGION_ID>.r.appspot.com/send"
```
The `<REGION_ID>` is an abbreviated code that Google assigns based on the region you select when you create your app[4].


cURL with application/x-www-form-urlencoded
```shell
$ curl -X POST -d 'to=to.mail@any-mail.com&subject=Send mail from GAE' "https://mail-dot-<YOUR_PROJECT_ID>.<REGION_ID>.r.appspot.com/send"
```

cURL with application/json
```shell
$ curl -X POST -H 'Content-Type: application/json' -d '{"to": "to.mail@any-mail.com", "subject": "Send mail from GAE"}' "https://mail-dot-<YOUR_PROJECT_ID>.<REGION_ID>.r.appspot.com/send"
```

ruby
```shell
$ gem install rest-client
```

```ruby
require 'rest-client'
send_url = 'https://mail-dot-<YOUR_PROJECT_ID>.<REGION_ID>.r.appspot.com/send'
file = File.open('./static/favicon.png')
params = {
  to: 'to.mail@any-mail.com',
  subject: 'Send Mail for Test',
  body: 'TESTTESTTESTTESTTESTTESTTEST',
  data: file
}
RestClient.post(send_url, params)
```

python
```shell
$ pip install requests
$ pip install requests-toolbelt
```

```python
import requests
from requests_toolbelt.multipart.encoder import MultipartEncoder
send_url = 'https://mail-dot-<YOUR_PROJECT_ID>.<REGION_ID>.r.appspot.com/send'
multipart_data = MultipartEncoder(
    fields={
            # a file upload field
            'data': ('favicon.png', open('./static/favicon.png', 'rb'), 'text/plain'),
            # plain text fields
            'to': 'to.mail@any-mail.com',
            'subject': 'Send Mail for Test',
            'body': 'TESTTESTTESTTESTTESTTESTTEST',
            }
    )
response = requests.post(send_url, data=multipart_data, headers={'Content-Type': multipart_data.content_type})
```

## Refs
- [1] https://cloud.google.com/compute/docs/tutorials/sending-mail/
- [2] https://cloud.google.com/about/locations/?region=asia-pacific#region
- [3] https://cloud.google.com/appengine/docs/standard/go/mail/#who_can_send_mail
- [4] https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed#gcloud_1
