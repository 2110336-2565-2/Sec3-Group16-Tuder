// import css
import './App.css';

// import Router for link website
import { Fragment } from 'react';

// import components
import Navbar from './components/Navbar.js';
import Home from './pages/Home.js'
// import Card from './pages/components/Card';
import SignIn from './pages/SignIn.js';
import SignUp from './pages/SignUp.js';
import Courses from './pages/Courses';
import Report from './pages/Report';

import { useSelector } from 'react-redux';


function App() {


  return (
    <Fragment>
      <Navbar />
    </Fragment>
  );
}
export default App;