import Login from "./login";
import SignUp from "./SignUp";





describe('Login test cases', () => {

    it("renders Login Page successfully", () => {
        const wrapper = shallow(
            <Login />
        );
        expect(wrapper).toMatchSnapshot();
    });

    it("simulate the click event on Button", () => {
        const wrapper = shallow(<Login />);
        expect(wrapper.find('Link').prop('to')).to.be.equal('/Login');
    });
})


describe('Login test cases', () => {

    it("renders Login Page successfully", () => {
        const wrapper = shallow(
            <SignUp />
        );
        expect(wrapper).toMatchSnapshot();
    });

    it("simulate the click event on Button", () => {
        const wrapper = shallow(<SignUp />);
        expect(wrapper.find('Link').prop('to')).to.be.equal('/SignUp');
    });
})

it('should show error when entered - Password', ()=>{
    wrapper.find('#password').simulate('change', {target: {value: 'password'}});
    expect(wrapper.find("#password").props().error).toBe(
        true);
    expect(wrapper.find("#password").props().helperText).toBe(
        'Wrong password format.');
  });

  it('should show error when entered - email', ()=>{
    wrapper.find('#email').simulate('change', {target: {value: ''}});
    expect(wrapper.find("#email").props().error).toBe(
        true);
    expect(wrapper.find("#email").props().helperText).toBe(
        'Wrong email format.');
  });