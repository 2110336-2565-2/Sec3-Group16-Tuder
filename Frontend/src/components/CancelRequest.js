import React from "react";
import styled from "styled-components"
import { submitAudditing } from "../handlers/cancellingRequestHandler";
import {toast} from 'react-hot-toast';
import { Timezone, DateFormat } from '../datas/DateFormat.js';





export default function CancelRequest(props){


    return(
        <Request status={props.status}>
            <ClassSection>
                <GridSection>
                    <ClassImg src={props.img} alt="classImg"/>
                </GridSection>
                

                <ClassInfoSection>
                    <ClassFlex>
                        <ClassInfo>
                            <InfoTitle min_w = "43px">Title :</InfoTitle>
                            <InfoDesc>{props.title}</InfoDesc>
                        </ClassInfo>
                        
                        <ClassInfo>
                            <InfoTitle min_w = "79px">Reporter :</InfoTitle>
                            <InfoDesc>{props.reporter}</InfoDesc>
                        </ClassInfo>
                        <ClassInfo>
                            <InfoTitle min_w = "105px">Report Date :</InfoTitle>
                            <InfoDesc>{(new Date(props.report_date)).toLocaleString(Timezone, DateFormat)}</InfoDesc>
                        </ClassInfo>
                        <ClassInfoButton>
                            <InfoTitle min_w = "60px">Status :</InfoTitle>
                            <InfoDesc>
                                <StatusBlock>
                                    {props.status}
                                </StatusBlock>
                            </InfoDesc>
                            
                        </ClassInfoButton>
                    </ClassFlex>
                    
                </ClassInfoSection>
            </ClassSection>
        </Request>
    )
}

// color be changed depends on the status
// const MgAuto = styled.div`
//     margin-top: auto;
//     margin-bottom: auto;
// `

const StatusBlock = styled.div`
    display: flex;
    padding: 7px;
    border-radius: 6px;
    background-color: ${(props) => {
        if (props.status === "rejected") {
            return "#FF0000";
        } else if (props.status === "approved") {
            return "#009900";
        } else {
            return "#FFFF00";
        }
    }};
    color: ${(props) => {
        if (props.status === "rejected") {
            return "#FFFFFF";
        } else if (props.status === "approved") {
            return "#FFFFFF";
        } else {
            return "#000000";
        }
    }};
    font-weight: 500;
`
const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: center;
    background-color: #EB7B42;
    border: 1px solid #DADADA ;
    color: white;
    min-height: 167px;
    gap: 20px;
    padding: 10px;
    cursor: pointer;
    border-radius: 10px;
    &:hover {
        border: 1px solid #000000 ;
        background-color: #EE8E45 ;
    }
`

const ClassSection = styled.div`    
    display: grid;
    grid-template-columns: 30% 70%;
    width: 100%;
    min-width: 750px;
    
`;

const GridSection = styled.div`
    display: grid;
`

const ClassImg = styled.img`
    min-width: 216px;
    height: 148px;
    margin: auto;
    border-radius: 10px;
`;

const ClassInfoSection = styled.div`
    gap: 30px;
    display: grid;
    padding: 10px;
    font-size: 30px;
    width: 100%;
    min-width: 500px;
`

const ClassFlex = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 10px;
`

const ClassInfo = styled.div`
    display: flex;
    flex-direction: row;
    gap: 5px;
    align-items: flex-start;
    font-weight: 350;
    font-size: 16px;
`
const ClassInfoButton = styled.div`
    display: flex;
    flex-direction: row;
    align-items: center;
    font-weight: 350;
    font-size: 16px;
    gap: 5px;
`

const InfoTitle = styled.div`
    display: flex;
    font-size: 16px;
    font-weight: 500;
    min-width: ${(props)=>{
        return props.min_w;
    }};
`;

const InfoDesc = styled.div`
    display: flex;
    font-size: 16px;
    font-weight: 300;
`;

