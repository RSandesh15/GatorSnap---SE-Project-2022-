import Typography from "@mui/material/Typography";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemText from "@mui/material/ListItemText";
import Grid from "@mui/material/Grid";
import React, { useState, useEffect } from "react";
import axios from "axios";

const products = [
  {
    name: "Product 1",
    desc: "A nice thing",
    price: "$9.99",
  },
  {
    name: "Product 2",
    desc: "Another thing",
    price: "$3.45",
  },
  {
    name: "Product 3",
    desc: "Something else",
    price: "$6.51",
  },
  {
    name: "Product 4",
    desc: "Best thing of all",
    price: "$14.11",
  },
  { name: "Shipping", desc: "", price: "Free" },
];

const addresses = ["1 MUI Drive", "Reactville", "Anytown", "99999", "USA"];
const payments = [
  { name: "Card type", detail: "Visa" },
  { name: "Card holder", detail: "Mr John Smith" },
  { name: "Card number", detail: "xxxx-xxxx-xxxx-1234" },
  { name: "Expiry date", detail: "04/2024" },
];

export default function Review() {
  const [fetchCart, setFetchedCartData] = useState([]);
  // const [total, addTotal] = useState(0);
  let total = 0;

  useEffect(() => {
    axios
      .get(`http://localhost:8085/fetchCartInfo`)
      .then((response) => {
        setFetchedCartData(response.data.data);
        console.log(response.data.data);
      });
  }, []);
  return (
    <React.Fragment>
      <Typography variant="h6" gutterBottom>
        Order summary
      </Typography>
      <List disablePadding>
        {fetchCart.map((product, i) => {
          // addTotal(total + product.price);
          total = total + product.price;
          return (
            <ListItem key={i} sx={{ py: 1, px: 0 }}>
              <ListItemText primary={product.title} secondary={product.desc}/>
              <ListItemText primary={product.imageId}/>
              <img src={product.wImageUrl} height={100} width={100} />
              <Typography variant="body2">{product.price}</Typography>
            </ListItem>
          );
        })}
        

        <ListItem sx={{ py: 1, px: 0 }}>
          <ListItemText primary="Total" />
          <Typography variant="subtitle1" sx={{ fontWeight: 700 }}>
            ${total}
          </Typography>
        </ListItem>
      </List>
      <Grid container spacing={2}>
      </Grid>
    </React.Fragment>
  );
}
