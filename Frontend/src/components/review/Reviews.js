import React from "react";
import styled from "styled-components";
import { Rating } from "react-simple-star-rating";
import Review from "./Review";

export default function Reviews({ reviews }) {
  // Use only first 2 decimal places
  const totalScore = reviews.total_score.toFixed(2);
  return (
    <Container>
      <TotalRating>
        <TotalScore>{totalScore}</TotalScore>
        <Rating
          initialValue={totalScore}
          readonly={true}
          allowFraction={true}
          size="20"
          fillColor="#FF5D02"
          style={ {"gridColumn": "2 / 3"} }
        />
        <TotalReview>
          {reviews.total_review} review{reviews.total_review > 1 ? "s" : ""}
        </TotalReview>
      </TotalRating>
      <HR />
      <ReviewList>
        {reviews.reviews.map((review, index) => (
          <Review review={review} key={index} />
        ))}
      </ReviewList>
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 80px;
  background-color: white;
  border-radius: 40px 40px 0px 0px;
  width: 100%;
`;

const TotalRating = styled.div`
  display: grid;
  gap: 5px 5px;
  align-self: flex-start;
`;

const HR = styled.hr`
    color: #C4C4C4;
    width: 95%;
    margin: 25px 0px 30px 0px;
`;

const ReviewList = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  width: 100%;
`;

const TotalScore = styled.span`
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 60px;
  font-weight: medium;
  grid-row: 1 / 3;
  line-height: 0;
`;

const TotalReview = styled.p`
  grid-column: 2 / 3;
`;
