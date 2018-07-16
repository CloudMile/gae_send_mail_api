import requests
from requests_toolbelt.multipart.encoder import MultipartEncoder

multipart_data = MultipartEncoder(
    fields={
            # a file upload field
            'data': ('favicon.png', open('favicon.png', 'rb'), 'text/plain'),
            # plain text fields
            'to': 'to.mail@mile.cloud',
            'subject': 'Send Mail for Test',
            'body': 'TESTTESTTESTTESTTESTTESTTEST',
            }
    )

response = requests.post('https://mail-dot-<YOUR_PROJECT_ID>.appspot.com/send', data=multipart_data, headers={'Content-Type': multipart_data.content_type})
