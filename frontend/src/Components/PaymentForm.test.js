import * as React from 'react';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';


it('should show error when entered - Card Name', ()=>{
    wrapper.find('#cardName').simulate('change', {target: {value: 'cardName'}});
    expect(wrapper.find("#cardName").props().error).toBe(
        true);
    expect(wrapper.find("#cardName").props().helperText).toBe(
        'Wrong card format.');
  });

  it('should show error when entered - Card Number', ()=>{
    wrapper.find('#cardNumber').simulate('change', {target: {value: 'cardNumber'}});
    expect(wrapper.find("#cardNumber").props().error).toBe(
        true);
    expect(wrapper.find("#cardNumber").props().helperText).toBe(
        'Wrong card number format.');
  });

  it('should show error when entered - expiry date', ()=>{
    wrapper.find('#expDate').simulate('change', {target: {value: 'expDate'}});
    expect(wrapper.find("#expDate").props().error).toBe(
        true);
    expect(wrapper.find("#expDate").props().helperText).toBe(
        'Wrong expiry date format.');
  });

  it('should show error when entered - CVV', ()=>{
    wrapper.find('#cvv').simulate('change', {target: {value: 'cvv'}});
    expect(wrapper.find("#cvv").props().error).toBe(
        true);
    expect(wrapper.find("#cvv").props().helperText).toBe(
        'Wrong CVV format.');
  });
