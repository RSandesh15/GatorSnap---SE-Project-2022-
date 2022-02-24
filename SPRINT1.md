# GatorSnap 
### By: Pulin Soni, Rishab Parmar, Aakansh Togani, Sandesh Ramesh

## Sprint-1 Functionality

### Backend: 
- For Sprint 1, we started by creating the project structure and boilerplate code by following the industry standards. We then proceeded forward to design and create the database schemas for the features planned for Sprint 1.   

- Moreover, we then fed dummy data to these databases that will mimic the original data that will be received with the front end. We then created the following APIs:

- /fetchImages: API to fetch all the images(product catalogs) and corresponding data from the SQLite database.
Method: GET
- /fetchGenreCategories: API to fetch all the various genres available from the database to display for seller upload page

- These APIs when called from the front end will help the data to be delivered to the UI when requested. Furthermore, we have partially worked on the /uploadSellerImage API which is a POST method-based API. This API will take the input from the seller upload form and upload the data in our database through this API. Moreover, the code to upload the image to the cloud storage has been coded but is yet to be tested.

- Worked on Marshalling and UnMarshalling of JSON data.

- Developed several data structs to store data in sqlite.
### Front-end:
- For Sprint 1, we started by creating the project structure and boilerplate code by following the industry standards. We then proceeded forward to design and create the application features planned for Sprint 1.

- Starting with the login page, we created functionalities such as : Checking for '@' in the email field, alerting if the field are left blank, displaying a pop up if incorrect credential are typed. Only logs in for user using specified mock.json file.

- Mock sign up page which would register all the required information for making a customer login. There are two different login pages for seller and buyer. 

- Seller Dashboard displays the seller profile, his total balances, and a list of purchases previously made. Also incorporates dashboard's app bar with seamless transitions. From the seller dashboard, the seller can re-direct to the picture upload page. 

- Seller Upload Page contains fields such as picture title, description, price, and genre. The details are submitted once the seller clicks the submit button. 

- User Landing Page: After loggin in successfully, the users can view the images uploaded by sellers in a grid manner. The images have been unmarshalled from the json file. 


### Video Walkthrough

Here is a walkthrough of what was achieved on the backend and database side for sprint 1. 
<img src='Gifs/backend_Sprint1.gif' title='Backend' width='' />

Here is a walkthrough of what was achieved on the frontend side for sprint 1. 
<img src='Gifs/Frontend_Sprint1.gif' title='Frontend' width='' />