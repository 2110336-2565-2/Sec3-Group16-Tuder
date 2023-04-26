import styled from "styled-components";
import React from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Timezone, DateFormat } from "../datas/DateFormat";
import { changeStatusHandler } from "../handlers/classesHandler";
import { toast } from "react-hot-toast";
import { getRole } from "../utils/jwtGet";

export default function ClassDetails(props) {
  const data = props.classDetail;
  const navigate = useNavigate();
  const role = getRole();

  const { id } = useParams();
  const handleCancel = () => {
    navigate("/cancel-request/" + id, {
      state: { course_name: data.course_name, tutor_name: data.tutor_name },
    });
  };

  var timeLeft =
    data.appointments !== undefined
      ? data.appointments.filter((appointment) => {
          return appointment.status === "comingsoon";
        })
      : 0;

  const attend = (e, id) => {
    e.preventDefault();
    changeStatusHandler(id, "pending").then((res) => {
      if (res.data.success) {
        toast.success("Attendded");
        window.location.reload();
      } else {
        toast.error("Failed Attended");
      }
    });
  };

  const postpone = (e, id) => {
    e.preventDefault();
    changeStatusHandler(id, "postponed").then((res) => {
      if (res.data.success) {
        toast.success("Postponed");
        window.location.reload();
      } else {
        toast.error("Failed Postponed");
      }
    });
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
              {data.course_name}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Tutor : {data.tutor_name}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Level : {data.level}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Description :
            </Detailrow>
            <Detailrow mgtop="0.2rem" fsize="16px" fweight="400">
              {data.course_description}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Estimate time left : {timeLeft.length} hrs
            </Detailrow>
          </DetailContent>
        </GridContent>
      </DetailGrid>
      <AppointmentRow>
        <AppointmentContent>
          <AppointmentTitle>Appointment</AppointmentTitle>

          {/* // add for loop here */}
          {data.appointments !== undefined ? (
            data.appointments.map((app) => {
              if (app.status === "verifying") {
                return (
                  <Grid key={app.id}>
                    <GridLeft>
                      <AppointmentDetail>
                        <AppointmentDetailRow>
                          <AppointmentDetailTitle>
                            Time Slot:
                          </AppointmentDetailTitle>
                          <AppointmentDetailContent>
                            {new Date(app.begin_at).toLocaleString(
                              Timezone,
                              DateFormat
                            )}{" "}
                            {" - "}{" "}
                            {new Date(app.end_at).toLocaleString(
                              Timezone,
                              DateFormat
                            )}
                          </AppointmentDetailContent>
                        </AppointmentDetailRow>
                        <AppointmentDetailRow>
                          <AppointmentDetailTitle>
                            Status:
                          </AppointmentDetailTitle>
                          <AppointmentDetailContent>
                            {app.status}
                          </AppointmentDetailContent>
                        </AppointmentDetailRow>
                      </AppointmentDetail>
                    </GridLeft>
                    <GridRight>
                      {role === "student" ? (
                        <>
                          <Button onClick={(e) => attend(e, app.id)}>
                            {" "}
                            Attend Class
                          </Button>
                          <SpaceBox></SpaceBox>
                          <Button onClick={(e) => postpone(e, app.id)}>
                            {" "}
                            Postpone Class
                          </Button>
                        </>
                      ) : (
                        <div></div>
                      )}
                    </GridRight>
                  </Grid>
                );
              } else {
                return (
                  <Grid key={app.id}>
                    <GridLeft>
                      <AppointmentDetail>
                        <AppointmentDetailRow>
                          <AppointmentDetailTitle>
                            Time Slot:
                          </AppointmentDetailTitle>
                          <AppointmentDetailContent>
                            {new Date(app.begin_at).toLocaleString(
                              Timezone,
                              DateFormat
                            )}{" "}
                            {" - "}{" "}
                            {new Date(app.end_at).toLocaleString(
                              Timezone,
                              DateFormat
                            )}
                          </AppointmentDetailContent>
                        </AppointmentDetailRow>
                        <AppointmentDetailRow>
                          <AppointmentDetailTitle>
                            Status:
                          </AppointmentDetailTitle>
                          <AppointmentDetailContent>
                            {app.status}
                          </AppointmentDetailContent>
                        </AppointmentDetailRow>
                      </AppointmentDetail>
                    </GridLeft>
                    <GridRight></GridRight>
                  </Grid>
                );
              }
            })
          ) : (
            <div></div>
          )}
        </AppointmentContent>
      </AppointmentRow>
      <CancelButtonSection>
        <CancelButton onClick={handleCancel}>Cancel Class</CancelButton>
      </CancelButtonSection>
    </>
  );
}

const SpaceBox = styled.div`
  padding: 4px;
`;

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
  grid-template-columns: 75% 25%;
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
  width: 165px;
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
