
import{
Button,
TextField,
Grid,
Paper,
AppBar,
Typography,
Toolbar,

} from "@material-ui/core";
import { render } from "@testing-library/react";
import Checkout2 from "../Components/Checkout2";

export default function Checkout() {
       render()
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