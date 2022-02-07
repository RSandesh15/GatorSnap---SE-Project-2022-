
import Login from './pages/login';
import { Route , Switch} from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import SignUp from './pages/SignUp';
import ShowcaseImages from './Components/ShowcaseImages';
import SellerLogin from './pages/SellerLogin';
import SellerUploadPage from './pages/SellerUploadPage';
 
function App() {
  return ( <div >
    <Switch>
      <Route exact path = "/"component = {Login} />
      <Route exact path = "/ShowcaseImages" component={ShowcaseImages} />
      <Route exact path = "/Dashboard" component={Dashboard} />
      <Route exact path = "/SignUp" component={SignUp} />
      <Route exact path = "/SellerLogin" component={SellerLogin} />
      <Route exact path = "/SellerUploadPage" component={SellerUploadPage} />
    </Switch> 
    </div>
  );
  
}


export default App;
