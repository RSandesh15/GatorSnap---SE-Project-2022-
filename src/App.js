
import Login from './pages/login';
import { Route , Switch} from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import SignUp from './pages/SignUp';
import ShowcaseImages from './Components/ShowcaseImages';
import SellerLogin from './pages/SellerLogin';
import SellerUploadPage from './pages/SellerUploadPage';
import UserLandingPage from './pages/userLandingPage';
import Checkout from './pages/Checkout';




 
function App() {
  console.log("Temp")
  return ( <div >
    <Switch>
      <Route exact path = "/"component = {Login} />
      <Route exact path = "/ShowcaseImages" component={ShowcaseImages} />
      <Route exact path = "/Dashboard" component={Dashboard} />
      <Route exact path = "/SignUp" component={SignUp} />
      <Route exact path = "/SellerLogin" component={SellerLogin} />
      <Route exact path = "/SellerUploadPage" component={SellerUploadPage} />
      <Route exact path = "/userLandingPage" component={UserLandingPage} />
      <Route exact path = "/Checkout" component={Checkout} />
      
    </Switch> 
    </div>
  );
  
}


export default App;
