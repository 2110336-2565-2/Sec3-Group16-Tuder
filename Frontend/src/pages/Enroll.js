import React, { useEffect, useState } from "react";
import styled from "styled-components";
import { useParams } from "react-router-dom";
import FormEnrollCourse from "../components/form/FormEnrollCourse";
import { fetchCourseByIdHandler } from "../handlers/searchCourseHandler";
import { fetchTutorSchedule } from "../handlers/tutorScheduleHandler";
import WaveFooter from "../components/global/WaveFooter";
import { IsStudent } from "../components/IsAuth";
import { convertBackendSchedulesToFrontend } from "../utils/profile/scheduleConverter";

export default function Enroll() {
  const { courseID } = useParams();
  const [course, setCourse] = useState({});
  const [courseSchedule, setCourseSchedule] = useState([]);
  useEffect(() => {
    fetchCourseByIdHandler(courseID).then((res) => {
      setCourse(res.data.data);
    });
    fetchTutorSchedule(courseID).then((res) => {
      setCourseSchedule(convertBackendSchedulesToFrontend(res.data.data));
    });
  }, [courseID]);
  console.log(course, "course");
  return (
    <Container>
      <IsStudent />
      <TitleWrapper>
        <Title>Enroll - {course.title}</Title>
      </TitleWrapper>
      <FormEnrollCourse course={course} courseSchedule={courseSchedule} />
      <WaveFooter backgroundColor="white" />
    </Container>
  );
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
