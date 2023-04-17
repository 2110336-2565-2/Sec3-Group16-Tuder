import React, { useEffect,useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import styled from "styled-components"

import Error from "../components/global/Error.js";
import Footer from "../components/global/Footer.js";
import Button from "../components/global/Button.js";

import ClassCard from "../components/tutor/ClassCard.js";

import { IsTutor } from "../components/IsAuth.js";
import { fetchStudent } from "../handlers/indivCourseHandler.js";

export default function IndivCourse(){

    const [data,setData] = useState([]);
    const [coursename,setCourseName] = useState();
    const courseid = useParams();

    useEffect(() => {
        const res = fetchStudent(courseid.id)
            .then((res) => {
                setData(res.data.data.data_card);
                setCourseName(res.data.data.course_name);
            })
            .catch((err) => {
                console.log(err)
                console.log("Course not found");
                setError(true);
            })
            }, []);


    return(
        <IsTutor>
            <Container>
                <Header>
                    <Title>{coursename}</Title>
                </Header>
                <ClassList>
                    {data===null||data===[]?(
                        <NoClassContainer>
                            <p>No class yet.</p>
                        </NoClassContainer>
                        ):(data.map((item,index) => 
                            <ClassCard data={item} key={index} />
                        ))
                    }
                </ClassList>
            </Container>
            <Footer />
        </IsTutor>
        )
 }

const Container = styled.div`
    position: relative;
    width: 100%;
    max-height: 800px; /* set your desired maximum height */
    height: auto;
    overflow-y: auto; /* add scrollbar if content exceeds max height */
    margin-bottom: 100px;
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

const ClassList = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
`;

const NoClassContainer = styled.div`
  display: flex;
  width: 100%;
  background-color: white;
  justify-content: center;
  align-items: center;
  min-height: 50vh;
`;