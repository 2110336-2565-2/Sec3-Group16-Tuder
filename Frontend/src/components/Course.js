import styled from 'styled-components';
import React from 'react';
import { Link } from 'react-router-dom';

export default function Course(props){

    const pathCourse = "/courses/" + props.course_id

    return (
        <>
            <CardHeader>
                <CardImg src={props.img}/>
                <br></br>
                <CardContent fsize="16px">
                    {props.title}
                </CardContent>
                <CardContent fsize="16px">
                    TOPIC : {props.topic}
                </CardContent>  
                <CardContent fsize="16px">
                    SUBJECT : {props.subject}
                </CardContent>  
                <CardContent fsize="12px">
                    Time: {props.time} hr
                </CardContent>
                <HorizonL></HorizonL>  
                <CardContent fsize="17px">
                    <Inrow>
                        <div>{props.tutor}</div>
                        <div>{props.price} Bath/hr</div>
                    </Inrow>
                </CardContent>
                <CardContentLink fsize="16px">
                    <LinkCourse to={pathCourse}>View</LinkCourse>
                </CardContentLink>    
            </CardHeader>
        </>
    )
}
const CardHeader = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
`

const CardContentLink = styled.div`
    margin-top: 10px;
    display: flex;
    justify-content: flex-end;
    font-size: ${(props) => {
        return props.fsize
    }};
`
const LinkCourse = styled(Link)`
    display: flex;
    padding: 5px;
    text-decoration: none;
    border: 2px solid #ff7008;
    border-radius: 5px;
    font-style: normal;
    font-weight: 500;
    color: #FFFFFF;
    background-color: #FF7008;

    &:hover {
        background-color: #FF7908;
        color: #ffffff;
    }

    cursor: pointer;
`

const HorizonL = styled.hr`
    margin-top: 10px;
`

const CardImg = styled.img`
    width: 270px;
    height: 180px;
    object-fit: cover;
    border-radius: 10px;
    margin: auto;
    object-fit: cover;
`

const CardContent = styled.div`
    margin-top: 10px;
    display: flex;
    font-size: ${(props) => {
        return props.fsize
    }};
`

const Inrow = styled.div`
    display: flex;
    width: 270px;
    justify-content: space-between;
`