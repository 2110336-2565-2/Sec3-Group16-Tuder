import React, {useState} from "react";
import styled from "styled-components";
import { submitSaveStateHandler,submitDeleteStateHandler } from "../handlers/AdminIssueReportHandler";
import { toast } from "react-hot-toast";
import { Timezone, DateFormat } from "../datas/DateFormat.js";

export default function AdminIssueReport(props) {
    const [statusState, setStatusState] = useState(
        statusState = props.status
    );
    const handleSave = (e) => {
        {
          const data = {
            IssueReportId: props.issuereport_id,
            status: statusState,
          };
          submitSaveStateHandler(data)
            .then((res) => {
              if (res.data.success) {
                toast.success("Success");
              } else {
                toast.error("Error");
              }
            })
            .catch((err) => {
              console.log(err);
            });
        }
      };
  const handlePrevState = (e) => {
    if(statusState == "ongoing" ){
        setStatusState("rejected")
    }else if(statusState == "rejected"){
        setStatusState("completed")
    }else if(statusState == "completed"){
        setStatusState("ongoing")
    }else{
        setStatusState("ongoing")
    }
    handleSave()
  };
  const handleNextState = (e) => {
    
    if(statusState == "ongoing" ){
        setStatusState("completed")
    }else if(statusState == "completed"){
        setStatusState("rejected")
    }else if(statusState == "rejected"){
        setStatusState("ongoing")
    }else{
        setStatusState("ongoing")
    }
    handleSave()
  };



  const handleDelete = (e) => {
    {
      const data = {
        IssueReportId: props.issuereport_id,
      };
      submitDeleteStateHandler(data)
        .then((res) => {
          if (res.data.success) {
            toast.success("Success");
          } else {
            toast.error("Error");
          }
        })
        .catch((err) => {
          console.log(err);
        });
        props.onDelete(props.issuereport_id);
    }
  };
  return (
    <Request status={statusState}>
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

        <StateSection>
          <PrevStateButton value="false" onClick={handlePrevState}>
            Prev
          </PrevStateButton>
          <StateInfo>
            <InfoTitle>Status:</InfoTitle>
            &nbsp;
            {statusState}
          </StateInfo>
          <NextStateButton value="false" onClick={handleNextState}>
            Next
          </NextStateButton>
        </StateSection>

        <ButtonSection>
          <DeleteButton disabled={statusState!="ongoing" ? false:true} value="false" onClick={handleDelete}>
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

const StateSection = styled.div`
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
const StateInfo = styled.div`
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
const PrevStateButton = styled.button`
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

const NextStateButton = styled.button`
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

const DeleteButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #b22222;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #ffffff;
  background-color: #b22222;

  &:hover {
    background-color: #dc143c;
    border: 2px solid #dc143c;
    color: #ffffff;
  }

  cursor: pointer;
`;
