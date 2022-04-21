import React, { useState, useEffect } from "react";
import axios from "axios";
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';
import React from 'react'
import Enzyme, { shallow } from 'enzyme'
import Adapter from 'enzyme-adapter-react-16'
import BasicModal from "./BasicModal";


Enzyme.configure({ adapter: new Adapter() })

describe('Test Case For BasicModal', () => {
  it('should render button', () => {
    const wrapper = shallow(<BasicModal />)
    const buttonElement  = wrapper.find('#Close');
    expect(buttonElement).toHaveLength(1);
    expect(buttonElement.text()).toEqual('Close');
  })
})