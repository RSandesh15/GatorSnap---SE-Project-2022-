import React, { useState, useEffect } from "react";
import axios from "axios";
import ImageList from "@mui/material/ImageList";
import ImageListItem from "@mui/material/ImageListItem";
import ImageListItemBar from "@mui/material/ImageListItemBar";
import IconButton from "@mui/material/IconButton";
import StarBorderIcon from "@mui/icons-material/StarBorder";
import Badge from "@material-ui/core/Badge";
import ShoppingCartIcon from "@material-ui/icons/ShoppingCart";
import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import ButtonGroup from "@material-ui/core/ButtonGroup";
import AddIcon from "@material-ui/icons/Add";
import RemoveIcon from "@material-ui/icons/Remove";
import "../App1.css";
import BasicModal from "../Components/BasicModal";
import { Link } from 'react-router-dom';

//Bootstrap
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.min.js";

function srcset(image, width, height, rows = 1, cols = 1) {
  return {
    src: `${image}?w=${width * cols}&h=${height * rows}&fit=crop&auto=format`,
    srcSet: `${image}?w=${width * cols}&h=${
      height * rows
    }&fit=crop&auto=format&dpr=2 2x`,
  };
}

export default function UserLandingPage() {
    
  // constructor(props) {
  //     super(props);
  //     this.state = { email: "jim@ufl.edu", id:"", authflag:1 };
  //     this.handleAddition = this.handleAddition.bind(this);
  //     this.handleChange = this.handleChange.bind(this);
  //     }
  // const email: "jim@ufl.edu"

  const handleAddition = (e, imageID) => {
    console.log("hi");
    e.preventDefault();
    fetch("http://localhost:8085/addToCart", {
      mode: "no-cors",
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-type": "application/json",
      },
      body: JSON.stringify({
        buyerEmailId: "aakanshtogani@ufl.edu",
        imageId: imageID,
      }),
    });

    setItemCount(Math.max(itemCount + 1));
  };
  const handleDeletion = (e, imageID) => {
    console.log("hi");
    e.preventDefault();
    fetch("http://localhost:8085/deleteFromCart", {
      mode: "no-cors",
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-type": "application/json",
      },
      body: JSON.stringify({
        buyerEmailId: "aakanshtogani@ufl.edu",
        imageId: imageID,
      }),
    });

    setItemCount(Math.max(itemCount - 1, 0));
  };
  //   handleChange(event) {
  //     this.setState({ email: event.state.email, id: event.state.id });
  //     }

  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
 
  const [fetchImages, setFetchedImageData] = useState([]);
  const obj = Object.entries(fetchImages);

  useEffect(() => {
    axios.get("http://localhost:8085/fetchImages").then((response) => {
      debugger;
      setFetchedImageData(response.data.data);
    });
  }, []);

  const [itemCount, setItemCount] = React.useState(0);
  const [imageID, setImageID] = React.useState();

  return (
    <div className="app">
      <nav
        class="navbar navbar-expand-lg bg-dark navbar-light d-none d-lg-block"
        id="templatemo_nav_top"
      >
        <div class="container text-light">
          <div class="w-100 d-flex justify-content-between">
            <div>
              <i class="fa fa-envelope mx-2"></i>
              <a
                class="navbar-sm-brand text-light text-decoration-none"
                href="#"
              >
                info@GatorSnaps.com
              </a>
              <i class="fa fa-phone mx-2"></i>
              <a
                class="navbar-sm-brand text-light text-decoration-none"
                href="#"
              >
                010-000-0000
              </a>
            </div>
            <div>
              <a class="text-light" href="#" target="_blank" rel="sponsored">
                <i class="fab fa-facebook-f fa-sm fa-fw me-2"></i>
              </a>
              <a class="text-light" href="#" target="_blank">
                <i class="fab fa-instagram fa-sm fa-fw me-2"></i>
              </a>
              <a class="text-light" href="#" target="_blank">
                <i class="fab fa-twitter fa-sm fa-fw me-2"></i>
              </a>
              <a class="text-light" href="#" target="_blank">
                <i class="fab fa-linkedin fa-sm fa-fw"></i>
              </a>
            </div>
          </div>
        </div>
      </nav>

      <nav class="navbar navbar-expand-lg navbar-light shadow">
        <div class="container d-flex justify-content-between align-items-center">
          <a
            class="navbar-brand text-success logo h1 align-self-center"
            href="#"
          >
            GatorSnaps!
          </a>

          <button
            class="navbar-toggler border-0"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#templatemo_main_nav"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="navbar-toggler-icon"></span>
          </button>

          <div
            class="align-self-center collapse navbar-collapse flex-fill  d-lg-flex justify-content-lg-between"
            id="templatemo_main_nav"
          >
            <div class="flex-fill">
              <ul class="nav navbar-nav d-flex justify-content-between mx-lg-auto">
                <li class="nav-item">
                  <a class="nav-link" href="#">
                    Home
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">
                    About
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">
                    Shop
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" href="#">
                    Contact
                  </a>
                </li>
              </ul>
            </div>
            <div class="navbar align-self-center d-flex">
              <div class="d-lg-none flex-sm-fill mt-3 mb-4 col-7 col-sm-auto pr-3">
                <div class="input-group">
                  <input
                    type="text"
                    class="form-control"
                    id="inputMobileSearch"
                    placeholder="Search ..."
                  />
                  <div class="input-group-text">
                    <i class="fa fa-fw fa-search"></i>
                  </div>
                </div>
              </div>
              <a
                class="nav-icon d-none d-lg-inline"
                href="#"
                data-bs-toggle="modal"
                data-bs-target="#templatemo_search"
              >
                <i class="fa fa-fw fa-search text-dark mr-2"></i>
              </a>
              <Badge color="secondary" badgeContent={itemCount}>
                
                  
                  <button>
                  <Link to = '/Checkout'>
                  <ShoppingCartIcon />
                      </Link>
                  </button>
                
              </Badge>
            </div>
          </div>
        </div>
      </nav>

      <div
        class="modal fade bg-white"
        id="templatemo_search"
        tabindex="-1"
        role="dialog"
        aria-labelledby="exampleModalLabel"
        aria-hidden="true"
      >
        <div class="modal-dialog modal-lg" role="document">
          <div class="w-100 pt-1 mb-5 text-right">
            <button
              type="button"
              class="btn-close"
              data-bs-dismiss="modal"
              aria-label="Close"
            ></button>
          </div>
          <form
            action=""
            method="get"
            class="modal-content modal-body border-0 p-0"
          >
            <div class="input-group mb-2">
              <input
                type="text"
                class="form-control"
                id="inputModalSearch"
                name="q"
                placeholder="Search ..."
              />
              <button
                type="submit"
                class="input-group-text bg-success text-light"
              >
                <i class="fa fa-fw fa-search text-white"></i>
              </button>
            </div>
          </form>
        </div>
      </div>
      <div
        className="background-size"
        style={{ display: "flex", justifyContent: "center" }}
      >
        <ImageList
          sx={{
            width: 500,
            height: 450,
            // Promote the list into its own layer in Chrome. This costs memory, but helps keeping high FPS.
            transform: "translateZ(0)",
          }}
          rowHeight={200}
          gap={1}
        >
          {fetchImages.map((item) => {
            const cols = 8;
            const rows = 2;

            return (
              <Box>
                <ImageListItem key={item.wImageUrl} cols={cols} rows={rows}>
                  <img
                    {...srcset(item.wImageUrl, 125, 100, rows, cols)}
                    alt={item.title}
                    loading="lazy"
                  />
                  <ImageListItemBar
                    sx={{
                      background:
                        "linear-gradient(to bottom, rgba(0,0,0,0.7) 0%, " +
                        "rgba(0,0,0,0.3) 70%, rgba(0,0,0,0) 100%)",
                    }}
                    title={item.title}
                    position="top"
                    actionIcon={
                      <IconButton
                        sx={{ color: "white" }}
                        aria-label={`star ${item.title}`}
                      >
                        <StarBorderIcon />
                      </IconButton>
                    }
                    actionPosition="left"
                  />
                </ImageListItem>
                <ButtonGroup>
                  <Button onClick={handleDeletion}>
                    {" "}
                    <RemoveIcon fontSize="small" />
                  </Button>
                  <Button onClick={(e) => handleAddition(e, item.imageId)}>
                    {" "}
                    <AddIcon fontSize="small" />
                  </Button>
                  <Button onClick={handleOpen}>Details</Button>

                  {open && <BasicModal imageId={item.imageId} />}
                </ButtonGroup>
              </Box>
            );
          })}
        </ImageList>
      </div>

      <section class="container py-5">
        <div class="row text-center pt-3">
          <div class="col-lg-6 m-auto">
            <h1 class="h1">Categories of The Month</h1>
            <p>
              Lorem Ipsum is simply dummy text of the printing and typesetting
              industry. Lorem Ipsum has been the industry's standard dummy text
              ever since the 1500s, when an unknown printer took a galley of
              type and scrambled it to make a type specimen book.
            </p>
          </div>
        </div>
        <div class="row">
          <div class="col-12 col-md-4 p-5 mt-3">
            <a href="#">
              <img
                src="https://www.cianellistudios.com/images/abstract-art/abstract-art-celebration.jpg"
                class="rounded-circle img-fluid border"
              />
            </a>
            <h5 class="text-center mt-3 mb-3">Nature</h5>
            <p class="text-center">
              <a class="btn btn-success">Go Shop</a>
            </p>
          </div>
          <div class="col-12 col-md-4 p-5 mt-3">
            <a href="#">
              <img
                src="https://www.cianellistudios.com/images/abstract-art/abstract-art-canvas-prints-me5.jpg"
                class="rounded-circle img-fluid border"
              />
            </a>
            <h2 class="h5 text-center mt-3 mb-3">Abstract</h2>
            <p class="text-center">
              <a class="btn btn-success">Go Shop</a>
            </p>
          </div>
          <div class="col-12 col-md-4 p-5 mt-3">
            <a href="#">
              <img
                src="https://www.cianellistudios.com/images/abstract-art/abstract-art-celebration.jpg"
                class="rounded-circle img-fluid border"
              />
            </a>
            <h2 class="h5 text-center mt-3 mb-3">Art</h2>
            <p class="text-center">
              <a class="btn btn-success">Go Shop</a>
            </p>
          </div>
        </div>
      </section>

      <section class="bg-light">
        <div class="container py-5">
          <div class="row text-center py-3">
            <div class="col-lg-6 m-auto">
              <h1 class="h1">Featured Product</h1>
              <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry. Lorem Ipsum has been the industry's standard dummy
                text ever since the 1500s, when an unknown printer took a galley
                of type and scrambled it to make a type specimen book.
              </p>
            </div>
          </div>
          <div class="row">
            <div class="col-12 col-md-4 mb-4">
              <div class="card h-100">
                <a href="#">
                  <img
                    src="https://www.cianellistudios.com/images/abstract-art/abstract-art-canvas-prints-ref1.jpg"
                    class="card-img-top"
                    alt="..."
                  />
                </a>
                <div class="card-body">
                  <ul class="list-unstyled d-flex justify-content-between">
                    <li>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-muted fa fa-star"></i>
                      <i class="text-muted fa fa-star"></i>
                    </li>
                    <li class="text-muted text-right">$240.00</li>
                  </ul>
                  <a href="#" class="h2 text-decoration-none text-dark">
                    Lorem Ipsum
                  </a>
                  <p class="card-text">
                    Lorem ipsum dolor sit amet, consectetur adipisicing elit.
                    Sunt in culpa qui officia deserunt.
                  </p>
                  <p class="text-muted">Reviews (24)</p>
                </div>
              </div>
            </div>
            <div class="col-12 col-md-4 mb-4">
              <div class="card h-100">
                <a href="#">
                  <img
                    src="https://www.cianellistudios.com/images/abstract-art/abstract-art-canvas-prints-romance1.jpg"
                    class="card-img-top"
                    alt="..."
                  />
                </a>
                <div class="card-body">
                  <ul class="list-unstyled d-flex justify-content-between">
                    <li>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-muted fa fa-star"></i>
                      <i class="text-muted fa fa-star"></i>
                    </li>
                    <li class="text-muted text-right">$480.00</li>
                  </ul>
                  <a href="#" class="h2 text-decoration-none text-dark">
                    Lorem Ipsum
                  </a>
                  <p class="card-text">
                    Lorem Ipsum Lorem Ipsum Lorem IpsumLorem Ipsum
                  </p>
                  <p class="text-muted">Reviews (48)</p>
                </div>
              </div>
            </div>
            <div class="col-12 col-md-4 mb-4">
              <div class="card h-100">
                <a href="#">
                  <img
                    src="https://www.cianellistudios.com/images/abstract-art/abstract-art-canvas-prints-sg1b.jpg"
                    class="card-img-top"
                    alt="..."
                  />
                </a>
                <div class="card-body">
                  <ul class="list-unstyled d-flex justify-content-between">
                    <li>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                      <i class="text-warning fa fa-star"></i>
                    </li>
                    <li class="text-muted text-right">$360.00</li>
                  </ul>
                  <a href="#" class="h2 text-decoration-none text-dark">
                    Lorem Ipsum
                  </a>
                  <p class="card-text">
                    Lorem Ipsum Lorem Ipsum Lorem IpsumLorem Ipsum Lorem Ipsum
                    Lorem Ipsum.
                  </p>
                  <p class="text-muted">Reviews (74)</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <footer class="bg-dark" id="tempaltemo_footer">
        <div class="container">
          <div class="row">
            <div class="col-md-4 pt-5">
              <h2 class="h2 text-success border-bottom pb-3 border-light logo"></h2>
              <ul class="list-unstyled text-light footer-link-list">
                <li>
                  <i class="fas fa-map-marker-alt fa-fw"></i>
                  India
                </li>
                <li>
                  <i class="fa fa-phone fa-fw"></i>
                  <a class="text-decoration-none" href="#">
                    000-000-0000
                  </a>
                </li>
                <li>
                  <i class="fa fa-envelope fa-fw"></i>
                  <a class="text-decoration-none" href="#">
                    info@company.com
                  </a>
                </li>
              </ul>
            </div>

            <div class="col-md-4 pt-5">
              <h2 class="h2 text-light border-bottom pb-3 border-light">
                Products
              </h2>
              <ul class="list-unstyled text-light footer-link-list">
                <li>
                  <a class="text-decoration-none" href="#">
                    Luxury
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    Sports
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    Nature
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    Tourist Places
                  </a>
                </li>
              </ul>
            </div>

            <div class="col-md-4 pt-5">
              <h2 class="h2 text-light border-bottom pb-3 border-light">
                Further Info
              </h2>
              <ul class="list-unstyled text-light footer-link-list">
                <li>
                  <a class="text-decoration-none" href="#">
                    Home
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    About Us
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    Locations
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    FAQs
                  </a>
                </li>
                <li>
                  <a class="text-decoration-none" href="#">
                    Contact
                  </a>
                </li>
              </ul>
            </div>
          </div>

          <div class="row text-light mb-4">
            <div class="col-12 mb-3">
              <div class="w-100 my-3 border-top border-light"></div>
            </div>
            <div class="col-auto me-auto">
              <ul class="list-inline text-left footer-icons">
                <li class="list-inline-item border border-light rounded-circle text-center">
                  <a
                    class="text-light text-decoration-none"
                    target="_blank"
                    href="#"
                  >
                    <i class="fab fa-facebook-f fa-lg fa-fw"></i>
                  </a>
                </li>
                <li class="list-inline-item border border-light rounded-circle text-center">
                  <a
                    class="text-light text-decoration-none"
                    target="_blank"
                    href="#"
                  >
                    <i class="fab fa-instagram fa-lg fa-fw"></i>
                  </a>
                </li>
                <li class="list-inline-item border border-light rounded-circle text-center">
                  <a
                    class="text-light text-decoration-none"
                    target="_blank"
                    href="#"
                  >
                    <i class="fab fa-twitter fa-lg fa-fw"></i>
                  </a>
                </li>
                <li class="list-inline-item border border-light rounded-circle text-center">
                  <a
                    class="text-light text-decoration-none"
                    target="_blank"
                    href="#"
                  >
                    <i class="fab fa-linkedin fa-lg fa-fw"></i>
                  </a>
                </li>
              </ul>
            </div>
            <div class="col-auto">
              <label class="sr-only" for="subscribeEmail">
                Email address
              </label>
              <div class="input-group mb-2">
                <input
                  type="text"
                  class="form-control bg-dark border-light"
                  id="subscribeEmail"
                  placeholder="Email address"
                />
                <div class="input-group-text btn-success text-light">
                  Subscribe
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="w-100 bg-black py-3">
          <div class="container">
            <div class="row pt-2">
              <div class="col-12">
                <p class="text-left text-light">
                  Copyright &copy; 2021 Gator Snaps | Designed by{" "}
                  <a rel="sponsored" href="#" target="_blank">
                    Gators
                  </a>
                </p>
              </div>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}

// {
//   <ImageList sx={{ width: 500, height: 450 }}>
//   {fetchImages.map((item) => (
//     <ImageListItem key={item.imageId}>
//       <img
//         src={`${item.wImageUrl}?w=248&fit=crop&auto=format`}
//         alt={item.title}
//         loading="lazy"
//       />
//       <ImageListItemBar
//         title={item.title}
//         subtitle={<span>by: {item.price}</span>}
//         position="below"
//       />
//     </ImageListItem>
//   ))}
// </ImageList>
// }
