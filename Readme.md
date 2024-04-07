# Auth Service

This is a Golang authentication service built using the Gin framework

## Features

- User sign in by email verification code
- Token-based authentication
- Google oath included
- Metric system allowed
- API documentation (check path in your browser `localhost:5000/swagger/index.html`)

## Prerequisites

- Go 1.22 or higher
- Gin framework
- PostgreSQL database
- Redis cache v.6.0.16 or higher 

## Get started


1. Clone the repository:

	```shell
	git clone <repository_url>
	```

2. Prepare special api keys with your credentials:
	- Google Oauth Client ID
	- Google Oauth Client Secret
	- Google Email App Password

3. Create the config file with `.json` resolution and set up service parameters:
	```json
	{
		"jwt" : {
			"secret" : "secret_key",	
			"accessMinutes" : 15,		
			"refreshDays" : 60			
		},
		"google": {
			"clientID": "google_client_id",				
			"clientSecret": "google_client_secret",		
			"redirectURL": "google_oauth_redirect_url",	
			"scopes": [
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile"
			]
		},
		"conductor" : {
			"senderEmail": "google_email_sender",	
			"senderPass": "google_email_password",	
			"senderPort": 587,
			"mailTitle": "Authorization code from Service!",
			"mailCodeLength" : 6,
			"mailCodeDuration" : 15,
			"mailTemplatePath": "static/verificationCodeMail.html", 
			"tokenSecret": "verification_secret_key", 				
			"codeRefreshDelay:": 2,
			"codeEnterAttempts": 10
		}
	}
	```
   - `accessMinutes` - How many minutes access token is active
   - `refreshDays` - How many days refresh token is active
   - `clientID` - Google Oauth client id
   - `clientSecret` - Google Oauth client secret
   - `redirectURL` - After auth the page will be redirected by this url
   - `scopes` - Types of an access for the Google user info
   - `senderEmail` - Verification code sender email from Gmail
   - `senderPass` - Verification code sender app password (not from google account)

5. Build the docker-compose images:
	```shell
	docker-compose build
	```

6. Run the docker-compose:
	```shell
	docker-compose up -d
	```

## Configuration

The application can be configured using the following environment variables:
- `PORT` : 			service port
- `CONFIG_PATH` : 	path to config file
- `DB_NAME` : 		postgres db name
- `DB_USER` : 		postgres username
- `DB_PASS` : 		postgres password
- `DB_ADDR` : 		postgres address
- `DB_SSL` : 		postgres ssl (not allowed)
- `CACHE_ADDR` : 	redis cache address
- `CACHE_DB` : 		redis cache db

## License

This project is licensed under the [MIT License](LICENSE).