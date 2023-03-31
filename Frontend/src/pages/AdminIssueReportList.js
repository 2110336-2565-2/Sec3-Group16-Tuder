import React, { useState, createContext, useContext } from "react";
import styled from "styled-components";
import WaveFooter from "../components/global/WaveFooter.js";
import { IsAdmin } from "../components/IsAuth.js";
import AdminIssueReportList from "../components/AdminIssueReportList.js";

const DataContext = createContext({
  data: [],
  setData: () => {},
});

export const useDataContext = () => useContext(DataContext);
export default function AdminIssueReportListPage() {
  const [data, setData] = useState([]);

  return (
    <Container>
      <IsAdmin>
        <DataContext.Provider value={{ data, setData }}>
          <ContainerWithHeight margintop="100px">
            <AdminIssueReportList />
          </ContainerWithHeight>
          <WaveFooterWrapper>
            <WaveFooter />
          </WaveFooterWrapper>
        </DataContext.Provider>
      </IsAdmin>
    </Container>
  );
}

const ContainerWithHeight = styled.div`
  display: flex;
  flex-direction: column;
  padding: 0px 30px;
  margin-top: ${(props) => {
    return props.margintop;
  }};
  justify-content: center;
`;

const WaveFooterWrapper = styled.div`
  position: absolute;
  bottom: 0;
  width: 100%;
`;

const Container = styled.div`
  position: relative;
  min-height: 100vh;
  width: 100%;
`;
