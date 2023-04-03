import React, {useState} from "react";
import styled from "styled-components"
import Footer from "../components/global/Footer.js";
import CourseDetails from "../components/CourseDetails.js";


export default function CourseDetail(){

    return(
        <>
            <Container>
                <CourseDetails />
            </Container>
            <Footer />
        </>
        
    )
}

const Container = styled.div`
    margin-top: 3rem;
    margin-bottom: 3rem;
    padding-left: 7%;
    padding-right: 7%;
`