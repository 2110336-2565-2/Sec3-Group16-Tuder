import toast from "react-hot-toast";
import TimePicker from "react-time-picker";
import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { FormP } from "./ProfileStyle";

//icon
import { ClockCircleOutlined, DeleteOutlined } from "@ant-design/icons";

export default function TimeSelector({
  label,
  id,
  name,
  value,
  onChange,
  width,
}) {
  const availableTime = value;
  const days = [
    { day: "SUN", value: "Sunday" },
    { day: "MON", value: "Monday" },
    { day: "TUE", value: "Tuesday" },
    { day: "WED", value: "Wednesday" },
    { day: "THU", value: "Thursday" },
    { day: "FRI", value: "Friday" },
    { day: "SAT", value: "Saturday" },
  ];
  const [selectedDay, setSelectedDay] = useState("Sunday");
  const [timeSlot, setTimeSlot] = useState([]);
  // To be added timeSlot
  const [timeSlotToAdd, setTimeSlotToAdd] = useState({
    from: "00:00",
    to: "00:00",
  });
  const handleTimeSlotToAdd = (time, type) => {
    // mm must be 00
    if (time && time.split(":")[1] !== "00") {
      toast.error("Time must be in 00 minutes");
      // set mm to 00
      time = time.split(":")[0] + ":00";
      console.log(time);
    }
    if (type === "to") {
      // Check if to time is greater than from time
      if (timeSlotToAdd.from > time) {
        toast.error("To time must be greater than from time");
        return;
      }
      setTimeSlotToAdd({ ...timeSlotToAdd, to: time });
    } else {
      // Check if from time is less than to time
      if (timeSlotToAdd.to < time) {
        toast.error("From time must be less than to time");
        return;
      }
      setTimeSlotToAdd({ ...timeSlotToAdd, from: time });
    }
    console.log(timeSlotToAdd);
  };

  useEffect(() => {
    const timeSlot = availableTime.filter((day) => day.day === selectedDay);
    timeSlot.length > 0 ? setTimeSlot(timeSlot[0].timeSlot) : setTimeSlot([]);
  }, [selectedDay, availableTime]);

  return (
    <FormP.InputComponent width={width}>
      <Label>{label}</Label>
      <Container>
        <DaySelector>
          {days.map((day) => {
            return (
              <DayButton
                isSelected={selectedDay === day.value}
                onClick={() => setSelectedDay(day.value)}
                key={day.value}
              >
                {day.day}
              </DayButton>
            );
          })}
        </DaySelector>
        <TimeSlotContainer>
          {timeSlot.map((time) => {
            return (
              <TimeSlot key={time.from + " " + time.start}>
                <SlotWrapper>
                  <ClockIcon />
                  <span>
                    {time.from} - {time.to}
                  </span>
                </SlotWrapper>
                <DeleteIcon />
              </TimeSlot>
            );
          })}
        </TimeSlotContainer>
        <AddMenu>
          <Title>Add time slot</Title>
          <AddFormWrapper>
            <FormP.InputComponent>
              <FormP.Label>From</FormP.Label>
              <CustomTimePicker
                onChange={(time) => handleTimeSlotToAdd(time, "from")}
                value={timeSlotToAdd.from}
                clockIcon={null}
                clearIcon={null}
                disableClock={true}
                format="HH:mm"
              />
            </FormP.InputComponent>
            <FormP.InputComponent>
              <FormP.Label>To</FormP.Label>
              <CustomTimePicker
                onChange={(time) => handleTimeSlotToAdd(time, "to")}
                value={timeSlotToAdd.to}
                clockIcon={null}
                clearIcon={null}
                disableClock={true}
                format="HH:mm"
              />
            </FormP.InputComponent>
          </AddFormWrapper>
          <Button>Add</Button>
        </AddMenu>
      </Container>
    </FormP.InputComponent>
  );
}

const Label = styled(FormP.Label)`
  align-self: center;
  margin-bottom: 20px;
`;

const Container = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  align-self: center;
  border: 2px solid #858585;
  border-radius: 10px;
`;

const DaySelector = styled.div`
  width: 100%;
  display: flex;
  justify-content: space-between;
  padding: 30px 30px 20px 30px;
  border-bottom: 2px solid #858585;
`;

const DayButton = styled.button`
  background-color: transparent;
  border: none;
  outline: none;
  color: ${(props) => (props.isSelected ? "#224957" : "#D9D9D9")};
  cursor: pointer;
`;

const TimeSlotContainer = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
`;

const TimeSlot = styled.div`
  width: 90%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 0px;
  color: #224957;
  border-bottom: 2px solid #d9d9d9;
`;

const SlotWrapper = styled.div`
  display: flex;
  gap: 18px;
  align-items: center;
`;

const ClockIcon = styled(ClockCircleOutlined)`
  font-size: 20px;
  color: #224957;
`;

const DeleteIcon = styled(DeleteOutlined)`
  font-size: 20px;
  color: #224957;
`;

const AddMenu = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  border-top: 2px solid #858585;
  padding: 20px 30px;
  gap: 15px;
`;

const AddFormWrapper = styled.div`
  width: 100%;
  display: flex;
  justify-content: space-between;
`;
const CustomTimePicker = styled(TimePicker)`
  & > div {
    padding: 5px;
    border: 1px solid black;
    border-radius: 6px;
    background-color: white;
    color: black;
  }
`;

const Title = styled.h2`
  font-size: 20px;
  color: #224957;
`;

const Button = styled.button`
  align-self: center;
  width: 100px;
  height: 30px;
  background-color: transparent;
  border: 2px solid #224957;
  border-radius: 10px;
  outline: none;
  cursor: pointer;

  &:hover {
    background-color: #224957;
    color: white;
  }
`;
