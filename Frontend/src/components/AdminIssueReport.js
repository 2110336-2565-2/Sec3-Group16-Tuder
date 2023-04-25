import React, { useState, useRef, useLayoutEffect } from "react";
import styled from "styled-components";
import {
  submitSaveStateHandler,
  submitDeleteStateHandler,
} from "../handlers/AdminIssueReportHandler";
import { toast } from "react-hot-toast";
import { Timezone, DateFormat } from "../datas/DateFormat.js";

export default function AdminIssueReport(props) {
  const [statusState, setStatusState] = useState(props.status);

  const handleSave = () => {
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
        console.log(err)
        toast.error("Something went wrong");
      });
  };
  
  const handleComplete = () => {
    setStatusState((prevState) => {
      let newState;
      if (prevState === "completed") {
        newState = "ongoing";
      }  else {
        newState = "completed";
      }
      return newState;
    });
  };

  const handleReject = () => {
    setStatusState((prevState) => {
      let newState;
      if (prevState === "rejected") {
        newState = "ongoing";
      }  else {
        newState = "rejected";
      }
      return newState;
    });
  };

  const firstUpdate = useRef(true);
  useLayoutEffect(() => {
    if (firstUpdate.current) {
      firstUpdate.current = false;
      return;
    }
    handleSave();

  });

  const isEnablebutton = () => {
    if (statusState === "ongoing") {
      return ;
    }
    return (
        <DeleteButton disabled={statusState != "ongoing" ? false : true}
          value="false" onClick={handleDelete} statusState = {statusState}>
            Delete
        </DeleteButton>
      );
  };

  const handleDelete = () => {
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
        toast.error("Something went wrong");
      });
    props.onDelete(props.issuereport_id);
  };

  return (
    <Request status={statusState}>
      <ClassSection>
        <ClassInfoSection>
          <ClassFlex>
            <ClassInfo>
              <InfoTitle min_w = "45px">Title :</InfoTitle>
              <InfoDesc>{props.title}</InfoDesc>              
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w = "99px">Description :</InfoTitle>
              <InfoDesc>{props.description}</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w = "72px">Contact :</InfoTitle>
              <InfoDesc>{props.contact}</InfoDesc>
            </ClassInfo>
            <ClassInfo>
              <InfoTitle min_w = "112px">Report_Date :</InfoTitle>
              <InfoDesc>{new Date(props.report_date).toLocaleString(Timezone, DateFormat)}</InfoDesc>
            </ClassInfo>
          </ClassFlex>
        </ClassInfoSection>

        <StateSection>
          <StateFlex>
            <StateInfo>
              <CompleteButton value="false" onClick={handleComplete} statusState = {statusState}>
                completed
              </CompleteButton>
              <RejectButton value="false" onClick={handleReject} statusState = {statusState}>
                rejected
              </RejectButton>
            </StateInfo>
            <StateInfo>
              <InfoTitle>Status :</InfoTitle>
              <InfoDesc>
                <StatusBlock statusState = {statusState}>
                  {statusState}
                </StatusBlock>
              </InfoDesc>
            </StateInfo>
            <StateInfo>
              {isEnablebutton()}
            </StateInfo>
            
          </StateFlex>
        </StateSection>
      </ClassSection>
    </Request>
  );
}


const Request = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: center;
    background-color: #EBEBEB;
    border: 0.5px solid #DADADA ;
    color: #000;
    min-height: 167px;
    gap: 20px;
    padding: 10px;
    border-radius: 10px;
    
`;

const ClassSection = styled.div`
    display: grid;
    grid-template-columns: 70% 30%;
    width: 100%;
    min-width: 750px;
`;

const ClassInfoSection = styled.div`
    gap: 30px;
    display: grid;
    padding: 10px;
    font-size: 30px;
    width: 100%;
    min-width: 500px;
`;

const StateSection = styled.div`
  display: grid;
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
    gap: 5px;
    align-items: flex-start;
    font-weight: 350;
    font-size: 16px;
`;
const StateInfo = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 5px;
  font-size: 20px;
`;

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
`

const CompleteButton = styled.button`
  width: 95px;
  height: 35px;
  border: 2px solid #10AA00;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  background-color: #009900;
  color: #FFFFFF;
  &:hover {
    background-color: #10AA00;
    color: #ffffff;
  }

  cursor: pointer;
`;

const RejectButton = styled.button`
  width: 95px;
  height: 35px;
  border: 2px solid #FF0F39;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  background-color: #FF0000;
  color: #FFFFFF;

  &:hover {
    background-color: #FF0F39;
    color: #ffffff;
  }

  cursor: pointer;
`;


const DeleteButton = styled.button`
  width: 80px;
  height: 35px;
  border: 2px solid #FF0000;
  border-radius: 5px;
  min-width: 160px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  background-color: ${(props) => {
    if (props.statusState === "ongoing") {
        return "#D3D3D3";
    } else {
        return "#FF0000";
    }
  }};
  color: ${(props) => {
    if (props.statusState === "ongoing") {
        return "#000000";
    }else {
        return "#FFFFFF";
    }
  }};

  &:hover {
    background-color: #FF2000;
    border: 2px solid #FF2000;
    color: #ffffff;
  }

  cursor: pointer;
`;


const ClassFlex = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 10px;
`

const StateFlex = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 10px;
`

const StatusBlock = styled.div`
    display: flex;
    padding: 7px;
    border-radius: 6px;
    background-color: ${(props) => {
        if (props.statusState === "rejected") {
            return "#FF0000";
        } else if (props.statusState === "completed") {
            return "#009900";
        } else {
            return "#FFFF00";
        }
    }};
    color: ${(props) => {
        if (props.statusState === "rejected") {
            return "#FFFFFF";
        } else if (props.statusState === "completed") {
            return "#FFFFFF";
        } else {
            return "#000000";
        }
    }};
    font-weight: 500;
`