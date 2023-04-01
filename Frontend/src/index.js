import React from 'react';
import ReactDOM from 'react-dom/client';
import Modal from 'react-modal';
import './index.css';
import App from './App';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Home from './pages/Home.js'
import SignIn from './pages/SignIn.js';
import SignUp from './pages/SignUp.js';
import Courses from './pages/Courses';
import Report from './pages/Report';
import ForgetPassword from './pages/ForgetPassword';
import Profile from './pages/Profile';
import EditProfile from './pages/EditProfile';
import Review from './pages/Review';
import ErrorPage from './pages/ErrorPage';
import ResetPassword from './pages/ResetPassword';
import ChangePassword from './pages/ChangePassword';
import EnterNewPassword from './pages/EnterNewPassword';
import CancelRequestList from './pages/CancelRequestList';
import UserCancelRequest from './pages/UserCancelRequest';
import CancelRequestDetail from './pages/CancelRequestDetail';
import { QueryClientProvider, QueryClient } from 'react-query';

Modal.setAppElement(document.getElementById('root'))

const queryClient = new QueryClient();

const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
    errorElement: <ErrorPage />,
    children: [
      {index: true, element: <Home /> },
      {
        path: "/report",
        element: <Report />,
      },
      {
        path: "/courses",
        element: <Courses />,
      },
      {
        path: "/sign-in",
        element: <SignIn/>,
      },
      {
        path: "/sign-up",
        element: <SignUp />,
      },
      {
        path: "/forget-password",
        element: <ForgetPassword />,
      },
      {
        path: "/profile",
        element: <Profile />,
      },
      {
        path: "/edit-profile",
        element: <EditProfile />,
      },
      {
        path: "/review/:id",
        element: <Review />,
      },
      {
        path: "/reset-password/:token",
        element: <ResetPassword />,
      },
      {
        path: "/change-password",
        element: <ChangePassword />,
      },
      {
        path: "/enter-new-password",
        element: <EnterNewPassword />,
      },
      {
        path: "/user-cancel-request/:id",
        element: <UserCancelRequest />,
      },
      {
        path: "/cancel-request-list",
        element: <CancelRequestList />,
      },
      {
        path: "/cancel-request-detail/:id",
        element: <CancelRequestDetail />,
      }
    ]
  }
])
const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <>
    <QueryClientProvider client={queryClient}>
    <RouterProvider router={router} />
    </QueryClientProvider>
  </>
);

