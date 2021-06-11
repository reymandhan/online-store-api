# Task 1 : Online Store

## Problem
We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following 
facts:
- Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from our Customer Service department stating that their orders have been canceled due to stock unavailability.
- These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major 
discounts to promote our store.

After checking in with our Customer Service and Order Processing departments, we received the following additional facts:
- Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity.
- The misreported items are those that performed very well on our 12.12 event.
- Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders
<br/>
<br/>

## Analysis
Based on the facts above, 
- Customers might have not provided with real time inventory quantity.
- There is a lot of concurrent transaction happen for the same data (items that performed well on the event) might cause the stock validation doesn't work correctly
<br/>
<br/>

## Scope and assumption
For the sake of simplicity and clarity, there are several assumptions for the scope of proposed solution. 
- No separate inventory system, item quantity stored on item table
- No user/customer table, to simplify only put customer name, address etc on order table. Also cartId and username will be use to track shopping cart
- No authentication and authorization process
- No Item category
- Shopping cart data will be deleted after success payment (data will be on order table)
- Simplify process to sequential process (cannot back to previous step):
    1. Add to cart (or remove from cart)
    2. Checkout (create order)
    3. Pay (update order status and item quantity)
<br/>
<br/>

## Proposed solution 
To prevent overselling, need to serialise concurrent changes to stock availability using DB row lock.
Payment process implemented as following:
- Retrieve Order and Order Item data
- Create Transaction for row lock
- Iterate Order Item
    - Retrieve item to get stock quantity with lock row (to block concurrect read and update for this item)
    - If stock not available (or not enough) return error, stop process
    - If stock available, update stock quantity reduced with paid item quantity, then continue process
- Update order status
- Clean user cart and cart item
- If process finished succesfully, then commit the transaction and lock row will be released
- If any of the step return error, rollback transaction and lock row will be released as well.
  

<br/>

## How to run
### Docker
- Open terminal/console
- Go to root directory of this code
- Run the app with `docker-compose up`

### Go
- make sure Postgres DB installed and running on local
- go version 1.16
- Open terminal/console
- Go to root directory of this code
- Copy `.env.example` and rename it `.env`
- Set your local Postgres credential on newly created `.env` file
- Download dependency with `go mod tidy`
- Run the app with `go run .`

Every time app run it will run migration (only once) and seed initial data.
to make sure app run and connected to DB, open browser access `localhost:8080/api/v1/health`, check the response.

## Test
Seed data will create product, cart and cart item.
Functional test will execute checkout and payment for those initial data.<br/>
There are 5 user that will pay concurrently, One of product will run out of stock.
One (or more) user will be failed to pay because of stock unavailable.

To run test, after run the app, run `go test`.
Re-run the test might have different result (depend on which user payment processed first).
To Re-run the test, re-run app first to refresh the data then run `go test` again

<br/>

Also postman collection ("Online Store.postman_collection.json") provided for all endpoint on this app 


