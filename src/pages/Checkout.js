import React, { useState, useEffect } from "react";
import axios from "axios";
import{
Button,
TextField,
Grid,
Paper,
AppBar,
Typography,
Toolbar,

} from "@material-ui/core";
import Checkout2 from "../Components/Checkout2";

export default function Checkout() {
    const [fetchCart, setFetchedCartData] = useState([]);
  useEffect(() => {
    axios.get(`http://localhost:8085/fetchCartInfo/aakansh.togani@ufl.edu`).then((response) => {
      
      setFetchedCartData(response.data.data);
       console.log(response.data.data);
    });
  }, []);
        return (
          
        <div>
        <AppBar position="static" alignitems="center" color="primary">
        <Toolbar>
        <Grid container justify="center" wrap="wrap">
        <Grid item>
        <Typography variant="h6">GatorSnaps!</Typography>
        </Grid>
        </Grid>
        </Toolbar>
        </AppBar>
        <Checkout2/>
        </div>
        
        
    );
}