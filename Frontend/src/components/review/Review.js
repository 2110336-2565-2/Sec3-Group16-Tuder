import React, { useState } from "react";
import styled from "styled-components";
import { Rating } from "react-simple-star-rating";
import formatDate from "../../utils/review/dateConverter";

export default function Review({ review }) {
  const [score, setScore] = useState(parseInt(review.score));

  return (
    <Container>
      {review.course_title? <h3>{review.course_title}</h3> : null}
      <RowWrapper>
        <Rating
          initialValue={score}
          readonly={true}
          size="20"
          fillColor="#FF5D02"
        />
        <ReviewTime>
          {review.review_time_at
            ? formatDate(review.review_time_at.split("T")[0])
            : review.review_time_at}
        </ReviewTime>
      </RowWrapper>
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

const RowWrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
`;

const ReviewTime = styled.p``;
