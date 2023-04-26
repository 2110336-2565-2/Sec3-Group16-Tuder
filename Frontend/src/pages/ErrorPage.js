import React from "react";
import Error from "../components/global/Error";
// import { useRouteError } from "react-router-dom";

export default function ErrorPage() {
  // const error = useRouteError();

  return (
    <>
      <Error error_msg="Sorry, this page does not exist." />
    </>
  );
}
