import React, { useState, createContext, useContext } from "react";
import { IsUser } from "../components/IsAuth";
import styled from "styled-components";
import ClassComponentList from "../components/ClassComponentList";

const DataContext = createContext({
  data: [],
  setData: () => {},
});

export const useDataContext = () => useContext(DataContext);
export default function ClassList() {
  const [data, setData] = useState([]);
  return (
    <>
      <IsUser>
        <DataContext.Provider value={{ data, setData }}>
          <ContainerWithHeight margintop="25px">
            <ClassComponentList />
          </ContainerWithHeight>
        </DataContext.Provider>
      </IsUser>
    </>
  );
}

const ContainerWithHeight = styled.div`
  display: flex;
  flex-direction: column;
  padding: 0px 40px;
  margin-top: ${(props) => {
    return props.margintop;
  }};
  justify-content: center;
`;
