import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { useParams } from "react-router-dom";
import useUsername from "../hooks/useUsername.js";
import Reviews from "../components/review/Reviews.js";
import Footer from "../components/global/Footer.js";
import { IsTutor } from "../components/IsAuth.js";
import { fetchTutorReviews } from "../handlers/tutorReviewHandler.js";
import { getTutorByUsername } from "../handlers/profile/getUserHandler.js";
import { tutorReviews } from "../datas/DummyReview";

export default function TutorReviews() {
  const othersUsername = useParams().username;
  // if path is /tutors/:username/reviews, then othersUsername is not undefined
  // indicating that the user is viewing other's profile
  const isOwner = othersUsername === undefined;
  const [myUsername, handleMyUsername] = useUsername();
  const username = isOwner ? myUsername : othersUsername;

  const [reviews, setReviews] = useState({});

  // console.log("username: ", username);
  // NOTE to backend: Change this to fetch reviews from backend then it should work
  useEffect(() =>  {
    const res = fetchTutorReviews(username).then((res) => {
      setReviews(res.data.data);
    });
  }, []);
  // ---------------------------------------------

  return (
    <Container>
      {isOwner ? <IsTutor /> : null}
      <TitleWrapper>
        <Title>
          {isOwner
            ? "My Reviews"
            : reviews.tutor_firstname + " " + reviews.tutor_lastname}
        </Title>
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