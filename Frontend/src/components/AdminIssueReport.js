import React, { useState, useEffect } from "react";
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
    console.log(data);
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
  };

  const handlePrevState = () => {
    setStatusState((prevState) => {
      let newState;
      if (prevState === "ongoing") {
        newState = "rejected";
      } else if (prevState === "rejected") {
        newState = "completed";
      } else if (prevState === "completed") {
        newState = "ongoing";
      } else {
        newState = "ongoing";
      }
      return newState;
    });
  };

  const handleNextState = () => {
    setStatusState((prevState) => {
      let newState;
      if (prevState === "ongoing") {
        newState = "completed";
      } else if (prevState === "completed") {
        newState = "rejected";
      } else if (prevState === "rejected") {
        newState = "ongoing";
      } else {
        newState = "ongoing";
      }
      return newState;
    });
  };

  useEffect(() => {
    handleSave();
  }, [statusState]);

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
              <PrevStateButton value="false" onClick={handlePrevState}>
                Prev
              </PrevStateButton>
              <NextStateButton value="false" onClick={handleNextState}>
                Next
              </NextStateButton>
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
              <DeleteButton disabled={statusState != "ongoing" ? false : true}
                value="false" onClick={handleDelete}>
                Delete
              </DeleteButton>
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
    background-color: #EB7B42;
    border: 0.5px solid #DADADA ;
    color: white;
    min-height: 167px;
    gap: 20px;
    padding: 10px;
    border-radius: 10px;
    border: 1px solid black;
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

const PrevStateButton = styled.button`
  width: 80px;
  height: 35px;
  border: 2px solid #ff7008;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #FF7008;
  background-color: #FFFFFF;

  &:hover {
    background-color: #FF7008;
    color: #ffffff;
  }

  cursor: pointer;
`;

const NextStateButton = styled.button`
  width: 80px;
  height: 35px;
  border: 2px solid #ff7008;
  border-radius: 5px;

  font-family: "Lexend Deca";
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #FF7008;
  background-color: #FFFFFF;

  &:hover {
    background-color: #FF7008;
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
  color: #ffffff;
  background-color: #FF0000;

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