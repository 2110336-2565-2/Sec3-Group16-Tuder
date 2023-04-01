import React from "react";
import styled from "styled-components"
import { submitAudditing } from "../handlers/cancellingRequestHandler";
import {toast} from 'react-hot-toast';
import { Timezone, DateFormat } from '../datas/DateFormat.js';





export default function CancelRequest(props){


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
                        <InfoTitle>Reporter:</InfoTitle>
                        &nbsp;
                        {props.reporter}
                    </ClassInfo>
                    <ClassInfo>
                        <InfoTitle>Report Date:</InfoTitle>
                        &nbsp;
                        {(new Date(props.report_date)).toLocaleString(Timezone, DateFormat)}
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

// color be changed depends on the status

const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: left;
    background-color: #EB7B42;

    width: 1000px;
    height: 167px;
    gap: 20px;
    padding: 10px;
    cursor: pointer;

    &:hover {
        background-color: #DAA520;
    }


    
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
    gap: 30px;
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
