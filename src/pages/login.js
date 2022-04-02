import React from "react";
import{
Button,
TextField,
Grid,
Paper,
AppBar,
Typography,
Toolbar,

} from "@material-ui/core";
import { Link } from 'react-router-dom';





class Login extends React.Component {
    
constructor(props) {
super(props);
this.state = { username: "", password:"", authflag:1 };
this.handleChange = this.handleChange.bind(this);
this.handleSubmit = this.handleSubmit.bind(this);
}
handleChange(event) {
this.setState({ username: event.state.username, password: event.state.password });
}

handleSubmit = (e) => {
    e.preventDefault();
    fetch ("https://reqres.in/api/login", {
       method: 'POST',
       headers:{
           'Accept' : 'application/json',
           'Content-type' : 'application/json'
       },
       body: JSON.stringify({
         email: this.state.username,
         password: this.state.password
      }),
  })
    .then((response) => response.json())
    .then((result) => {
      if(result.token === "QpwL5tke4Pnpja7X4"){
        alert("You are logged in.");
        console.log(result.token)
       } else {
           alert("Please check your login information.");
           console.log(result)
       }
    });
  }
  

render() {
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
<br />
<br />


<center>
<Typography variant="h6">Buyers</Typography>
</center>
<Grid container spacing={-5} justify="center" direction="row">
<Grid item>
<Grid
container
direction="column"
justify="center"
spacing={2}
className="login-form"
>
<Paper
variant="elevation"
elevation={5}
className="login-background"
>
<Grid item>
<Typography component="h1" variant="h5">
Sign in
</Typography>
</Grid>
<Grid item>
<form onSubmit={this.handleSubmit}>
<Grid container direction="column" spacing={2}>
<Grid item>
<TextField
type="email"
placeholder="Email"
fullWidth
name="username"
variant="outlined"
value={this.state.username}
onChange={(event) =>
this.setState({
[event.target.name]: event.target.value,
})
}
required
autoFocus
/>
</Grid>
<Grid item>
<TextField
type="password"
placeholder="Password"
fullWidth
name="password"
variant="outlined"
value={this.state.password}
onChange={(event) =>
this.setState({
[event.target.name]: event.target.value,
})
}
required
/>
</Grid>
<Grid item>
<Button
variant="contained"
color="primary"
type="submit"
className="button-block"
>
Submit
</Button>
</Grid>
</Grid>
</form>
</Grid>
<Grid item>
<Link to = '/SignUp'>
Forgot Password?
</Link>
</Grid>
<Grid item>
<Link to = '/SignUp'>
Create New Account!
</Link>
</Grid>
</Paper>
</Grid>
</Grid>
</Grid>

        
                        
</div>

);
}
}
export default Login;