import styled from "styled-components";
import { useQuery } from "react-query";
import { fetchAdminIssueReportHandler } from "../handlers/AdminIssueReportHandler";
import { useDataContext } from "../pages/AdminIssueReportList";
import AdminIssueReport from "../components/AdminIssueReport";
import React from "react";

export default function AdminIssueReportList() {
  
  const { data, setData } = useDataContext();

  const { isLoading, error } = useQuery(
    "AdminIssueReport",
    () => {
      fetchAdminIssueReportHandler()
        .then((res) => {
          console.log(res.data.result);
          if (res.data.success) {
            if (res.data.result !== null) setData(res.data.result);
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    {
      refetchOnWindowFocus: false,
    }
  );

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }
  console.log(data);
  if (data === null) {
    return <div>Error</div>;
  }

  if (data === []) {
    return <div>Empty</div>;
  }

  return (
    <Container>
      <h1>AdminIssueReport Requests</h1>
      <RequestListPage>
        <RequestListcontent>
          {data.map((item) => (
            <div key={item.issuereport_id}>
              <AdminIssueReport
                key={item.issuereport_id}
                issuereport_id={item.issuereport_id}
                title={item.title}
                description={item.description}
                contact={item.contact}
                report_date={item.report_date}
                status={item.status}
              />
            </div>
          ))}
        </RequestListcontent>
      </RequestListPage>
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 10px;
  margin-top: 5px;
  margin-bottom: 20px;
`;

const RequestListPage = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 5%;
`;

const RequestListcontent = styled.div`
  display: flex;
  flex-direction: column;
  gap: 20px;
`;
