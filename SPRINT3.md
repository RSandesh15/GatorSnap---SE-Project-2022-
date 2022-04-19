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
Create a new React project to get started. Optionally add TypeScript
1. npm create react-app cypress-test-react --template typescript
You also need @cypress/react, which is the primary way to mount and interact with components. 
1. npm add cypress @cypress/react @cypress/webpack-dev-server --dev

### Uniting with Jest
If you are new to React, we recommend using Create React App. It is ready to use and ships with Jest. You will only need to add react-test-renderer for rendering snapshots. If you have an existing application you'll need to install a few packages to make everything work well together. We are using the babel-jest package and the react babel preset to transform our code inside of the test environment. Use the following for setup:
1. yarn add --dev react-test-renderer
2. yarn add --dev jest babel-jest @babel/preset-env @babel/preset-react react-test-renderer

'it' or 'test' You would pass a function to this method, and the test runner would execute that function as a block of tests. The first rule is that any files found in any directory with the name __test__ are considered a test. If you put a JavaScript file in one of these folders, Jest will try to run it when you call Jest, for better or for worse.

To run the test cases, we need to use the command:
1. npm run test OR
2. npm test
