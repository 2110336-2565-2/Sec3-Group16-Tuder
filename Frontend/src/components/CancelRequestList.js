import styled from "styled-components";
import { useQuery } from "react-query";
import { fetchCancellingRequestsHandler } from "../handlers/cancellingRequestHandler";
import { useDataContext } from "../pages/CancelRequestList";
import CancelRequest from "../components/CancelRequest";
import React from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-hot-toast";
export default function CancelRequestList() {
  const { data, setData } = useDataContext();
  const navigate = useNavigate();

  const { isLoading, error } = useQuery(
    "cancellingClass",
    () => {
      fetchCancellingRequestsHandler()
        .then((res) => {
          if (res.data.success) {
            if (res.data.data !== null) {
              setData(res.data.data);
            }
          }
        })
        .catch((err) => {
          console.log(err);
          toast.error("Something went wrong");
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

  if (data === null) {
    return <div>Error</div>;
  }

  if (data === []) {
    return <div>Empty</div>;
  }

  return (
    <Container>
      <h1>Cancellation Requests</h1>
      <RequestListPage>
        <RequestListcontent>
          {data.map((item) => (
            <div
              key={item.cancelRequestId}
              onClick={(e) =>
                navigate("/cancel-requests/" + item.cancelRequestId)
              }
            >
              <CancelRequest
                key={item.cancelRequestId}
                title={item.title}
                img={item.img_url}
                reporter={item.reporter}
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
  margin-bottom: 25px;
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
