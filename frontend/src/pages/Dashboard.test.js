import * as React from 'react';
import { styled, createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import MuiDrawer from '@mui/material/Drawer';
import Box from '@mui/material/Box';
import MuiAppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import List from '@mui/material/List';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import Badge from '@mui/material/Badge';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Link from '@mui/material/Link';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import NotificationsIcon from '@mui/icons-material/Notifications';
import { mainListItems, secondaryListItems } from '../Components/listItems';
import Chart from '../Components/Chart';
import Deposits from '../Components/Balances';
import Cart from '../Components/Cart';
import Balances from '../Components/Balances';
import Avatars from '../Components/Avatars';
import Dashboard from './Dashboard';


describe('Dashboard test cases', () => {

    it("renders Dashboard Page successfully", () => {
        const wrapper = shallow(
            <Dashboard />
        );
        expect(wrapper).toMatchSnapshot();
    });

    it("simulate the click event on Button", () => {
        const wrapper = shallow(<Dashboard />);
        expect(wrapper.find('Link').prop('to')).to.be.equal('/Dashboard');
    });
})