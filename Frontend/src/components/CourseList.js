import styled from "styled-components";
import Course from "../components/Course";
import { useQuery } from "react-query";
import { fetchCourseHandler } from "../handlers/searchCourseHandler";
import { useDataContext } from "../pages/Courses";
import React from "react";

export default function CourseList() {
  const { data, setData } = useDataContext();

  const { isLoading, error } = useQuery(
    "courses",
    () => {
      fetchCourseHandler()
        .then((res) => {
          if (res.data.success) {
            setData(res.data.result);
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
    <>
      <CourseListPage>
        {data.map((item) => (
          <CourseListcontent key={item.course_id}>
            <Course
              course_id={item.course_id}
              title={item.title}
              topic={item.topic}
              tutor={item.tutor_name}
              subject={item.subject}
              time={item.estimate_time}
              price={item.price_per_hour}
              img={item.course_picture_url}
            />
          </CourseListcontent>
        ))}
      </CourseListPage>
    </>
  );
}

const CourseListPage = styled.div`
  display: flex;
  flex-direction: row;
  gap: 20px;
  justify-content: center;
  align-items: center;
`;

const CourseListcontent = styled.div`
  width: 300px;
  display: flex;
  box-shadow: 0px 1px 4px rgba(0, 0, 0, 0.25);
  border-radius: 10px;
  padding: 10px 10px 10px 10px;
  justify-content: center;
`;
