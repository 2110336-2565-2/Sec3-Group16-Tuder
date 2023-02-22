// import css
import './App.css';

// import Router for link website
import React, { Fragment } from 'react';

// import components
import Navbar from './components/Navbar.js';
import Footer from './components/SignUpFooter';
import Home from './pages/Home.js'
// import Card from './pages/components/Card';
import SignIn from './pages/SignIn.js';
import SignUp from './pages/SignUp.js';
import Courses from './pages/Courses';
import Report from './pages/Report';
import Profile from './pages/Profile.js';

import styled from 'styled-components';



function App() {
  
  return (
    <Fragment>
      <Navbar />
    </Fragment>
    
  );
}


export default App;