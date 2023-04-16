import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { useParams } from "react-router-dom";
import Error from "../components/global/Error";
import FormEnrollCourse from "../components/form/FormEnrollCourse";
import { fetchCourseByIdHandler } from "../handlers/searchCourseHandler";
import { fetchTutorSchedule } from "../handlers/tutorScheduleHandler";
import WaveFooter from "../components/global/WaveFooter";
import { IsEnroll } from "../components/IsAuth";
import { convertBackendSchedulesToFrontend } from "../utils/profile/scheduleConverter";

export default function Enroll() {
  const [error, setError] = useState(false);
  const { courseID } = useParams();
  const [course, setCourse] = useState({});
  const [courseSchedule, setCourseSchedule] = useState([]);
  useEffect(() => {
    fetchCourseByIdHandler(courseID)
      .then((res) => {
        setCourse(res.data.data);
      })
      .catch((err) => {
        setError(true);
      });
    fetchTutorSchedule(courseID)
      .then((res) => {
        setCourseSchedule(convertBackendSchedulesToFrontend(res.data.data));
      })
      .catch((err) => {
        setError(true);
      });
  }, [courseID]);
  console.log(course, "course");
  if (error) {
    return (
      <>
        <Error error_msg="Sorry, this course does not exist." />
        <WaveFooter backgroundColor="white" />
      </>
    );
  } else {
    return (
      <IsEnroll>
        <Container>
          <TitleWrapper>
            <Title>Enroll - {course.title}</Title>
          </TitleWrapper>
          <FormEnrollCourse course={course} courseSchedule={courseSchedule} />
          <WaveFooter backgroundColor="white" />
        </Container>
      </IsEnroll>
    );
  }
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #eb7b42;
  width: 100%;
`;

const Title = styled.h1`
  color: white;
`;

const TitleWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 150px;
`;
