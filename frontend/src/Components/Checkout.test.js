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
import Checkout from "../pages/Checkout";
import Checkout2 from "./Checkout2";



Enzyme.configure({ adapter: new Adapter() })

describe('Test Case For Checkout2', () => {
  it('should render button', () => {
    const wrapper = shallow(<Checkout2 />)
    const buttonElement  = wrapper.find('#Back');
    expect(buttonElement).toHaveLength(1);
    expect(buttonElement.text()).toEqual('Back');
  })
})