import React from "react";
import styled from "styled-components"
import { submitApprovmentHandler, submitRejectionHandler } from "../handlers/cancellingClassHandler";
import {toast} from 'react-hot-toast';




export default function CancelRequest(props){

    // const submitApprove = (classId) => {

    //     // confirmation before submit
    //     if(!window.confirm("Are you sure to approve this request?")){
    //         return;
    //     }

    //     submitApprovmentHandler(classId).then((res) => {
    //         if(res.data.success){
    //             toast.success("Approve Successfully");
    //             //delete component
    //             props.removeItem(classId);
                
                
    //         }
    //     }).catch((err) => {
    //         console.log(err);
    //         toast.error("Approve Failed");
            
    //     }
    //     )
    // }

    // const submitReject = (classId) => {

    //     // confirmation before submit
        
    //     if(!window.confirm("Are you sure to reject this request?")){
    //         return;
    //     }

    //     submitRejectionHandler(classId).then((res) => {
    //         if(res.data.success){
    //             toast.success("Rejected Successfully",{iconTheme: {primary: '#DAA520', secondary: '#fff'}});
    //             //delete component
    //             props.removeItem(classId);
                
    //         }
    //     }).catch((err) => {
    //         console.log(err);
    //         toast.error("Reject Failed");
    //     }
    //     )
    // }


    // var button = () => {
    //     if (props.type !== "Topic") {
    //         return (
    //             <>
    //                     <ApproveButton onClick={(e) => submitApprove(props.classId)}>Approve</ApproveButton>
    //                     <RejectButton onClick={(e) => submitReject(props.classId)}>Reject</RejectButton>
    //                 </>
    //         )
    //     } else {
    //         return (
    //             <>
    //             <ApproveButton type="Topic">Approve</ApproveButton>
    //             <RejectButton type="Topic">Reject</RejectButton>

    //             </>
    //         )
    //     }
    // }

    return(
        <Request>
            <ClassSection>
                <ClassImg src={props.img} alt="classImg"/>

                <ClassInfoSection>

                    <ClassInfo>
                        <InfoTitle>Title:</InfoTitle>
                        &nbsp;
                        {props.title}
                    </ClassInfo>
                    <ClassInfo>
                        <InfoTitle>Course:</InfoTitle>
                        &nbsp;
                        {props.course}
                    </ClassInfo>
                    <ClassInfo>
                        <InfoTitle>Reporter:</InfoTitle>
                        &nbsp;
                        {props.reporter}
                    </ClassInfo>
                    <ClassInfo>
                        <InfoTitle>Report Date:</InfoTitle>
                        &nbsp;
                        {props.report_date}
                    </ClassInfo>
                    <ClassInfo>
                        <InfoTitle>Status:</InfoTitle>
                        &nbsp;
                        {props.status}
                    </ClassInfo>
                </ClassInfoSection>
            </ClassSection>
        </Request>
    )
}


const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: left;
    background-color: #EB7B42;
    width: 1000px;
    height: 167px;
    gap: 20px;
    padding: 10px;
    
`

const ClassSection = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: center;
    gap: 20px;
`;

const ClassImg = styled.img`
    width: 264px;
    height: 148px;
`;

const ClassInfoSection = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: left;
    gap: 10%;
    margin-top: 15px;
    margin-right: 10px;
    margin-left: 10px;
`

const ClassInfo = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: left;

    font-size: 20px;
    margin-top: -2.5%;
    margin-bottom: -2.5%;
`

const InfoTitle = styled.div`
    display: flex;
    font-size: 20px;
    font-weight: bold;
`;
