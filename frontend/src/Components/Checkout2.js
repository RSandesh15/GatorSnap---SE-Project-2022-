import React, { useState, useEffect } from "react";
import CssBaseline from "@mui/material/CssBaseline";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Toolbar from "@mui/material/Toolbar";
import Paper from "@mui/material/Paper";
import Stepper from "@mui/material/Stepper";
import Step from "@mui/material/Step";
import StepLabel from "@mui/material/StepLabel";
import Button from "@mui/material/Button";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import AddressForm from "./AddressForm";
import PaymentForm from "./PaymentForm";
import Review from "./Review";

import axios from "axios";
import { Mode } from "@mui/icons-material";

function Copyright() {
  return (
    <Typography variant="body2" color="text.secondary" align="center">
      {"Copyright © "}
      <Link color="inherit" href="https://mui.com/">
        GatorSnaps!
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const steps = ["Billing address", "Review your order", "Payment"];

function getStepContent(step) {
  switch (step) {
    case 0:
      return <AddressForm />;
    case 1:
      return <Review />;
    case 2:
      return <PaymentForm />;
    default:
      throw new Error("Unknown step");
  }
}

const theme = createTheme();

export default function Checkout2() {
  const [activeStep, setActiveStep] = React.useState(0);
  const [clientSecret, setClientSecret] = useState([]);
  // debugger;
  console.log(clientSecret);
  const handleNext = () => {
    setActiveStep(activeStep + 1);
  };

  const handleBack = () => {
    setActiveStep(activeStep - 1);
  };
  // const handlePlaceOrder = (e) => {
  //   axios
  //     .post(
  //       "http://10.3.6.194:8085/checkoutAndProcessPayment",
  //       {mode: "no-cors"},
  //       {
  //         body: JSON.stringify({
  //           //token: "",
  //           buyerEmailId: "parmar.rishab@gmail.com",
  //           //paymentIntentId: "pi_3Kq1JvE2RN3PJKON1kdsZxAb",
  //         }),
  //       },
  //       {
  //         headers: {
  //           Accept: "*/*",
  //           "Content-type": "application/json",
  //           "access-control-allow-origin": "*",
  //           "Access-Control-Allow-Headers": "*",
  //           "access-control-allow-methods": "*"
            
  //         },
  //       }
  //     )
      
  // };

  const handlePlaceOrder = (e) => {
    setActiveStep(activeStep + 1);
    console.log("in here");
    e.preventDefault();
    let url = "http://localhost:8085/emailProduct";
    fetch(url, {
      mode: "no-cors",
      method: "POST",
  headers: {
    Accept: "application/json",
    "Content-type": "application/json",
      },
      body: JSON.stringify({

          token: "",
          buyerEmailId: "mehuljhaver@ufl.edu",
          paymentIntentId: "pi_3KqQeXE2RN3PJKON1PgoiGRH"
      }),
    }).then(function (response) {
            console.log(response.data);
           });
  };

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <AppBar
        position="absolute"
        color="default"
        elevation={0}
        sx={{
          position: "relative",
          borderBottom: (t) => `1px solid ${t.palette.divider}`,
        }}
      ></AppBar>
      <Container component="main" maxWidth="sm" sx={{ mb: 4 }}>
        <Paper
          variant="outlined"
          sx={{ my: { xs: 3, md: 6 }, p: { xs: 2, md: 3 } }}
        >
          <Typography component="h1" variant="h4" align="center">
            Checkout
          </Typography>
          <Stepper activeStep={activeStep} sx={{ pt: 3, pb: 5 }}>
            {steps.map((label) => (
              <Step key={label}>
                <StepLabel>{label}</StepLabel>
              </Step>
            ))}
          </Stepper>
          <React.Fragment>
            {activeStep === steps.length ? (
              <React.Fragment>
                <Typography variant="h5" gutterBottom>
                  Thank you for your order.
                </Typography>
                <Typography variant="subtitle1">
                  Your order number is #2001539. We have emailed your order
                  confirmation, and will send you an update when your order has
                  shipped.
                </Typography>
              </React.Fragment>
            ) : (
              <React.Fragment>
                {getStepContent(activeStep)}
                <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
                  {activeStep !== 0 && (
                    <Button onClick={handleBack} sx={{ mt: 3, ml: 1 }}>
                      Back
                    </Button>
                  )}

                  <Button
                    variant="contained"
                    onClick={
                      activeStep === steps.length - 1
                        ? handlePlaceOrder
                        : handleNext
                    }
                    sx={{ mt: 3, ml: 1 }}
                  >
                    {activeStep === steps.length - 1 ? "Place Order" : "Next"}
                  </Button>
                </Box>
              </React.Fragment>
            )}
          </React.Fragment>
        </Paper>
        <Copyright />
      </Container>
    </ThemeProvider>
  );
}
