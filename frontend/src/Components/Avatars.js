import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Stack from '@mui/material/Stack';
import { CenterFocusStrong } from '@mui/icons-material';
import {
  Button,
  TextField,
  Grid,
  Paper,
  AppBar,
  Typography,
  Toolbar,
  
  } from "@material-ui/core";
  import { Link } from 'react-router-dom';

export default function Avatars() {
  return (
    <center>
      <Avatar
        alt="Suraj Mishra"
        src='https://images.unsplash.com/photo-1589118949245-7d38baf380d6'
        sx={{ width: 80, height: 80 }}
      />
      <div>
         <h1>Name : Suraj Mishra</h1> 
        <Link to = '/SellerUploadPage'>
        ADD IMAGES!
        </Link>
        
      </div>
      </center>
      
     
  );
}
