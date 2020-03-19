# Corona App


## Goals
- decrease number of contacts outside of their regular social circle
- help older people without access to internet -> let children, nephews add requests for other people
- decrease workload for emergency hotline by providing assitance for self diagnosis using the app 
- decrease workload on hospitals, emergency sites by signaling to the user where resources, beds, oxygen... are available

## Protoype

https://pr.to/W864FN/


TODO: Fina a name https://namelix.com/
TODO: Logo https://looka.com/

## Backend

### TODO
- salt passwords
- API spec 

Database Schema: https://dbdiagram.io/d/5e714fd54495b02c3b88612f
Free PostgreSQL DB on ElephantSQL (20 MB limit) (credentials on request) -> Move to managed digitalocean DB


Run Backend: go build && ./coronapp_server


Send Request to create a user:
curl -X POST \                                                                         ✘ 130
  -H 'Content-Type: application/json' \
  -d '{"firstName":"Joe", "lastName":"Smith", "email":"asd@def.de", "password":"pass", "mobile":"+4917122"}' \
  localhost:1323/users/create



## Frontend
### TODO


vue native, flutter, swiftui...


## Links

- all testing sites in Berlin https://www.berlin.de/en/news/coronavirus/6100254-6098215-coronavirus-examination-centres-in-berli.en.html

