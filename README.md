# google-proxy-heroku

Run a google reverse proxy using golang on Heroku.

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Installation 
-  Deploy Button

Use the [Deploy to Heroku](https://heroku.com/deploy) button above to create a
copy of the app, then Set `UPSTREAM_SERVER` config vars.

-  Git Deploy

Steps：

1. git clone https://github.com/fromxt/google-proxy-heroku.git

2. cd google-proxy-heroku ,then modify `main.go`
 ```
    upstream_server := "https://www.google.com"
```
3. heroku login -i

4. heroku create apps: <APP name>
  
5. git add -A

6. git commit -m "init"

7. git push heroku master

Note:
If a VPS not available, you can sign up [goormIDE](https://www.goorm.io/) account to insatll the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli),which is a powerful cloud IDE service.  
