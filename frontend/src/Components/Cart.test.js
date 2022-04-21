import * as React from 'react';
import Link from '@mui/material/Link';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Title from './Title';
import Cart from './Cart';

describe('Avatars Component test cases', () => {

    it("renders Cart Component successfully", () => {
        const wrapper = shallow(
            <Cart />
        );
        expect(wrapper).toMatchSnapshot();
    });

    it("simulate the click event on Button", () => {
        const wrapper = shallow(<Cart />);
        expect(wrapper.find('Link').prop('to')).to.be.equal('/Cart');
    });
})