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
