import React, { useState, useEffect } from "react";
import styled from "styled-components";
import Reviews from "../components/review/Reviews.js";
import Footer from "../components/global/Footer.js";
import Error from "../components/global/Error.js";
import { IsStudent } from "../components/IsAuth.js";
import { courseReviews } from "../datas/DummyReview";
import { getCourseReviewHandler } from "../handlers/courseReviewHandler.js";
import { useParams } from "react-router-dom";

export default function CourseReviews() {
  const [error, setError] = useState(false);
  const [reviews, setReviews] = useState([]);
  const { courseID } = useParams();

  useEffect(() => {
    getCourseReviewHandler(courseID)
      .then((res) => {
        setReviews(res.data.data);
      })
      .catch((err) => {
        setError(true);
      });
  }, []);
  // ---------------------------------------------
  if (error) {
    return (
      <>
        <Error error_msg="Sorry, this course does not exist." />
        <Footer />
      </>
    );
  } else {
    return (
      <Container>
        <IsStudent>
          <TitleWrapper>
            <Title>{reviews.course_title}</Title>
          </TitleWrapper>
          {reviews.total_review > 0 ? (
            <Reviews reviews={reviews} />
          ) : (
            <NoReviewContainer>
              <p>No review yet.</p>
            </NoReviewContainer>
          )}
          <Footer />
        </IsStudent>
      </Container>
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

const NoReviewContainer = styled.div`
  display: flex;
  width: 100%;
  background-color: white;
  justify-content: center;
  align-items: center;
  min-height: 50vh;
`;
