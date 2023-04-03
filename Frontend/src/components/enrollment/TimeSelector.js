import toast from "react-hot-toast";
import TimePicker from "react-time-picker";
import React, { useState, useEffect } from "react";
import styled from "styled-components";
import { FormP } from "../profile/ProfileStyle";
import {
  adjustFormat,
  isOverlapped,
  getMergedTimeSlot,
} from "../../utils/profile/timeHandler.js";
import { convertBackendSchedulesToFrontend } from "../../utils/profile/scheduleConverter";

//icon
import { ClockCircleOutlined, DeleteOutlined } from "@ant-design/icons";

export default function TimeSelector({
  label,
  id,
  name,
  courseSchedule,
  onChange,
  width,
}) {
  const days = [
    { day: "SUN", value: "Sunday" },
    { day: "MON", value: "Monday" },
    { day: "TUE", value: "Tuesday" },
    { day: "WED", value: "Wednesday" },
    { day: "THU", value: "Thursday" },
    { day: "FRI", value: "Friday" },
    { day: "SAT", value: "Saturday" },
  ];

  const [selectedSchedule, setSelectedSchedule] = useState(() => {
    // create initialSelectedSchedule that add isSelect property to each timeSlot
    const initialSelectedSchedule = courseSchedule.map((day) => ({
      day: day.day,
      timeSlot: day.timeSlot.map((timeSlot) => ({
        ...timeSlot,
        isSelected: false,
      })),
    }));
    return initialSelectedSchedule;
  });

  const [selectedDay, setSelectedDay] = useState("Sunday");
  function handleCheckBoxChange(from, to) {
    // find the timeSlot that is selected
    const selectedTimeSlot = selectedSchedule
      .filter((day) => day.day === selectedDay)[0]
      .timeSlot.filter((timeSlot) => timeSlot.from === from && timeSlot.to === to)[0];
    // change the isSelected property of the timeSlot
    selectedTimeSlot.isSelected = !selectedTimeSlot.isSelected;
    // set the selectedSchedule
    setSelectedSchedule([...selectedSchedule]);
    onChange(selectedSchedule);
  }
  console.log(selectedSchedule, "selectedSchedule");
  return (
    <FormP.InputComponent width={width}>
      <Label>{label}</Label>
      <Container>
        <DaySelector>
          {days.map((day) => {
            return (
              <DayButton
                type="button"
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
          {
            //select schedule which day is selected
            selectedSchedule
              .filter((day) => day.day === selectedDay)[0]
              .timeSlot.map((time) => {
                return (
                  <TimeSlot key={time.from + " " + time.to}>
                    <SlotWrapper>
                      <ClockIcon />
                      <span>
                        {time.from} - {time.to}
                      </span>
                      <CheckBox
                        type="checkbox"
                        checked={time.isSelected}
                        onChange={()=>handleCheckBoxChange(time.from, time.to)}
                      />
                    </SlotWrapper>
                  </TimeSlot>
                );
              })
          }
        </TimeSlotContainer>
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
`;

const SlotWrapper = styled.div`
  width: 100%;
  display: flex;
  gap: 25px;
  align-items: center;
  justify-content: center;
`;

const ClockIcon = styled(ClockCircleOutlined)`
  font-size: 20px;
  color: #224957;
`;

const CheckBox = styled.input`
    margin-right: 5px;
    accent-color: #EB7B42;
    width: 15px;
    height: 15px;
`