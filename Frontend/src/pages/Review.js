import React, { useState,useEffect } from "react";
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
import { toast } from 'react-hot-toast';
import submitReviewHandler, { fetchTutorHandler } from "../handlers/submitReviewHandler";
import { fetchCourseHandler} from "../handlers/submitReviewHandler";

export default function Review() {
  const studentID = getUserID();
  const { courseID } = useParams();
  console.log(studentID)
  console.log(courseID)
  const [rating, setRating] = useState(0);
  const [reviewMsg, setReviewMsg] = useState("");
  const navigate = useNavigate();
  const [course, setCourse] = useState({});
  const [tutor, setTutor] = useState({});
  // Note to backend: Change this to fetch course details from backend
  useEffect(() => {
    fetchCourseHandler(courseID).then((res) => {
      setCourse(res.data.data)
      fetchTutorHandler(res.data.data.tutor_id).then((res) => {
        setTutor(res.data.data)
      });
    });
  }, []);

  // ---------------------------------------------
  
  function handleRatingClick(newRating) {
    setRating(newRating);
  }
  async function handleSubmit(e) {
    e.preventDefault();
    const reviewPayload = {
      "course_id" : courseID,
      "rating" : rating,
      "review_message" : reviewMsg,
    }
    try{
      await submitReviewHandler(reviewPayload)
      toast.success('Submit Review Success')
    }
    catch(err){
      toast.error('Submit Review Failed')
    }
    
    // Note to backend: Change this to send review to backend
    // each field is already stored in variable corresponding to its name

  }

  return (
    <Container>
      {/* <IsStudent /> */}
      <Form onSubmit={handleSubmit}>
        <Topic>Course Review</Topic>
        <Detail>
          {course.title}
          <br />
          {tutor.firstname + " " + tutor.lastname}
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
