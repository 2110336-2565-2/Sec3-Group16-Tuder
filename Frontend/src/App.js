// import css
import './App.css';

// import Router for link website
import React, { Fragment, useEffect } from 'react';
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

import { createContext, useContext } from 'react';
import useRole from './hooks/useRole';
import { verify } from './utils/jwtGet';

const DataContext = createContext({
  role:'guest',
  handleRole: () => {}
});

export const useDataContext = () => useContext(DataContext);

function App() {
  const [role, handleRole] = useRole();


  return (
    <Fragment>
      <DataContext.Provider value={[{role}, {handleRole}]}>
      <Toaster />
      <Navbar />
      </DataContext.Provider>
    </Fragment>
    
  );
}


export default App;