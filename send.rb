require 'rest-client'

send_url = 'https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send'

file = File.open('./static/favicon.png')

params = {
  to: 'to.mail@any-mail.com',
  subject: 'Send Mail for Test from ruby',
  body: 'TESTTESTTESTTESTTESTTESTTEST',
  data: file
}

RestClient.post(send_url, params)
