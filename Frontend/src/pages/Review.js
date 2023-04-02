import React, { useState } from "react";
import { Rating } from "react-simple-star-rating";
import styled from "styled-components";
import { getUserID } from "../utils/jwtGet";
import { useParams } from 'react-router-dom';
import { useNavigate } from "react-router-dom";
import Button from "../components/global/Button";
import TextInput from "../components/profile/TextInput";
import Footer from "../components/global/Footer";
import { IsStudent } from '../components/IsAuth.js'
import { confirm } from "../components/global/customToast";

export default function Review() {
  const studentID = getUserID();
  const { courseID } = useParams();
  console.log(studentID)
  console.log(courseID)
  const [rating, setRating] = useState(0);
  const [reviewMsg, setReviewMsg] = useState("");
  const navigate = useNavigate();
  // Request backend to implement this API
  const course = {
    course_name: "Search Engine Optimization",
    tutor_firstname: "John",
    tutor_lastname: "Doe",
  };
  
  function handleRatingClick(newRating) {
    setRating(newRating);
  }
  function handleSubmit(e) {
    e.preventDefault();

  }

  return (
    <Container>
      {/* <IsStudent /> */}
      <Form onSubmit={handleSubmit}>
        <Topic>Course Review</Topic>
        <Detail>
          {course.course_name}
          <br />
          {course.tutor_firstname + " " + course.tutor_lastname}
        </Detail>
        <RatingSection>
          <Rating onClick={handleRatingClick} initialValue={rating} />
          <StarCount>{rating} stars</StarCount>
        </RatingSection>
        <TextInput
          type="textArea"
          label="Description"
          id="review_msg"
          name="review_msg"
          value={reviewMsg}
          onChange={(e) => setReviewMsg(e.target.value)}
          width="100%"
        />
        <ButtonSection>
          <Button variance="cancel" onClick={() => confirm("Are you sure you want to discard this form?", ()=>navigate("/"))}>
            Cancel
          </Button>
          <Button variance="submit" type="submit" >Send</Button>
        </ButtonSection>
      </Form>
      <Footer />
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: #fdedeb;
`;

const Form = styled.form`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  width: 80%;
  background-color: white;
  padding: 30px 75px;
  margin: 150px 0px;
  border-radius: 10px;
`;

const Topic = styled.h1`
    text-align: center;
`;

const Detail = styled.h2`
  text-align: center;
  font-size: 16px;
  font-weight: lighter;
  margin-top: 15px;
`;

const RatingSection = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 15px;
`;

const StarCount = styled.span`
  font-size: 16px;
  font-weight: 500px;
`;

const ButtonSection = styled.div`
  display: flex;
  justify-content: flex-end;
  align-items: center;
  width: 100%;
  gap: 20px;
  margin-top: 30px;
`;
