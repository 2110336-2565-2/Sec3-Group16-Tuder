import styled from "styled-components";
import React from "react";
import { useNavigate } from "react-router-dom";

export default function ClassDetails(props) {
  const data = props.classDetail;
  const navigate = useNavigate();

  const handleCancel = () => {
    navigate("/cancel-request/" + data.id);
  };

  return (
    <>
      <DetailGrid>
        <GridContent>
          <DetailImgFix>
            <DetailImg src={data.course_picture_url} />
          </DetailImgFix>
        </GridContent>
        <GridContent>
          <DetailContent>
            <Detailrow fsize="30px" fweight="700">
              {data.title}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              {data.ReviewCount} review(s)
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Tutor : {`${data.TutorFirstname} ${data.TutorLastname}`}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Level : {data.level}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Description :
            </Detailrow>
            <Detailrow mgtop="0.2rem" fsize="16px" fweight="400">
              {data.description}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Estimate time left : {data.estimate_time} / {data.estimate_time}{" "}
              hrs
            </Detailrow>
          </DetailContent>
        </GridContent>
      </DetailGrid>
      <AppointmentRow>
        <AppointmentContent>
          <AppointmentTitle>Appointment</AppointmentTitle>

          {/* // add for loop here */}
          <Grid>
            <GridLeft>
              <AppointmentDetail>
                <AppointmentDetailRow>
                  <AppointmentDetailTitle>Time Slot:</AppointmentDetailTitle>
                  <AppointmentDetailContent>Hee</AppointmentDetailContent>
                </AppointmentDetailRow>
                <AppointmentDetailRow>
                  <AppointmentDetailTitle>Status:</AppointmentDetailTitle>
                  <AppointmentDetailContent>Ongoing</AppointmentDetailContent>
                </AppointmentDetailRow>
              </AppointmentDetail>
            </GridLeft>
            <GridRight>
              <Button>Cancel</Button>
            </GridRight>
          </Grid>

          <Grid>
            <GridLeft>
              <AppointmentDetail>
                <AppointmentDetailRow>
                  <AppointmentDetailTitle>Time Slot:</AppointmentDetailTitle>
                  <AppointmentDetailContent>Hee</AppointmentDetailContent>
                </AppointmentDetailRow>
                <AppointmentDetailRow>
                  <AppointmentDetailTitle>Status</AppointmentDetailTitle>
                  <AppointmentDetailContent>Ongoing</AppointmentDetailContent>
                </AppointmentDetailRow>
              </AppointmentDetail>
            </GridLeft>
            <GridRight>
              <Button>Cancel</Button>
            </GridRight>
          </Grid>
        </AppointmentContent>
      </AppointmentRow>
      <CancelButtonSection>
        <CancelButton onClick={handleCancel}>Cancel Class</CancelButton>
      </CancelButtonSection>
    </>
  );
}

const DetailRight = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-top: 1rem;
`;

const Content = styled.div`
  display: flex;
  font-size: ${(props) => {
    return props.fsize;
  }};
  font-weight: ${(props) => {
    return props.fweight;
  }};
`;

const DetailGrid = styled.div`
  display: grid;
  grid-template-columns: 45% 55%;
  width: 100%;
  min-width: 750px;
`;
const GridContent = styled.div`
  display: grid;
`;
const ButtonCourse = styled.button`
  display: flex;
  width: 80px;
  padding: 5px;
  justify-content: center;
  text-decoration: none;
  border: 2px solid #ff7008;
  border-radius: 5px;
  color: #ffffff;
  background-color: #ff7008;

  &:hover {
    background-color: #ff7908;
    color: #ffffff;
  }

  cursor: pointer;
`;

const DetailImgFix = styled.div`
  margin-right: auto;
  margin-left: auto;
  margin-top: 1rem;
  width: 390px;
  height: 260px;
`;

const DetailImg = styled.img`
  display: flex;
  width: 100%;
  height: 100%;
  min-width: 270px;
  min-height: 180px;
  object-fit: cover;
  border-radius: 10px;
  object-fit: cover;
`;

const DetailContent = styled.div`
  display: flex;
  flex-direction: column;
  padding: 10px;
  min-width: 450px;
`;
const Detailrow = styled.div`
  display: flex;
  margin-top: ${(props) => {
    return props.mgtop;
  }};
  font-size: ${(props) => {
    return props.fsize;
  }};
  font-weight: ${(props) => {
    return props.fweight;
  }};
`;

const AppointmentRow = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin-top: 1rem;
  width: 100%;
  min-width: 750px;
`;

const AppointmentContent = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  min-width: 750px;

  gap: 1rem;
`;

const AppointmentTitle = styled.div`
  display: flex;
  font-size: 30px;
  font-weight: 700;
  margin-bottom: 1rem;
`;
const Grid = styled.div`
  display: grid;
  grid-template-columns: 80% 20%;
  width: 100%;
  min-width: 750px;
  border: 1px solid #e0e0e0;
  border-radius: 10px;
`;

const GridLeft = styled.div`
  grid-column: 1/2;
`;

const GridRight = styled.div`
  grid-column: 2/3;
  justify-self: center;
  align-self: center;
`;

const Button = styled.button`
  display: flex;
  justify-content: center;
  text-decoration: none;
  font-size: 16px;
  padding: 10px 20px;
  border: 1.5px solid #ff7008;
  border-radius: 5px;
  color: #ff7908;
  background-color: #ffffff;

  &:hover {
    background-color: #ff7908;
    color: #ffffff;
  }

  cursor: pointer;
`;

const AppointmentDetail = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  padding: 1rem;
`;

const AppointmentDetailRow = styled.div`
  display: flex;
  flex-direction: row;
  margin-bottom: 1rem;
`;

const AppointmentDetailTitle = styled.div`
  display: flex;
  font-size: 16px;
  font-weight: 400;
`;

const AppointmentDetailContent = styled.div`
  display: flex;
  font-size: 16px;
  font-weight: 400;
  margin-left: 1rem;
`;

const CancelButtonSection = styled.div`
  display: flex;
  justify-content: center;
  margin-top: 1rem;
  width: 100%;
  min-width: 500px;
`;

const CancelButton = styled.button`
  display: flex;
  justify-content: center;
  text-decoration: none;
  font-size: 16px;
  padding: 10px 20px;
  border: 1.5px solid #ec0a25;
  border-radius: 5px;
  color: #ffffff;
  background-color: #ec0a25;

  &:hover {
    background-color: rgb(255, 68, 68);
    color: #ffffff;
  }

  cursor: pointer;
`;
