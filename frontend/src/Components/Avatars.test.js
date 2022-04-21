import Avatars from './Avatars';
import { Link } from 'react-router';
import SellerUploadPage from '../pages/SellerUploadPage';


describe('Avatars Component test cases', () => {

    it("renders Avatars Component successfully", () => {
        const wrapper = shallow(
            <SellerUploadPage />
        );
        expect(wrapper).toMatchSnapshot();
    });

    it("simulate the click event on Button", () => {
        const onButtonClick = sinon.spy();
        const wrapper = shallow(<SellerUploadPage />);
        expect(wrapper.find('Link').prop('to')).to.be.equal('/SellerUploadPage');
    });
})