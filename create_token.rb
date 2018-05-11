require 'jwt'
payload = { user: 'passwords' }

# IMPORTANT: set nil as password parameter
token = JWT.encode payload, nil, 'none'
# eyJhbGciOiJub25lIn0.eyJkYXRhIjoidGVzdCJ9.
puts token



hmac_secret = 'My Secret'
token = JWT.encode payload, hmac_secret, 'HS256'
token = JWT.encode payload, hmac_secret, algorithm='HS256', header_fields={typ: "JWT"}
# eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjoidGVzdCJ9.pNIWIL34Jo13LViZAJACzK6Yf0qnvT_BuwOxiMCPE-Y
puts token


# Set password to nil and validation to false otherwise this won't work
decoded_token = JWT.decode token, nil, false

decoded_token = JWT.decode token, hmac_secret, true, { algorithm: 'HS256' }
# Array
# [
#   {"data"=>"test"}, # payload
#   {"alg"=>"none"} # header
# ]
puts decoded_token