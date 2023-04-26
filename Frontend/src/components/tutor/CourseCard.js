import React, { useEffect,useState } from "react";
import styled from "styled-components";
import Button from "../global/Button.js";
import { changeStatus } from "../../handlers/changeCourseStatus.js";
import { confirm } from "../global/ConfirmToast.js";
import { Link } from "react-router-dom";
import { toast } from "react-hot-toast";

export default function CourseCard({course,buttonStatus,setEventTrigger}) {

    const [status, setStatus] = useState(buttonStatus)

    const handleClick = () => {
        const data = status === "Open" ? {status: "closed"} : {status: "open"}
        const res = changeStatus(course.id,data)
            .then((res) => {               
                setStatus(res.data.data.status==="open" ? "Close" : "Open")
                setEventTrigger(prevState => !prevState)
            })
            .catch((err) => {
                console.log(err)
                toast.error("Something went wrong");
            });
      };

    useEffect(() => {
        setStatus(course.status==="open" ? "Close" : "Open")
    }, [course]);

    const route = `/my-course/${course.id}`

    return (
    <Request>
        <CourseSection to={route}>
          <ItemSection>
            <CourseImg src={course.course_picture_url} alt="CourseImg" />
          </ItemSection>
  
          <CourseInfoSection>
            <CourseFlex>
              <CourseInfo>
                <h3>{course.title}</h3>
              </CourseInfo>
              <CourseInfo>
                <InfoTitle min_w="43px">Level :</InfoTitle>
                <InfoDesc>{course.level}</InfoDesc>
              </CourseInfo>
  
              <CourseInfo>
                <InfoTitle min_w="79px">Estimated time :</InfoTitle>
                <InfoDesc>{course.estimated_time}</InfoDesc>
              </CourseInfo>
              <CourseInfo>
                <InfoTitle min_w="45px">Status :</InfoTitle>
                <InfoDesc>{course.status}</InfoDesc>
              </CourseInfo>
              <CourseInfoButton>
                <InfoTitle min_w="60px">Number of classes :</InfoTitle>
                <InfoDesc>{course.num_of_class}</InfoDesc>
              </CourseInfoButton>
            </CourseFlex>
          </CourseInfoSection>
        </CourseSection>
        <ItemSection>
            {/* put button in here */}
            <StatusButton onClick={()=>
            confirm(`Are you sure to ${status} ${course.title}?`
        , () => handleClick()
      )}>{status}</StatusButton>
        </ItemSection>
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
  margin: 20px 50px 20px 50px;
`;

const CourseSection = styled(Link) `
  text-decoration: none;
  color:black;
  display: grid;
  grid-template-columns: 30% 40% 30%;
  width: 100%;
`;

const ItemSection = styled.div`
  align-items: flex-start;
  display: flex;
  margin: 10px 40px 10px 0px ;
`;

const CourseImg = styled.img`
  width: 216px;
  height: 148px;
  margin: auto;
  border-radius: 10px;
`;

const CourseInfoSection = styled.div`
  gap: 30px;
  display: grid;
  padding: 10px;
  font-size: 30px;
  width: 100%;
  min-width: 500px;
`;

const CourseFlex = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 10px;
`;

const CourseInfo = styled.div`
  display: flex;
  flex-direction: row;
  gap: 5px;
  align-items: flex-start;
  font-weight: 350;
  font-size: 16px;
`;
const CourseInfoButton = styled.div`
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

const StatusButton = styled(Button)`
    width: 200px;
    align-text: center;
`;

