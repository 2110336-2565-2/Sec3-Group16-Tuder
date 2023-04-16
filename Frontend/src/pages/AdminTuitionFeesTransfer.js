import React from "react";
import styled from "styled-components";
import toast from "react-hot-toast";
import QRPayment from "../components/global/QRPayment";
import Button from "../components/global/Button";
import Footer from "../components/global/Footer";
import { IsAdmin } from "../components/IsAuth";
import { useNavigate, useLocation } from "react-router-dom";
import { changeStatusHandler } from "../handlers/classesHandler";
import { Timezone, DateFormat } from "../datas/DateFormat.js";

export default function AdminTuitionFeesTransfer() {
  function fieldNameToTitle(fieldName) {
    return fieldName
      .split("_")
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(" ");
  }
  // get current URL
  const location = useLocation();
  const item = location.state.data;
  const locationPath = location.pathname;
  const navigate = useNavigate();
  const data = {
    course_title: item.title,
    course_picture: item.img,
    student_name: item.student_name,
    tutor_name: item.tutor_name,
    tutor_id: item.tutorID,
    subject: item.subject,
    topic: item.topic,
    date: new Date(item.appointmentEndAt).toLocaleString(Timezone, DateFormat),
    price_per_hour: item.price_per_hour,
    appointmentID: item.appointmentID,
  };

  const tableContent = [
    "student_name",
    "tutor_name",
    "subject",
    "topic",
    "date",
    "price_per_hour",
  ];
  function paymentCallback(appointmentID) {
    // Change this to PUT /appointment/updatestatus/:appointmentID with status = "completed"
    changeStatusHandler(appointmentID, "completed").then((res) => {
      toast.success("Update status successfully");
      navigate("/admin-tuition-fees");
    }).catch((err) => {
      toast.error("Something went wrong");
    });
    // ----------------------------------
  }
  return (
    <>
      <IsAdmin />
      <Container>
        <Wrapper>
          <Title>Tuition Fee - {data.course_title}</Title>
          <Card>
            <ImageWrapper>
              <Image src={data.course_picture} />
            </ImageWrapper>
            <Table>
              <thead>
                <Tr>
                  <Th>Information</Th>
                </Tr>
              </thead>
              <tbody>
                {tableContent.map((item, index) => (
                  <Tr key={index}>
                    <Td>
                      <FieldName>{fieldNameToTitle(item)}</FieldName>
                    </Td>
                    <Td>{data[item]}</Td>
                  </Tr>
                ))}
              </tbody>
            </Table>
          </Card>
          <QRWrapper>
            <QRPayment
              amount={data["price_per_hour"]}
              tutorID={data["tutor_id"]}
              appointmentID={data["appointmentID"]}
              callback={paymentCallback}
              callbackData={data["appointmentID"]}
            />
          </QRWrapper>
          <ButtonWrapper>
            <Button
              variance="cancel"
              onClick={() =>
                navigate(
                  locationPath.substring(0, locationPath.lastIndexOf("/"))
                )
              }
            >
              Back
            </Button>
          </ButtonWrapper>
        </Wrapper>
      </Container>
      <Footer />
    </>
  );
}

const Container = styled.div`
  display: flex;

  justify-content: center;
  align-items: center;
  background-color: #ebebeb;
  width: 100%;
  min-height: 100vh;
  padding: 50px 0;
`;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 90%;
  height: 90%;
  background-color: white;
  padding: 50px;
  border-radius: 10px;
  box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.75);
`;

const Title = styled.h1`
  font-size: 35px;
  font-weight: bold;
  color: #45424a;
`;

const Card = styled.div`
  display: flex;
  justify-content: center;
  width: 80%;
  margin-top: 50px;
`;

const ImageWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 50%;
  height: 400px;
  background-color: red;
`;

const Image = styled.img`
  width: 100%;
  height: 100%;
  object-fit: cover;
`;

const Table = styled.table`
  border-collapse: collapse;
  border-spacing: 0;
  border: 1px solid #e5e7eb;
  background-color: white;
`;

const Tr = styled.tr`
  border-bottom: 1px solid #e5e7eb;
`;

const Th = styled.th`
  text-align: left;
  font-size: 24px;
  font-weight: medium;
  padding: 10px 15px;
`;

const Td = styled.td`
  text-align: right;
  font-size: 18px;
  font-weight: light;
  padding: 10px 15px;
`;

const FieldName = styled.p`
  text-align: left;
  font-weight: 500;
  font-size: 20px;
`;

const QRWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 80%;
  margin-top: 50px;
`;

const ButtonWrapper = styled.div`
  display: flex;
  width: 70%;
  justify-content: flex-end;
  margin-top: 50px;
`;
