import React from "react";
import styled from "styled-components";
import { fetchAdminIssueReportHandler } from "../handlers/AdminIssueReportHandler";
import { toast } from "react-hot-toast";
import { Timezone, DateFormat } from "../datas/DateFormat.js";

export default function AdminIssueReport(props) {
  const handlePrevStage = (e) => {};
  const handleNextStage = (e) => {};
  const handleSave = (e) => {};
  const handleDelete= (e) => {};ç
  return (
    <Request status={props.status}>
      <ClassSection>
        <ClassInfoSection>
          <ClassInfo>
            <InfoTitle>Title:</InfoTitle>
            &nbsp;
            {props.title}
          </ClassInfo>
          <ClassInfo>
            <InfoTitle>description:</InfoTitle>
            &nbsp;
            {props.description}
          </ClassInfo>
          <ClassInfo>
            <InfoTitle>Contract:</InfoTitle>
            &nbsp;
            {props.contact}
          </ClassInfo>
          <ClassInfo>
            <InfoTitle>Report_Date:</InfoTitle>
            &nbsp;
            {new Date(props.report_date).toLocaleString(Timezone, DateFormat)}
          </ClassInfo>
        </ClassInfoSection>

        <StageSection>
          <PrevStageButton value="false" onClick={handlePrevStage}>
            Prev
          </PrevStageButton>
          <StageInfo>
            <InfoTitle>Status:</InfoTitle>
            &nbsp;
            {props.status}
          </StageInfo>
          <NextStageButton value="true" onClick={handleNextStage}>
            Next
          </NextStageButton>
        </StageSection>

        <ButtonSection>
          <SaveButton value="false" onClick={handleSave}>
            Save
          </SaveButton>

          <DeleteButton value="true" onClick={handleDelete}>
            Delete
          </DeleteButton>
        </ButtonSection>
      </ClassSection>
    </Request>
  );
}

// color be changed depends on the status

const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: left;
    background-color: ${(props) => {
      if (props.status === "rejected") {
        return "#EBEBEB";
      } else if (props.status === "approved") {
        return "#009900";
      } else {
        return "#EB7B42";
      }
    }};

    border-radius: 10px;
    border: 1px solid #DADADA ;
    shadow: 0px 4px 4px rgba(0, 0, 0, 0.45);

    width: 1000px;
    height: 167px;
    gap: 20px;
    padding: 10px;
    cursor: pointer;

    &:hover {
        background-color: ${(props) => {
          if (props.status === "rejected") {
            return "#FBFBFB";
          } else if (props.status === "approved") {
            return "#00AA00";
          } else {
            return "#EE8E45";
          }
        }}


    
`;

const ClassSection = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 20px;
`;

const ClassInfoSection = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: left;
  gap: 20px;
  margin-top: 15px;
  margin-right: 10px;
  margin-left: 10px;
`;

const StageSection = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: left;
  gap: 20px;
  margin-top: 15px;
  margin-right: 10px;
  margin-left: 10px;
`;

const ButtonSection = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: left;
  gap: 20px;
  margin-top: 15px;
  margin-right: 10px;
  margin-left: 10px;
`;
const ClassInfo = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: left;

  font-size: 10px;
  margin-top: -2.5%;
  margin-bottom: -2.5%;
`;
const StageInfo = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: left;

  font-size: 20px;
  margin-top: 2.5%;
  margin-bottom: 2.5%;
`;

const InfoTitle = styled.div`
  display: flex;
  font-size: 10px;
  font-weight: bold;
`;
const PrevStageButton = styled.button`
width: 100px;
height: 50px;
border: 2px solid #ff7008;
border-radius: 5px;

font-family: "Lexend Deca";
font-style: normal;
font-weight: 500;
font-size: 16px;
line-height: 20px;
text-align: center;
color: #ffffff;
background-color: #ff7008;

&:hover {
  background-color: #ff8009;
  border: 2px solid #ff8008;
  color: #ffffff;
}

cursor: pointer;
`;

const NextStageButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #ff7008;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #ffffff;
  background-color: #ff7008;

  &:hover {
    background-color: #ff8009;
    border: 2px solid #ff8008;
    color: #ffffff;
  }

  cursor: pointer;
`;
const SaveButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #808000;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #ffffff;
  background-color: #808000;

  &:hover {
    background-color: #008000;
    border: 2px solid #008000;
    color: #ffffff;
  }

  cursor: pointer;
`;
const DeleteButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #B22222;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #ffffff;
  background-color: #B22222;
 

  &:hover {
    background-color: #DC143C;
    border: 2px solid #DC143C;
    color: #ffffff;
  }

  cursor: pointer;
`;
