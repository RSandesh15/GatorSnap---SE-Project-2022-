# GatorSnap 
### By: Pulin Soni, Rishab Parmar, Aakansh Togani, Sandesh Ramesh

## Sprint-1 Functionality

### Backend: 

We have worked on and developed the following APIs :

- /checkoutAndProcessPayment: In this API, after adding the products to buy to the cart, the user hits checkout. Furthermore, whenever the buyer puts in their payment information and hits the pay button, a call is made to the Stripe payment gateway that returns with the payment intent. This payment intent when passed to the front end, the callback code is next called to invoke /emailProduct.

- /emailProduct:  

  - After receiving the client secret from the above API call, this client secret is then confirmed to be executed from the front-end side which states whether the payment was successful or not 

  - Moreover, if it is successful, this API is invoked. By doing so, we then start emailing the bought unwatermarked images one by one.  

  - We first fetch the original unwatermarked image from the S3 cloud bucket and then save it locally on the server to mail 

  - We then create a mail template, attach the above downloaded image to the mail and mail it to the registered email id of the buyer 

  - After this point, we then update or remove the cart content regarding the transaction and consequently, we also add the recently made transaction to the previously bought schema. 

- Setup and integrated the Stripe payment gateway for checkout and purchases. The Stripe payment gateway returns a client secret after the payment information has been entered by the buyer and once confirmed, we return the paymentIntent Id for further processing

- Enabled mailing services for mailing the respective image to the respective buyer using gomail API for optimal development

- Wrote test cases for the above mentioned APIs and more, both correct and incorrect test cases are working accordingly and are giving the desired output



### Front-end:
- For Sprint 3, we integrated 3 important API addtoCart, DeletefromCart and FetchCartInfo. All the API mentioned used are of POST method.
- Another set of pages and components have been created for checkout , payment and review information specifically. The amount of items are added and a total is displayed to the user.
- Login authentication is been also implemented using an API call which response accordingly to the token returned.


### Video Walkthrough

Here is a walkthrough of what was achieved on the backend and database side for sprint 3. 
<img src='Gifs/Sprint3_recording.gif' title='Backend' width='' />

Here is a walkthrough of what was achieved on the frontend side for sprint 3. 
<img src='Gifs/FrontEnd_Sprint3.gif' title='Frontend' width='' />

### Cypress Testing
Conducted cypress tests on additional functionalities for multiple pages. This testing included - 1. Page redirection 2. Field authentication 3. Button responses 4. Field validation
