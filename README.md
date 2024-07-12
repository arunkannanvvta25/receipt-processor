# Receipt Processor

A Receipt Processor Service which generates a id for the given json and calculates the points based on the rules.

## Implementation:

* Layered Architecture (aka N-Tier Architecture) for API project
* Strategy Design Pattern for Points Calculation
* Abstraction for DataAccessLayer and ServiceLayer (ReceiptService)
* Map for storing receipts (O(1) retrieval)
* Added unit tests for critical methods
* Dockerfile for containerization
* UUID for ID generation

## Tech Stack

* Go
* Docker

## Steps for Running using Docker Container

create a Docker Image  

`docker build -f Dockerfile -t my-receipt-processor .`

Run the image in a container  

`docker run --name receipt-processor-container -d -p 5001:8080 my-receipt-processor`

Run `docker ps` to see the container that got created.

Endpoints(Application will run in port 5001):

http://localhost:5001/receipts/process  
http://localhost:5001/receipts/{id}/points

if 5001 is already in use , then use a different port.


## Steps for Running in local

Run the following  
`go get .`  
`go run .`

To run the unit tests run `go test ./tests -v`

Endpoints(Application will run in port 8080):

http://localhost:8080/receipts/process  
http://localhost:8080/receipts/{id}/points

---

### API Endpoint: Process Receipts

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.


## Endpoint: Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.

## Sample Input JSON:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```