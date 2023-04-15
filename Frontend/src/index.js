import React from "react";
import ReactDOM from "react-dom/client";
import Modal from "react-modal";
import "./index.css";
import App from "./App";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Home from "./pages/Home.js";
import SignIn from "./pages/SignIn.js";
import SignUp from "./pages/SignUp.js";
import Courses from "./pages/Courses";
import Report from "./pages/Report";
import ForgetPassword from "./pages/ForgetPassword";
import Profile from "./pages/Profile";
import EditProfile from "./pages/EditProfile";
import Review from "./pages/Review";
import ErrorPage from "./pages/ErrorPage";
import ResetPassword from "./pages/ResetPassword";
import ChangePassword from "./pages/ChangePassword";
import EnterNewPassword from "./pages/EnterNewPassword";
import CancelRequestList from "./pages/CancelRequestList";
import CourseReviews from "./pages/CourseReviews";
import UserCancelRequest from "./pages/UserCancelRequest";
import CancelRequestDetail from "./pages/CancelRequestDetail";
import AdminIssueReportList from "./pages/AdminIssueReportList";
import CourseDetail from "./pages/CourseDetail";
import TutorReviews from "./pages/TutorReviews";
import Enroll from "./pages/Enroll";
import ClassDetail from "./pages/ClassDetail";
import ClassList from "./pages/ClassList";
import { QueryClientProvider, QueryClient } from "react-query";

Modal.setAppElement(document.getElementById("root"));

const queryClient = new QueryClient();

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Courses /> },
      {
        path: "/report",
        element: <Report />,
      },
      {
        path: "/sign-in",
        element: <SignIn />,
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
        children: [
          {
            index: true,
            element: <Profile />,
          },
          {
            path: "reviews",
            element: <TutorReviews />,
          },
          {
            path: "edit-profile",
            element: <EditProfile />,
          },
        ],
      },
      {
        path: "/students",
        children: [
          {
            path: ":username",
            element: <Profile />,
          },
        ],
      },
      {
        path: "/tutors",
        children: [
          {
            path: ":username",
            children: [
              {
                index: true,
                element: <Profile />,
              },
              {
                path: "reviews",
                element: <TutorReviews />,
              },
            ],
          },
        ],
      },
      {
        path: "/reviews/:courseID",
        element: <Review />,
      },
      {
        path: "/courses/:courseID/enroll",
        element: <Enroll />,
      },
      {
        path: "/courses/:courseID/reviews",
        element: <CourseReviews />,
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
        path: "/cancel-request/:id",
        element: <UserCancelRequest />,
      },
      {
        path: "/cancel-requests",
        element: <CancelRequestList />,
      },
      {
        path: "/cancel-requests/:id",
        element: <CancelRequestDetail />,
      },
      {
        path: "/issuereports",
        element: <AdminIssueReportList />,
      },
      {
        path: "/courses/:id",
        element: <CourseDetail />,
      },
      {
        path: "/classes",
        children: [
          {
            index: true,
            element: <ClassList />,
          },
          {
            path: ":id",
            element: <ClassDetail />,
          },
        ],
      },
    ],
  },
]);
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
  <>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </>
);
