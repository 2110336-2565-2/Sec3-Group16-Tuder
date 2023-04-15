import styled from "styled-components";
import { useQuery } from "react-query";
import { fetchAdminIssueReportHandler } from "../handlers/AdminIssueReportHandler";
import { useDataContext } from "../pages/AdminIssueReportList";
import AdminIssueReport from "../components/AdminIssueReport";
import React, { useState } from "react";

export default function AdminIssueReportList() {
  const [ item, setItem ] = useState([]);

  function deleteIssue(id) {
    try {
      setItem(prevNotes => prevNotes.filter(issueItem => issueItem.issuereport_id  !== id));
    } catch (error) {
      console.error("Error deleting issue report:", error);
    }
    
  }
  const { isLoading, error } = useQuery(
    "AdminIssueReport",
    () => {
      fetchAdminIssueReportHandler()
        .then((res) => {
          console.log(res.data.result);
          if (res.data.success) {
            if (res.data.result !== null) setItem(res.data.result);
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
  console.log(item);
  if (item === null) {
    return <div>Error</div>;
  }

  if (item === []) {
    return <div>Empty</div>;
  }


  return (
    <Container>
      <h1>AdminIssueReport Requests</h1>
      <RequestListPage>
        <RequestListcontent>
          {
              item.map(issueItem => (
              <AdminIssueReport
                key={issueItem.issuereport_id}
                issuereport_id={issueItem.issuereport_id}
                title={issueItem.title}
                description={issueItem.description}
                contact={issueItem.contact}
                report_date={issueItem.report_date}
                status={issueItem.status}
                onDelete={deleteIssue}
              />
          ))
          }
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