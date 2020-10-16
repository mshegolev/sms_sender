# sms_sender
sms_sender example 

1. Register on site https://sms.ru
2. Get token. Open page https://sms.ru/?panel=api and get api_id: example "89CC5CED-6B96-FDAD-E65B-5BC31E1CEBBB"
3. Put token to ```api_id``` variable.
4. Put phone number to ```phoneNumber``` variable like 89111234567.
5. Put message to ```message``` variable like "hello world!".
6. ```go run sms_sender.go```
