
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

import { useSelector } from 'react-redux';


function App() {

  const {role} = useSelector(state => state.role);
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navbar role={role} />}>
          <Route index element={<Home />}></Route>
          <Route path="Report" element={<Report />}></Route>
          <Route path="Courses" element={<Courses />}></Route>
          <Route path="SignIn" element={<SignIn />}></Route>
          <Route path="SignUp" element={<SignUp />}></Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
export default App;
