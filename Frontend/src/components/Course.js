import styled from 'styled-components';
import React from 'react';

export default function Course(props){
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
            </CardHeader>
        </>
    )
}
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

const CardHeader = styled.div`
    display: flex;
    justify-content: center;
    flex-direction: column;
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

const Fdiv = styled.div`
    font-size: ${(props) => {
        return props.fsize
    }};
`