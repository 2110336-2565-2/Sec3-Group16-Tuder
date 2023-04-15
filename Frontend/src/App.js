// import css
import "./App.css";

// import Router for link website
import React, { Fragment, useEffect } from "react";
import { Toaster } from "react-hot-toast";

// import components
import Navbar from "./components/Navbar.js";

import { createContext, useContext } from "react";
import useRole from "./hooks/useRole";
import { useState } from "react";
const DataContext = createContext({
  role: "guest",
  handleRole: () => {},
});

export const useDataContext = () => useContext(DataContext);

function App() {
  const [role, handleRole] = useRole();

  return (
    <Fragment>
      <DataContext.Provider value={[{ role }, { handleRole }]}>
        <Toaster />
        <Navbar />
      </DataContext.Provider>
    </Fragment>
  );
}

export default App;
