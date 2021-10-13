import requests
from requests_toolbelt.multipart.encoder import MultipartEncoder

send_url = 'https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send'

multipart_data = MultipartEncoder(
    fields={
            # a file upload field
            'data': ('favicon.png', open('./static/favicon.png', 'rb'), 'text/plain'),
            # plain text fields
            'to': 'to.mail@any-mail.com',
            'subject': 'Send Mail for Test from python',
            'body': 'TESTTESTTESTTESTTESTTESTTEST',
            }
    )

response = requests.post(send_url, data=multipart_data, headers={'Content-Type': multipart_data.content_type})
