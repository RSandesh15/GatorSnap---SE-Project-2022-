import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Stack from '@mui/material/Stack';
import { CenterFocusStrong } from '@mui/icons-material';


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
         <h4>Buyer</h4>
      </div>
      </center>
      
     
  );
}
