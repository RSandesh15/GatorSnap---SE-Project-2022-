# Gator SnapStore 
## Members: 
### Backend:
- Pulin Soni
- Rishab Parmar
### Frontend:
- Aakansh Togani
- Sandesh Ramesh

### About Gator SnapStore
A web application for UF students to upload their artwork/photographs for potential buyers to browse and purchase according to their liking. We have integrated Google authentication to make the login process as secure as possible. The images uploaded are displayed with a watermark on the webpages but the un-watermarked images are send to the buyer's email ID after the payment is processed. We have integrated Stripe as our payment portal. Other functionalities such as addding and deleting images from Cart, searching of images present in the database are also implemented.

### Future Work
We can think about scaling our application using noSQL databases such as mongoDB to attract more userbase. We can also integrate new payment methods by taking advantage of blockchain technology and deploying each image on the blockcahin with a smart contract to act as an nft. We can deploy this website for students and take their review as part of our beta testing phase.

### Demo video

https://user-images.githubusercontent.com/93216515/164368572-b8ed818e-e7e1-4509-ae06-25dd1bb46f5b.mp4

### Backend Testing video
<img src='Gifs/backend_testing.gif' title='Demo' width='' />

### Cypress Testing video
Conducted cypress tests on additional functionalities for multiple pages. This testing included - 1. Page redirection 2. Field authentication 3. Button responses 4. Field validation
Create a new React project to get started. Optionally add TypeScript
1. npm create react-app cypress-test-react --template typescript
You also need @cypress/react, which is the primary way to mount and interact with components. 
1. npm add cypress @cypress/react @cypress/webpack-dev-server --dev

<img src='Gifs/cypress_testing.gif' title='Demo' width='' />

### Link to API Documentation
[API documentation](https://uflorida-my.sharepoint.com/:w:/g/personal/parmar_rishab_ufl_edu/EVL6ZXFHf2dLpko6o5w2DwQBPgCj16-c7Ur--bWRycFzUQ?e=1C7GaN)

### Link to project board
[Sprint 4](https://github.com/RSandesh15/GatorSnap---SE-Project-2022-/projects/3)
[Sprint 3](https://github.com/RSandesh15/GatorSnap---SE-Project-2022-/projects/2)
[Sprint 1/2](https://github.com/RSandesh15/GatorSnap---SE-Project-2022-/projects/1)

### Sprint 4 Deliverables
### Backend:
We have developed the following set of APIs for Sprint 4: 

- /google/login: This API is called to invoke the Google authentication for users to log in using their google login details. 

- /google/callback: This API is called to handle the callback function for google authentication. It is used to generate a token and to set a cookie for aiding in future API calls. 

- /logout: This API is called to end the session of a user that was logged into the system. 

- /fetchSellerInfo: When the seller logs in to the system, this API is invoked. This API gives all the previous sales that were conducted by the seller, stating information such as the seller and buyer information, the price at which the product was bought, the precise time of the transaction, etc. 

- Clearing all the dead and test code, code refinement and adding comments wherever necessary 

- Wrote test cases for the above mentioned APIs and more, both correct and incorrect test cases are working accordingly and are giving the desired output 

### Frontend:
- Completed Login Integration with a an API call, google/login.
- Imported the cookie and used it for all the API calls.
- Integration of Stripe (Bypassing CORS policy), using stripe created paymentIntentID. Used this stripe created paymentIntentID for enhancing security on the payment portal which is STRIPE.
- Calling the emailProduct API and succefully sending email to the registered user.
- Created more test cases for every webpage to validate functionality.

## Thank You!! Hope you enjoyed our project idea and implemetation.
