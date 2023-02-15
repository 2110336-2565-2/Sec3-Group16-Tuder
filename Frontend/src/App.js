import { useState } from 'react';
// import css
import './App.css';

// import Router for link website
import { BrowserRouter, Route, Routes } from 'react-router-dom';

// import components
import Navbar from './components/Navbar.js';
import Home from './pages/Home.js'
// import Card from './pages/components/Card';
import SignIn from './pages/SignIn.js';
import SignUp from './pages/SignUp.js';
import Courses from './pages/Courses';
import Report from './pages/Report';
import ForgetPassword from './pages/ForgetPassword';



function App() {
  const [role, setRole] = useState('guest');
  const [signUpRole, setSignUpRole] = useState('student');
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navbar role={role} />}>
          <Route index element={<Home />}></Route>
          <Route path="Report" element={<Report />}></Route>
          <Route path="Courses" element={<Courses />}></Route>
          <Route path="SignIn" element={<SignIn />}></Route>
          <Route path="SignUp" element={<SignUp role={signUpRole}/>}></Route>
          <Route path="Forgetpassword" element={<ForgetPassword />}></Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
export default App;