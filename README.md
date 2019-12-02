# nepal-it-pool
A simple way to list tech workforce in Nepal.

Simple api made on GO and ElasticSearch with json data.

Goto https://nepal-it-pool-3pjwxgvmna-uc.a.run.app/ for listings     

## Filters

All filters are query string param i.e https://nepal-it-pool-3pjwxgvmna-uc.a.run.app/?name=ujwal&skills=php,java

So here some filters 

* name (e.g name=ujwal)
* designation (e.g designation=Software Engineer) 
* actively_job_searching (e.g actively_job_searching=false)
* max_exp (e.g max_exp=5)
* min_exp (e.g min_exp=2)
* skills (e.g skills=java,php)
* is_intern (e.g is_intern=false)
* address (e.g address=Kathmandu,Nepal)
* state (e.g state=3)

For sorting use
* sort=key only `"experience", "name", "current_company","is_intern","actively_job_searching"` keys are available
* sort_type=asc or desc , by default asc is used

 
 
 ## How to list yourself ?
 * Fork this repo
 * Create the file at `app/crowdSourceData` with your github username like `app/crowdSourceData/ujwaldhakal.json`. Add your info in json format like
 To list your info here please edit 
 
 ```
 {
    "name": "Ujwal Dhakal ",
    "email": "kevin.ujwal@gmail.com",
    "designation" : "Software Engineer",
    "experience" : "5",
    "skills" : ["php","node","go"],
     "current_company" : "Pagevamp",
     "is_intern" : false,
    "actively_job_searching": false,
     "address" : "Kathmandu,Nepal",
     "state" : 3,
    "github_url" : "https://github.com/ujwaldhakal" ,
    "linkedin_url" : ""
     
 }
```
* Push and make a PR
   
## Removal Procedure
If you want yourself to get removed from the list please remove your listing from `app/crowdSourceData` and send us PR .
Note-: Please send with the same github account that is listed on list
     
  
## Contribution Guidelines
Go and elasticsearch has been used for this api.

## Setup
* Clone this repo
* `docker-compose up`
* Go `localhost:5000` you should see data or if you want to have your custom virtual host `- VIRTUAL_HOST=itpool.pv` change this line of docker file as well as in your `etc/hosts` too.
* Send PR with clear details 


Any suggestions to make search better please create issue with your detailed solutions :) 

## Disclaimer 
We wont be responsible for any misuse of data and we believe that the data you provided are for public general use :)  

[![Generic badge](https://github.com/ujwaldhakal/nepal-it-pool/workflows/Build%20and%20Push/badge.svg)](https://github.com/ujwaldhakal/nepal-it-pool/actions)
