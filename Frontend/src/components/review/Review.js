import React, { useState } from "react";
import styled from "styled-components";
import { Rating } from "react-simple-star-rating";

export default function Review({ review }) {
  const [score, setScore] = useState(parseInt(review.score));
  return (
    <Container>
      <Rating
        initialValue={score}
        readonly={true}
        size="20"
        fillColor="#FF5D02"
      />
      <ReviewMsg>{review.review_msg}</ReviewMsg>
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 20px 20px;
  width: 100%;
  min-height: 160px;
  background-color: #ebebeb;
`;

const ReviewMsg = styled.p`
  margin-top: 20px;
`;
