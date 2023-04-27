import React from "react";
import { getRole } from "../utils/jwtGet.js";
import { useNavigate } from "react-router";
import styled from "styled-components";

import { Timezone, DateFormat } from "../datas/DateFormat.js";

export default function MyClass(props) {
  const role = getRole();
  const course_picture_url = props.data.course_picture_url;
  const course_name = props.data.course_name;
  const tutor_name = props.data.tutor_name;
  const navigate = useNavigate();
  const remaining = props.data.remaining;
  const status = props.data.status;
  // the first appointment in the array which is in state coming soon
  const upcoming =
    remaining != 0
      ? new Date(props.data.upcoming_class).toLocaleString(Timezone, DateFormat)
      : "No Upcoming Class";

  var button = null;
  if (status === "completed" && role === "student") {
    button = (
      <Button onClick={(e) => navigate(`/reviews/${props.data.course_id}`)}>
        Review Course
      </Button>
    );
  } else if (status !== "canceled" && status !== "cancelling" && status !== "completed") {
    button = (
      <Button onClick={(e) => navigate(`/classes/${props.data.match_id}`)}>
        View Info
      </Button>
    );
  } 

  return (
    <Request>
      <ClassSection>
        <GridSection>
          <ClassImg src={course_picture_url} alt="classImg" />
        </GridSection>

        <ClassInfoSection>
          <ClassFlex>
            <ClassInfo>
              <h3>{course_name}</h3>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w="43px">Tutor :</InfoTitle>
              <InfoDesc>{tutor_name}</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w="79px">Upcoming Class :</InfoTitle>
              <InfoDesc>{upcoming}</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w="105px">Remaining :</InfoTitle>
              <InfoDesc>{remaining} hrs</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w="105px">Status :</InfoTitle>
              <InfoDesc>{status}</InfoDesc>
            </ClassInfo>
          </ClassFlex>
        </ClassInfoSection>
        <GridSection>
          {/* put button in here */}
          <FlexButton>
            {button}
          </FlexButton>
        </GridSection>
      </ClassSection>
    </Request>
  );
}

const Request = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  background-color: white;
  border: 1px solid #000000;
  min-height: 167px;
  min-width: 1000px;
  gap: 20px;
  padding: 10px;
  border-radius: 10px;
  margin: 10px;
`;

const ClassSection = styled.div`
  display: grid;
  grid-template-columns: 30% 50% 20%;
  width: 100%;
`;

const GridSection = styled.div`
  display: grid;
  align-self: center;
  justify-self: center;
`;

const ClassImg = styled.img`
  max-width: 216px;
  height: 148px;
  margin: auto;
  border-radius: 10px;
`;

const ClassInfoSection = styled.div`
  gap: 30px;
  display: grid;
  padding: 10px;
  font-size: 30px;
  width: 100%;
  min-width: 500px;
`;

const ClassFlex = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 10px;
`;

const ClassInfo = styled.div`
  display: flex;
  flex-direction: row;
  gap: 5px;
  align-items: flex-start;
  font-weight: 350;
  font-size: 16px;
`;
const ClassInfoButton = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  font-weight: 350;
  font-size: 16px;
  gap: 5px;
`;

const InfoTitle = styled.div`
  display: flex;
  font-size: 14px;
  font-weight: 450;
  min-width: ${(props) => {
    return props.min_w;
  }};
`;

const InfoDesc = styled.div`
  display: flex;
  font-size: 14px;
  font-weight: 300;
`;

const Button = styled.button`
  border: 2px solid #ff8008;
  border-radius: 10px;
  padding: 10px;
  font-size: 16px;
  font-weight: 400;
  color:white;
  cursor: pointer;
  background-color: #ff7008;

  &:hover {
    background-color: #ff8009;
    border: 2px solid #ff8008;
    color: #ffffff;
  }
`;

const FlexButton = styled.div`
  display: flex;
`