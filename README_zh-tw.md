# Mail Server on GAE

## 摘要

本文將說明如何在 GAE 架設 API Mail Server。
GCP 上的 GCE 不得作為 Smtp Mail Server 使用，必須串接 gsuite 或第三方服務
但如果 GCP 內部需要使用寄信服務（如，定期寄信給稽核單位），串接第三方服務又顯得不太安全。
而 GAE 本身就有寄信服務，因此利用 GAE 架設 API Mail Server。

## 啟用 GAE

在 GCP 上建立新專案
1. 選擇 App Engine
2. 選擇程式語言，本文使用 Golang（其實選擇語言只會影響後面的 tutorial，如果打算略過 tutorial，選擇哪個都一樣）
3. 選擇部署的位置，目前 GAE 沒有台灣的點，因此選擇日本[1]；一旦選好 region 後，之後所有的 services 都是在這個 region，無法修改；如果有更換 region 的需求，直接建立新專案
4. 隨後會出現 tutorial，可依需求來決定是否繼續照著 tutorial 實作；本文選擇略過
5. 頁面重新整理後出現完整的 GAE 畫面

## 設定寄件者
在 GAE 頁面 → Settings → Email senders → ADD
注意，能當 sender

- 該專案的擁有者
- 已該專案 ID 為域名；假設此專案 ID 為 `hello-world-2018`
  可以使用 `mail@hello-world-2018.appspotmail.com` ，這個 email 不需要加入
- 詳細可以參考 [2]


## 取得 Mail Server Project Source Code

打開 Cloud Shell，如果無法開啟可以嘗試使用無痕模式
這個程式碼是由 golang 開發，先查詢 Cloud Shell 的 go path 在哪
```shell
$ echo $GOPATH
/home/<GCP_UESR>/gopath:/google/gopath
```

已當前 user 的 home 為主，建立一個 gopath
```shell
$ mkdir -p ~/gopath
```

取得 source code
```shell
$ go get -u github.com/CloudMile/gae_send_mail_api
```

出現下面訊息只是警告，可忽略
```shell
package github.com/CloudMile/gae_send_mail_api: no Go files in /home/<GCP_UESR>/gopath/src/github.com/CloudMile/gae_send_mail_api
```

取得此專案使用的 http controller lib
```shell
$ go get -u github.com/gorilla/mux
```

切到專案下
```shell
$ cd ~/gopath/src/github.com/CloudMile/gae_send_mail_api/
```

部署
```shell
$ make deploy PROJECT_ID='<YOUR_PROJECT_ID>' FROM='mail@<YOUR_PROJECT_ID>.appspotmail.com'
```

## 測試與使用
cURL with multipart/form-data
```shell
$ curl -X POST \
-F "to=to.mail@mile.cloud" \
-F "cc=cc1.mail@mile.cloud,cc2.mail@mile.cloud" \
-F "bcc=bcc1.mail@mile.cloud,bcc2.mail@mile.cloud" \
-F "subject=Send mail from GAE" \
-F "data=@./favicon.png" \
-F "body=upload file" \
"https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send"
```

cURL with application/x-www-form-urlencoded
```shell
$ curl -X POST -d 'to=to.mail@mile.cloud&subject=Send mail from GAE' "https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send"
```

cURL with application/json
```shell
$ curl -X POST -H 'Content-Type: application/json' -d '{"to": "to.mail@mile.cloud", "subject": "Send mail from GAE"}' "https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send"
```

使用 ruby
```shell
$ gem install rest-client
```

```ruby
require 'rest-client'
send_url = 'https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send'
file = File.open('./static/favicon.png')
params = {
  to: 'to.mail@mile.cloud',
  subject: 'Send Mail for Test',
  body: 'TESTTESTTESTTESTTESTTESTTEST',
  data: file
}
RestClient.post(send_url, params)
```

使用 python
```shell
$ pip install requests
$ pip install requests-toolbelt
```

```python
import requests
from requests_toolbelt.multipart.encoder import MultipartEncoder
send_url = 'https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send'
multipart_data = MultipartEncoder(
    fields={
            # a file upload field
            'data': ('favicon.png', open('./static/favicon.png', 'rb'), 'text/plain'),
            # plain text fields
            'to': 'to.mail@mile.cloud',
            'subject': 'Send Mail for Test',
            'body': 'TESTTESTTESTTESTTESTTESTTEST',
            }
    )
response = requests.post(send_url, data=multipart_data, headers={'Content-Type': multipart_data.content_type})
```

## Refs

- [1] https://cloud.google.com/about/locations/?region=asia-pacific#region
- [2] https://cloud.google.com/appengine/docs/standard/go/mail/#who_can_send_mail
