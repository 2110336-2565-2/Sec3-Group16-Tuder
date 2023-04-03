import React, { useState, useEffect } from "react";
import styled from "styled-components";
import Reviews from "../components/review/Reviews.js";
import Footer from "../components/global/Footer.js";
import { IsStudent } from "../components/IsAuth.js";
import { courseReviews } from "../datas/DummyReview";

export default function CourseReviews() {
  const [reviews, setReviews] = useState([]);

  // NOTE to backend: Change this to fetch reviews from backend then it should work
  useEffect(() => {
    setReviews(courseReviews);
  }, []);
  // ---------------------------------------------

  return (
    <Container>
      {/* <IsStudent /> */}
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

const NoReviewContainer = styled.div`
  display: flex;
  width: 100%;
  background-color: white;
  justify-content: center;
  align-items: center;
  min-height: 50vh;
`;
