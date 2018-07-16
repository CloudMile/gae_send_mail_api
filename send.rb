require 'rest_client'

file = File.open('./favicon.png')

params = {
  to: 'to.mail@mile.cloud',
  subject: 'Send Mail for Test',
  body: 'TESTTESTTESTTESTTESTTESTTEST',
  data: file
}

RestClient.post('https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send', params)
