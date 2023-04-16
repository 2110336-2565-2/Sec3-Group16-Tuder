import React, { useEffect,useState } from "react";
import { useNavigate } from "react-router-dom";
import styled from "styled-components"
import { toast } from "react-hot-toast";
import Error from "../components/global/Error.js";
import Footer from "../components/global/Footer.js";
import Button from "../components/global/Button.js";

import CourseCard from "../components/tutor/CourseCard.js";

import { IsTutor } from "../components/IsAuth.js";
import { fetchTutorCourses } from "../handlers/tutorCoursesHandler.js";

export default function TutorCourses(){

    const navigate = useNavigate();
    const [error, setError] = useState(false);
    const [courses, setCourses] = useState([]);
    const [eventTrigger, setEventTrigger] = useState(false);

    useEffect(() => {
        const res = fetchTutorCourses()
            .then((res) => {
                setCourses(res.data.data);
            })
            .catch((err) => {
                console.log(err)
                toast.error("Something went wrong");
                setError(true);
            });

    }, [eventTrigger]);

    if (error) {
        return (
          <>
            <Error error_msg="Sorry, this tutor does not exist." />
            <Footer />
          </>
        );
      } else {

        return(
            <IsTutor>
                <Container>
                    <Header>
                        <Title>My Courses</Title>
                    </Header>
                    <CourseList>
                        {courses.map((course, index) => {
                            if (course.status === 'open') {
                                return <CourseCard course={course} key={index} buttonStatus="Close" setEventTrigger={setEventTrigger}/>
                            } else {
                                return <CourseCard course={course} key={index} buttonStatus="Open" setEventTrigger={setEventTrigger}/>
                            }
                            })}
                    </CourseList>
                </Container>
                <Wrapper>
                    <AddButton onClick ={()=> navigate("/add-course")}>Add New Course</AddButton>
                </Wrapper>
                <Footer />
            </IsTutor>
        )
    }
}

const Container = styled.div`
    position: relative;
    width: 100%;
    max-height: 800px; /* set your desired maximum height */
    height: auto;
    overflow-y: auto; /* add scrollbar if content exceeds max height */
`;


const Title = styled.h1`
    color: black;
`;

const Header = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 150px;
`;

const Wrapper = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
`;

const AddButton = styled(Button)`
    bottom: 0;
    margin: 50px;
    width: 250px;
    font-size: 16px;
    font-weight: 500;
`;

const CourseList = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
`;
