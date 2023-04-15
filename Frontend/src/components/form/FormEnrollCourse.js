import React, { useState } from "react";
import styled from "styled-components";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import TimeSelector from "../enrollment/TimeSelector";
import PaymentModal from "../global/PaymentModal";
import { getStudentID } from "../../utils/jwtGet";
import { convertFrontendSchedulesToBackend } from "../../utils/profile/scheduleConverter";
import { submitEnrollFormHandler } from "../../handlers/match/enrollHandler";

export default function FormEnrollCourse({ course, courseSchedule }) {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const navigate = useNavigate();
  const studentID = getStudentID();
  const [totalHours, setTotalHours] = useState(0);
  const [selectedSchedule, setSelectedSchedule] = useState([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  function handleSubmit(e) {
    e.preventDefault();
    if (isSubmitting) return;
    setIsSubmitting(true);
    const paymentLoadingToast = toast.loading("Loading...");
    // Filter out the timeSlot that is not selected
    const filteredSchedule = selectedSchedule.map((day) => ({
      day: day.day,
      timeSlot: day.timeSlot.filter((timeSlot) => timeSlot.isSelected),
    }));
    // Convert selectedSchedule to backend format
    const backendSchedule = convertFrontendSchedulesToBackend(filteredSchedule);
    if (backendSchedule.length === 0) {
      toast.dismiss(paymentLoadingToast);
      toast.error("Please select at least one time slot");
      setIsSubmitting(false);
      return;
    }
    // Create data to send to backend
    const data = {
      course_id: course.id,
      total_hour: parseInt(totalHours),
      schedule: backendSchedule,
    };
    console.log(data);
    toast.dismiss(paymentLoadingToast);
    setIsModalOpen(true);
    setIsSubmitting(false);

    // submitEnrollFormHandler(data)
    //   .then((res) => {
    //     toast.dismiss(loadingToast);
    //     toast.success("Enroll Success");
    //     navigate("/");
    //     setIsSubmitting(false);
    //   })
    //   .catch((err) => {
    //     toast.dismiss(loadingToast);
    //     toast.error("Enroll Failed");
    //     setIsSubmitting(false);
    //   });
  }
  console.log("courseSchedule: ", courseSchedule);
  return (
    <Container>
      <PaymentModal
        title={"Enroll - " + course.title}
        isOpen={isModalOpen}
        setIsOpen={setIsModalOpen}
        courseID={course.id}
        amount={totalHours * course.price_per_hour}
      />
      <Form onSubmit={handleSubmit}>
        {courseSchedule.length > 0 ? (
          <TimeSelector
            label="Select Time"
            id="schedule"
            name="schedule"
            courseSchedule={courseSchedule}
            onChange={setSelectedSchedule}
            width="50%"
          />
        ) : (
          <p>Loading</p>
        )}
        <Label htmlFor="totalHours">Fill Duration</Label>
        <DurationInput
          type="number"
          id="totalHours"
          name="totalHours"
          value={totalHours}
          min="1"
          onChange={(e) => setTotalHours(e.target.value)}
        />
        <Wrapper>
          <TotalPrice>
            Total Price: {(totalHours * course.price_per_hour).toLocaleString()}{" "}
            Baht
          </TotalPrice>
          <Button type="submit">Enroll</Button>
        </Wrapper>
      </Form>
    </Container>
  );
}

const Container = styled.div`
  width: 100%;
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 50px 0px;
`;

const Form = styled.form`
  width: 80%;
  background-color: white;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border: 1px solid #858585;
  border-radius: 10px;
  padding: 20px;
`;

const Label = styled.label`
  margin-top: 20px;
`;

const DurationInput = styled.input`
  padding: 10px;
  width: 100px;
  height: 40px;
  border: 1px solid #ccc;
  border-radius: 4px;
  margin: 20px;
`;
const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
`;

const TotalPrice = styled.p`
  font-size: 20px;
  font-weight: bold;
  color: #eb4242;
  align-self: flex-start;
`;

const Button = styled.button`
  width: 400px;
  height: 40px;
  background-color: #eb4242;
  color: white;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  margin: 20px 0px;
`;
