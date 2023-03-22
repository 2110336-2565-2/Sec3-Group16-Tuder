// import css
import './App.css';

// import Router for link website
import React, { Fragment } from 'react';
import { Toaster } from 'react-hot-toast';

// import components
import Navbar from './components/Navbar.js';
import Footer from './components/footer/SignUpFooter';
import Home from './pages/Home.js'
// import Card from './pages/components/Card';
import SignIn from './pages/SignIn.js';
import SignUp from './pages/SignUp.js';
import Courses from './pages/Courses';
import Report from './pages/Report';
import Profile from './pages/Profile.js';
import EditProfile from './pages/EditProfile.js';
// import ChangePassword from './pages/ChangePassword.js';

import styled from 'styled-components';



function App() {
  
  return (
    <Fragment>
      <Toaster />
      <Navbar />
    </Fragment>
    
  );
}


export default App;