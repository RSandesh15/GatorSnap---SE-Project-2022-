import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Axios from 'axios';


it('should show error when entered - First Name', ()=>{
    wrapper.find('#firstName').simulate('change', {target: {value: 'firstName'}});
    expect(wrapper.find("#firstName").props().error).toBe(
        true);
    expect(wrapper.find("#firstName").props().helperText).toBe(
        'Wrong Name format.');
  });

  it('should show error when entered - Last Name', ()=>{
    wrapper.find('#lastName').simulate('change', {target: {value: 'lastName'}});
    expect(wrapper.find("#lastName").props().error).toBe(
        true);
    expect(wrapper.find("#lastName").props().helperText).toBe(
        'Wrong Name format.');
  });

  it('should show error when entered - Email Address', ()=>{
    wrapper.find('#email').simulate('change', {target: {value: 'email'}});
    expect(wrapper.find("#email").props().error).toBe(
        true);
    expect(wrapper.find("#email").props().helperText).toBe(
        'Wrong email format.');
  });

  it('should show error when entered - Password', ()=>{
    wrapper.find('#password').simulate('change', {target: {value: 'password'}});
    expect(wrapper.find("#password").props().error).toBe(
        true);
    expect(wrapper.find("#password").props().helperText).toBe(
        'Wrong password format.');
  });

  