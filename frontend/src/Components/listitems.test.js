import * as React from 'react';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import ListSubheader from '@mui/material/ListSubheader';
import DashboardIcon from '@mui/icons-material/Dashboard';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import PeopleIcon from '@mui/icons-material/People';
import BarChartIcon from '@mui/icons-material/BarChart';
import LayersIcon from '@mui/icons-material/Layers';
import AssignmentIcon from '@mui/icons-material/Assignment';
import Dashboard from '../pages/Dashboard';




it('should render proper icon on show and hide', async () => {
    const component = shallow(<Dashboard />);
    checkOpenState(component, false);
    checkCloseState(component, true);
  
    toggleState(component);
  
    checkOpenState(component, true);
    checkCloseState(component, false);
})