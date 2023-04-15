import styled from "styled-components";
import { useQuery } from "react-query";
import { fetchAdminTuitionFeeHandler } from "../handlers/AdminTuitionFeeHandler";
import { useDataContext } from "../pages/AdminTuitionFeeList";
import AdminTuitionFee from "../components/AdminTuitionFee";
import React from "react";
import { useNavigate } from "react-router-dom";

export default function AdminTuitionFeeList() {
  const { data, setData } = useDataContext();
  const navigate = useNavigate();

  const { isLoading, error } = useQuery(
    "tuitionfee",
    () => {
      fetchAdminTuitionFeeHandler()
        .then((res) => {
          if (res.data.success) {
            if (res.data.result !== null) setData(res.data.result);
          }
          console.log(res.data)
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
      <h1>AdminTuitionFee Requests</h1>
      <RequestListPage>
        <RequestListcontent>
          {data.map((item) => (
            <div
               key={item.appointmentID}
              onClick={(e) =>
                navigate("/admin-tuition-fees/" + item.appointmentID)
              }
              >
              <AdminTuitionFee
                key={item.appointmentID}
                AppointmentID={item.appointmentID}
                AppointmentBeginAt = {item.appointmentBeginAt}
                AppointmentEndAt = {item.appointmentEndAt}
                AppointmentStatus = {item.appointmentStatus}
                MatchID = {item.matchID}
                StudentID = {item.studentID}
                TutorID = {item.tutorID}
                Student_name = {item.student_name}
                Tutor_name = {item.tutor_name}
                Title={item.title}
                Subject={item.subject}
                Topic={item.topic}
                Img={item.img}
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
