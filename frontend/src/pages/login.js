import React, { useState, useEffect } from "react";
import axios from "axios";
import Back from "../Components/background.jpg";
import {
  Button,
  TextField,
  Grid,
  Paper,
  AppBar,
  Typography,
  Toolbar,
} from "@material-ui/core";
import { Link } from "react-router-dom";

class Login extends React.Component {
  constructor(props) {
    super(props);
    this.state = { username: "", password: "", authflag: 1 };
    this.handleChange = this.handleChange.bind(this);
    this.handleLogin = this.handleLogin.bind(this);
  }
  handleChange(event) {
    this.setState({
      username: event.state.username,
      password: event.state.password,
    });
  }

  handleLogin = (e) => {
   
    e.preventDefault();
    console.log("In submit")
    if (typeof window !== 'undefined') {
      window.location.href = "http://localhost:8085/google/login";
 }
  };

  render() {
    return (
      <div
        style={{
          backgroundImage: `url(${Back})`,
          height: "100vh",
          marginTop: "30px",
          fontSize: "50px",
          backgroundSize: "cover",
          backgroundRepeat: "no-repeat",
        }}
      >
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
          <Typography variant="h4">Buyers</Typography>
        </center>
        <br />
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
                  <Typography component="h1" variant="h5" align="center">
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
                          
                          className="button-block"
                        >
                          Submit
                        </Button>
                      </Grid>
                    </Grid>
                    
                  </form>
                  <br/>
                  <Typography component="h4" variant="h9" align="center">
                      OR
                    </Typography>
                    <br />
                    <Typography component="h1" variant="h5" align="center">
                      Sign In With Google
                    </Typography>
                    <br />
                    <Button
                      onClick = {this.handleLogin}
                      variant="contained"
                      color="primary"
                      className="button-block"
                    >
                      Google Sign In
                    </Button>
                    <Grid>
                      <br />
                    </Grid>
                </Grid>
                <br />

                <Grid>
                  <Link to="/SignUp">Forgot Password?</Link>
                </Grid>
                <Grid item>
                  <Link to="/SignUp">Create New Account!</Link>
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
