import { render, screen } from '@testing-library/react';
import App from './App';
import React from 'react';
import { configure, shallow, mount } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import TextField from './AddressForm';
import TextField from '@material-ui/core/TextField';


test('renders learn react link', () => {
  expect(true).toBeTruthy();
});


configure({adapter: new Adapter()});

describe('<TextField />', ()=> {
  let shallow;

  beforeAll(() => {
    shallow = createShallow();
  });
  let wrapper;
  beforeEach(()=>{
    wrapper =  shallow(<TextField />);
  });

  it('should render one <TextField /> element.', ()=>{
    expect(wrapper.find(TextField)).toHaveLength(1);
  });


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

  it('should show error when entered - Address 1', ()=>{
    wrapper.find('#address1').simulate('change', {target: {value: 'address1'}});
    expect(wrapper.find("#address1").props().error).toBe(
        true);
    expect(wrapper.find("#address1").props().helperText).toBe(
        'Wrong Address1 format.');
  });

  it('should show error when entered - Address 2', ()=>{
    wrapper.find('#address2').simulate('change', {target: {value: 'address2'}});
    expect(wrapper.find("#address2").props().error).toBe(
        true);
    expect(wrapper.find("#address2").props().helperText).toBe(
        'Wrong Address2 format.');
  });

  it('should show error when entered - City', ()=>{
    wrapper.find('#city').simulate('change', {target: {value: 'city'}});
    expect(wrapper.find("#city").props().error).toBe(
        true);
    expect(wrapper.find("#city").props().helperText).toBe(
        'Wrong city format.');
  });

  it('should show error when entered - State', ()=>{
    wrapper.find('#state').simulate('change', {target: {value: 'state'}});
    expect(wrapper.find("#state").props().error).toBe(
        true);
    expect(wrapper.find("#state").props().helperText).toBe(
        'Wrong state format.');
  });

  it('should show error when entered - Country', ()=>{
    wrapper.find('#country').simulate('change', {target: {value: 'country'}});
    expect(wrapper.find("#country").props().error).toBe(
        true);
    expect(wrapper.find("#country").props().helperText).toBe(
        'Wrong Country format.');
  });
})
