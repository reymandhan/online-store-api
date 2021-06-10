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
- There is no inventory quantity validation on the system that cause overselling.
<br/>
<br/>
## Proposed Solution
For the sake of simplicity and clarity, there are several assumptions for the scope of proposed solution. 
- No separate inventory system, item quantity stored on item table
- No user/customer table, to simplify only put customer name, address etc on order table. Also cartId will be use to track shopping cart instead of user id
- No authentication and authorization process
- No Item category or item review data
- Shopping cart data will be deleted after success payment (data will be on order table)

### Changes 
- Update item quantity after each success payment
- Check item availability before each step (add to cart, checkout, and payment) to prevent overselling
- Return additional info to customer on how many item already on other customer shopping cart (if any) compare to latest item quantity to give the customer sense of urgency to immediately checkout and pay before stock runs out. (Especially for items whose stock is low)

<br/>

## Prerequisite to run
- Docker

