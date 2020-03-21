# Corozone


## Goals
- decrease number of contacts outside of their regular social circle
- help older people without access to internet -> let children, nephews add requests for other people
- decrease workload for emergency hotline by providing assitance for self diagnosis using the app 
- decrease workload on hospitals, emergency sites by signaling to the user where resources, beds, oxygen... 
are available and when they do not need to visit a doctor or facility
- educate layman to collect test samples and send them to a testing facility (?)
- allow testing facilities to report results to the user via the app and forward the data to gesundheitsamt

## Outside help

- contacts to doctors to evaluate symptoms when a user does not need to call emergency hotline
- contact to epidemiologist to make sure we don't increase contacts by helping neighbours
- contacts to medical facilities to let them provide their utilization rate
- contacts to international organisation (RKI, WHO...). Apple does not allow any third party corona related apps, needs to be vetted. 
## Protoype

Screenshot: https://postimg.cc/zVYnLRT8

https://pr.to/W864FN/


* TODO: Find a name https://namelix.com/
* TODO: Logo https://looka.com/

## Backend

### TODO
- salt passwords
- API spec 

Database Schema: https://dbdiagram.io/d/5e714fd54495b02c3b88612f
Free PostgreSQL DB on ElephantSQL (20 MB limit) (credentials on request) -> Move to managed digitalocean DB (15 Euro/Month)


Run Backend: go build && ./coronapp_server


Example request to create a user. This will be the link between Frontend and Backend:
curl -X POST \                                                                  
  -H 'Content-Type: application/json' \
  -d '{"firstName":"Joe", "lastName":"Smith", "email":"asd@def.de", "password":"pass", "mobile":"+4917122"}' \
  localhost:1323/users/create



## Frontend
### TODO


vue native, flutter, swiftui...


## Non-Tech Links

- all testing sites in Berlin (-> DB) https://www.berlin.de/en/news/coronavirus/6100254-6098215-coronavirus-examination-centres-in-berli.en.html
- https://covapp.charite.de/

## Tech Links

- Go tutorials
  - https://www.sohamkamani.com/blog/golang/sql-transactions/
  - https://www.restapiexample.com/golang-tutorial/creating-golang-api-echo-framework-postgresql/
  - https://www.youtube.com/watch?v=9VEJyPFz7WY&list=PLFmONUGpIk0YwlJMZOo21a9Q1juVrk4YY&index=3

## Kooperationen
- Bluetooth Low Energy id tracking #wirvsvirus https://github.com/InfectionChain
- Bluetooth tracking Singapore https://tracetogether.de.uptodown.com/android
- Links zu allen Telegram Gruppen für Nachbarschaftshilfe (evtl mit einem Bot Gesuche posten https://listling.org/lists/pwfjfkpjmesjjinm/solidarische-nachbarschaftshilfe)
