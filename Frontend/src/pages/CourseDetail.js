import React, { useState } from "react";
import styled from "styled-components";
import Footer from "../components/global/Footer.js";
import CourseDetails from "../components/CourseDetails.js";
import { useParams } from "react-router-dom";
import { useQuery } from "react-query";
import { fetchCourseByIdHandler } from "../handlers/searchCourseHandler.js";
import { toast } from "react-hot-toast";
export default function CourseDetail() {
  const [data, setData] = useState({});
  const { id } = useParams();
  const { isLoading, error } = useQuery(
    "course",
    () => {
      fetchCourseByIdHandler(id)
        .then((res) => {
          if (res.data.success) {
            if (res.data.data !== null) setData(res.data.data);
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

  if (isLoading || Object.keys(data).length === 0) {
    return <div>Loading...</div>;
  }

  if (data === {}) {
    return <div>Incorrect ClassID</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <>
      <Container>
        <CourseDetails courseDetail={data} />
      </Container>
      <Footer />
    </>
  );
}

const Container = styled.div`
  margin-top: 3rem;
  margin-bottom: 3rem;
  padding-left: 5%;
  padding-right: 5%;
`;
