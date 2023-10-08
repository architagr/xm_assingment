# xm_assingment

this code had 2 microservice
1. auth
2. company

## Auth microservice 
this service just has 2 routes 
a. GET: auth/healthCheck
  this is to chek if service is running 
b. POST : auth/getToken
  this request expects request body in below format 
  {
    "username":"",
    "password": ""
}
this will return token which will be used in the company microservice

## Company Microservice
this service has 5 routes 
a. GET: company/healthCheck
  this is to chek if service is running
b. GET: company/:id
  this return data of company if the company with id exists
c. POST: company/
  this adds a new company and expects a barer token and request body as below
  {
    "name": "Archit company",
    "numberOfEmployee": 1,
    "registered": true,
    "type": 1
  }
d. PATCH: company/:id
  this updates company and expects a barer token and request body as below
  {
    "name": "Archit company",
    "numberOfEmployee": 1,
    "registered": true,
    "type": 1
  }
e. 
d. DELETE: company/:id
  this deletes company and expects a barer token and request body as below


note: for now I have used a MongoDb and have added the connection to my mongodb in the config.json 
