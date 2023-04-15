import React from "react";
import styled from "styled-components";
import { submitAudditing } from "../handlers/cancellingRequestHandler";
import { toast } from "react-hot-toast";
import { Timezone, DateFormat } from "../datas/DateFormat.js";

export default function AdminTuitionFee(props) {
  return (
    <Request>
      <ClassSection>
        <GridSection>
          <ClassImg src={props.Img} alt="classImg" />
        </GridSection>

        <ClassInfoSection>
          <ClassFlex>

          <ClassInfo>
              <InfoTitle min_w = "45px">Title :</InfoTitle>
              <InfoDesc>{props.Title}</InfoDesc>              
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w = "99px">Description :</InfoTitle>
              <InfoDesc>{props.Subject}</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w = "72px">Contact :</InfoTitle>
              <InfoDesc>{props.Topic}</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w = "112px">Report_Date :</InfoTitle>
              <InfoDesc>{new Date(props.AppointmentID).toLocaleString(Timezone, DateFormat)}</InfoDesc>
            </ClassInfo>
      

          </ClassFlex>
        </ClassInfoSection>
      </ClassSection>
    </Request>
  );
}

// color be changed depends on the status

const StatusBlock = styled.div`
  display: flex;
  padding: 7px;
  border-radius: 6px;
  background-color: ${(props) => {
    if (props.state === "rejected") {
      return "#FF0000";
    } else if (props.state === "approved") {
      return "#009900";
    } else {
      return "#FFFF00";
    }
  }};
  color: ${(props) => {
    if (props.state === "rejected") {
      return "#FFFFFF";
    } else if (props.state === "approved") {
      return "#FFFFFF";
    } else {
      return "#000000";
    }
  }};
  font-weight: 500;
`;
const Request = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  background-color: #eb7b42;
  border: 1px solid #dadada;
  color: white;
  min-height: 167px;
  gap: 20px;
  padding: 10px;
  cursor: pointer;
  border-radius: 10px;
  &:hover {
    box-shadow: 5px 6px 11px -3px rgba(0, 0, 0, 0.81);
    -webkit-box-shadow: 5px 6px 11px -7px rgba(0, 0, 0, 0.81);
    -moz-box-shadow: 5px 6px 11px -7px rgba(0, 0, 0, 0.81);
    background-color: #ee8e45;
  }
`;

const ClassSection = styled.div`
  display: grid;
  grid-template-columns: 30% 70%;
  width: 100%;
  min-width: 750px;
`;

const GridSection = styled.div`
  display: grid;
`;

const ClassImg = styled.img`
  max-width: 216px;
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
`;

const ClassFlex = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 10px;
`;

const ClassInfo = styled.div`
  display: flex;
  flex-direction: row;
  gap: 5px;
  align-items: flex-start;
  font-weight: 350;
  font-size: 16px;
`;
const ClassInfoButton = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  font-weight: 350;
  font-size: 16px;
  gap: 5px;
`;

const InfoTitle = styled.div`
  display: flex;
  font-size: 16px;
  font-weight: 500;
  min-width: ${(props) => {
    return props.min_w;
  }};
`;

const InfoDesc = styled.div`
  display: flex;
  font-size: 16px;
  font-weight: 300;
`;
