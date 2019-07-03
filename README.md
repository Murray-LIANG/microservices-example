# Simple Example of Microservices

A simple movie system with users and their booked movies.

## General Data Flow

For each service, the data flow is like below:

 ```console
 requests / response
          ^
          |
          v
  +---------------+
  |    Routers    |
  +---------------+
          ^
          |
          v
        JSON
          ^
          |
          v
  +---------------+
  |  Controllers  |
  +---------------+
          ^
          |
          v
       Models
          ^
          |
          v
  +---------------+
  |     Data      |
  +---------------+
          ^
          |
          v
        BSON
          ^
          |
          v
  +---------------+
  |      mgo      |
  +---------------+
```

## Deployment

Use `docker-compose` to deploy services.

### Steps

1. `docker-compose build`
2. `docker-compose up`
3. Add below line to file `/etc/hosts` of localhost.
```bash
# file /etc/hosts

# original content
# ...

127.0.0.1 db movies.local users.local bookings.local monitor.local
```

## Usage

Use the Chrome extension `Restlet` or `Postman` to send requests to services.

### Create
`POST` to http://users.local/users with body:
```json
{
  "data": {
    "name": "ryan",
    "lastname": "liang"
  }
}
``` 

### Get
`GET` http://users.local/users

The response returned:
```json
{
    "data": [
        {
            "id": "5d1c4bff2b5f1900012f0d32",
            "name": "ryan",
            "lastname": "liang"
        }
    ]
}
``` 

### Delete
`DELETE` http://users.local/users/5d1c48652b5f1900012f0d31