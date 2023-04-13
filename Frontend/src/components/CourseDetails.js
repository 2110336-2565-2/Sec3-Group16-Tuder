import styled from "styled-components";
import React from "react";
import { getRole } from "../utils/jwtGet";

export default function CourseDetails(props) {
  const data = props.courseDetail;
  const button = () => {
    if (getRole() !== "admin" && getRole() !== "tutor") {
      return (
        <Content fsize="16px" fweight="400">
          <ButtonCourse>Enroll</ButtonCourse>
        </Content>
      );
    }
  };

  return (
    <>
      <DetailGrid>
        <GridContent>
          <DetailImgFix>
            <DetailImg src={data.course_picture_url} />
          </DetailImgFix>
        </GridContent>
        <GridContent>
          <DetailContent>
            <Detailrow fsize="30px" fweight="700">
              {data.title}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              {data.ReviewCount} review(s)
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Tutor : {`${data.TutorFirstname} ${data.TutorLastname}`}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Level : {data.level}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Description :
            </Detailrow>
            <Detailrow mgtop="0.2rem" fsize="16px" fweight="400">
              {data.description}
            </Detailrow>
            <Detailrow mgtop="1rem" fsize="16px" fweight="400">
              Estimate time : {data.estimate_time}
            </Detailrow>
            <DetailRight>
              <Content fsize="25px" fweight="400">
                {data.price_per_hour} Bath/hour
              </Content>
            </DetailRight>
            <DetailRight>{button()}</DetailRight>
          </DetailContent>
        </GridContent>
      </DetailGrid>
    </>
  );
}

const DetailRight = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  margin-top: 1rem;
`;

const Content = styled.div`
  display: flex;
  font-size: ${(props) => {
    return props.fsize;
  }};
  font-weight: ${(props) => {
    return props.fweight;
  }};
`;

const DetailGrid = styled.div`
  display: grid;
  grid-template-columns: 45% 55%;
  width: 100%;
  min-width: 750px;
`;
const GridContent = styled.div`
  display: grid;
`;
const ButtonCourse = styled.button`
  display: flex;
  width: 80px;
  padding: 5px;
  justify-content: center;
  text-decoration: none;
  border: 2px solid #ff7008;
  border-radius: 5px;
  color: #ffffff;
  background-color: #ff7008;

  &:hover {
    background-color: #ff7908;
    color: #ffffff;
  }

  cursor: pointer;
`;

const DetailImgFix = styled.div`
  margin-right: auto;
  margin-left: auto;
  margin-top: 1rem;
  width: 390px;
  height: 260px;
`;

const DetailImg = styled.img`
  display: flex;
  width: 100%;
  height: 100%;
  min-width: 270px;
  min-height: 180px;
  object-fit: cover;
  border-radius: 10px;
  object-fit: cover;
`;

const DetailContent = styled.div`
  display: flex;
  flex-direction: column;
  padding: 10px;
  min-width: 450px;
`;
const Detailrow = styled.div`
  display: flex;
  margin-top: ${(props) => {
    return props.mgtop;
  }};
  font-size: ${(props) => {
    return props.fsize;
  }};
  font-weight: ${(props) => {
    return props.fweight;
  }};
`;
