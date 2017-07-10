# HTML Template populated from MongoDB (Go)
Just a project more for myself to refer back to when working on future GoLang projects. Renders HTML template from MongoDB/Heroku queries/data. Service pulls data from **rails-mongodb-heroku** production ( https://rails-mongodb-heroku.herokuapp.com/people )

# Setup
***To Run:***

*Set Environment Variables via Commands or in Bash File*

export MONGODB_URI="mongodb://restofurl"

export MONGODB_DB="mongodb_database_name"

export PORT="8080"

*In local testing, make sure mongoDB is running by executing:*
  ```
  mongod
  ```
*In a new terminal window:*

  ```
  git clone https://github.com/brianhodges/golang-mongodb-htmltemplate
  cd golang-mongodb-htmltemplate
  go run main.go
  ```
  
*Then simply navigate in your browser to:* 
 
 **DEVELOPMENT**
 
    http://localhost:8080/
 
 **PRODUCTION (HEROKU)**
 
    https://golang-mongodb-htmltemplate.herokuapp.com/
    
You can query by last_name and/or first_name by appending params to url. Case insensitive. 

Ex: https://golang-mongodb-htmltemplate.herokuapp.com?last_name=Bending
