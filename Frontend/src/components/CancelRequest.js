import React from "react";
import styled from "styled-components"
import { submitApprovmentHandler, submitRejectionHandler } from "../handlers/cancellingClassHandler";
import {toast} from 'react-hot-toast';




export default function CancelRequest(props){

    const submitApprove = (classId) => {

        // confirmation before submit
        if(!window.confirm("Are you sure to approve this request?")){
            return;
        }

        submitApprovmentHandler(classId).then((res) => {
            if(res.data.success){
                toast.success("Approve Successfully");
                //delete component
                props.removeItem(classId);
                
                
            }
        }).catch((err) => {
            console.log(err);
            toast.error("Approve Failed");
            
        }
        )
    }

    const submitReject = (classId) => {

        // confirmation before submit
        
        if(!window.confirm("Are you sure to reject this request?")){
            return;
        }

        submitRejectionHandler(classId).then((res) => {
            if(res.data.success){
                toast.success("Rejected Successfully",{iconTheme: {primary: '#DAA520', secondary: '#fff'}});
                //delete component
                props.removeItem(classId);
                
            }
        }).catch((err) => {
            console.log(err);
            toast.error("Reject Failed");
        }
        )
    }


    var button = () => {
        if (props.type !== "Topic") {
            return (
                <>
                        <ApproveButton onClick={(e) => submitApprove(props.classId)}>Approve</ApproveButton>
                        <RejectButton onClick={(e) => submitReject(props.classId)}>Reject</RejectButton>
                    </>
            )
        } else {
            return (
                <>
                <ApproveButton type="Topic">Approve</ApproveButton>
                <RejectButton type="Topic">Reject</RejectButton>

                </>
            )
        }
    }

    return(
        <Request>
            <ClassInfo>{props.classId}</ClassInfo>
            <ClassInfo>{props.course}</ClassInfo>
            <ClassInfo>{props.tutor}</ClassInfo>
            <ClassInfo>{props.student}</ClassInfo>
            <ClassInfo>{props.subject}</ClassInfo>
            <ClassInfo>{props.total_hour}</ClassInfo>
            <ClassInfo>{props.success_hour}</ClassInfo>
            <ClassInfo>{props.price}</ClassInfo>
            {button()}


        </Request>
    )
}


const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: left;
    gap: 20px;
    border: 1px solid black;
    border-radius: 16px;
    padding: 10px;
    
`

const ApproveButton = styled.button`
    visibility: ${props => props.type === "Topic" ? "hidden" : "visible"};
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
    visibility: ${props => props.type === "Topic" ? "hidden" : "visible"};
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
    justify-content: center;
    width: 150px;   
    height: 100px;
    gap: 10px;
    font-size: 20px;
    margin-top: 15px;
    margin-right: 10px;
`
