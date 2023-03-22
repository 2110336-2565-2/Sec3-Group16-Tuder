import React from "react";
import styled from "styled-components"



export default function CancelRequest(props){

    return(
        <Request>
            <ClassInfo>{props.classId}</ClassInfo>
            <ClassInfo>{props.course}</ClassInfo>
            <ClassInfo>{props.tutor}</ClassInfo>
            <ClassInfo>{props.subject}</ClassInfo>
            <ClassInfo>{props.time}</ClassInfo>
            <ClassInfo>{props.price}</ClassInfo>

            <ApproveButton>Approve</ApproveButton>
            <RejectButton>Reject</RejectButton>


        </Request>
    )
}


const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: center;
    gap: 20px;
    border: 1px solid black;
    border-radius: 16px;
    padding: 10px;
    
`

const ApproveButton = styled.button`
    background-color: #4CAF50;
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 4px 2px;
    cursor: pointer;
    border-radius: 16px;
    &:hover {
        background-color: rgb(10, 200, 0);
    }
`

const RejectButton = styled.button`
    background-color: #f44336;
    border: none;
    color: white;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    margin: 4px 2px;
    cursor: pointer;
    border-radius: 16px;
    &:hover {
        background-color: rgb(260, 10, 0);
    }
`

const ClassInfo = styled.div`
    display: flex;
    flex-direction: column;
    gap: 10px;
    font-size: 20px;
    margin-top: 15px;
    margin-right: 10px;
`
