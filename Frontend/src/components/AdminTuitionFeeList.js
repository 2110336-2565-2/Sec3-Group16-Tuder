import styled from "styled-components";
import { useQuery } from "react-query";
import { fetchCancellingRequestsHandler } from "../handlers/cancellingRequestHandler";
import { useDataContext } from "../pages/AdminTuitionFeeList";
import AdminTuitionFee from "../components/AdminTuitionFee";
import React from "react";
import { useNavigate } from "react-router-dom";

export default function AdminTuitionFeeList() {
  const { data, setData } = useDataContext();
  const navigate = useNavigate();

  const { isLoading, error } = useQuery(
    "cancellingClass",
    () => {
      fetchAdminTuitionFeeHandler()
        .then((res) => {
          if (res.data.success) {
            if (res.data.data !== null) setData(res.data.data);
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
              // key={item.AdminTuitionFeeId}
              // onClick={(e) =>
              //   navigate("/cancel-requests/" + item.AdminTuitionFeeId)
              // }
            >
              <AdminTuitionFee
                key={item.AdminTuitionFeeId}
                title={item.title}
                img={item.img_url}
                student={item.student}
                tutor={item.tutor}
                date={item.date}

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
