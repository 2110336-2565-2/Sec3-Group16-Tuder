
import React from 'react';
import styled from 'styled-components';
import { Timezone, DateFormat } from '../datas/DateFormat.js';
import { useNavigate } from 'react-router-dom';
import { submitAudittingHandler } from '../handlers/cancellingRequestHandler';
import { toast } from 'react-hot-toast';

export default function CancelRequestInfo(props){

    const navigate = useNavigate();

    const handleBack = () => {
        navigate('/cancel-requests');
    }

    const handleSubmit = (e) => {

        const data = {
            cancelRequestId: props.cancelRequestId,
            approve: (e.target.value) === "true"
        }

        

        submitAudittingHandler(data).then((res) => {
            if(res.data.success){
                toast.success('Success');
                navigate('/cancel-requests');
            }
            else{
                toast.error('Error');
            }
        }).catch((err) => {
            console.log(err);
        })
    }

    var submitButton = null;
    if (props.status !== "approved" && props.status !== "rejected") {
        submitButton = (
            <RightButtonSection>
                <RejectButton value="false" onClick={handleSubmit}>
                    Reject
                </RejectButton>
                <ApproveButton value="true" onClick={handleSubmit}>
                    Approve
                </ApproveButton>    
        </RightButtonSection>
        )
    }


    return (
        <ContainerWithHeight margintop='25px'>
                    <Title>
                        {props.title}
                    </Title>
                    <Image src={props.img} />
                    <ReportInfoSection>
                        <Reporter>
                            Reporter: {props.reporter}
                        </Reporter>
                        <DescriptionSection>

                            <DescriptionTopic>
                                Description: 
                            </DescriptionTopic>
                            <Description>
                                {props.description}
                            </Description>
                        </DescriptionSection>

                        <StatusSection>

                            <Reporter>
                                Report Date: {(new Date(props.report_date)).toLocaleString(Timezone, DateFormat) }
                            </Reporter>
                            <Status>
                                Status: {props.status}
                            </Status>
                        </StatusSection>

                        <ButtonSection>
                            <LeftButtonSection>
                                <BackButton onClick={handleBack}>
                                    Back
                                </BackButton>
                            </LeftButtonSection>
                            {submitButton}
                           
                        </ButtonSection>
                        
                    </ReportInfoSection>

                </ContainerWithHeight>
    )
}

const ContainerWithHeight = styled.div`
    display: flex;
    flex-direction: column;
    padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};
    justify-content: center;
    align-items: center;
    gap : 40px;
`;


const Title = styled.div`
    width: 625px;
    height: 44px;

    font-family: 'Lexend Deca';
    font-style: normal;
    font-weight: 700;
    font-size: 35px;
    line-height: 44px;
`
const Image = styled.img`
    width: 658px;
    height: 369px;
    object-fit: cover;
`

const ReportInfoSection = styled.div`
    display: flex;
    flex-direction: column;
    gap: 60px;
    width: 625px;
    height: 100%;
    justify-content: center;

`

const Reporter = styled.div`
    width: 150%;
    height: 10%;
    // margin-bottom: 20px;
`;

const DescriptionSection = styled.div`
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 625px;
    height: 10%;
`;

const Status = styled.div`
    width: 150%;
    height: 10%;
    text-align: right;
`;

const StatusSection = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    width: 100%;
    margin-bottom: 20px;
    justify-content: right;
`;


const DescriptionTopic = styled.div`
    width: 625px;
    height: 20%;
`;

const Description = styled.div`
    width: 625px;
`;


const ButtonSection = styled.div`
  display: flex;
  flex-direction: row;
  gap: 20px;
  width: 100%;
  margin-bottom: 20px;
  justify-content: flex-end; 

`;

const LeftButtonSection = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    width: 100%;
    margin-bottom: 20px;
    justify-content: flex-start;
`;

const RightButtonSection = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    width: 100%;
    margin-bottom: 20px;
    justify-content: flex-end;
`;

const BackButton = styled.button`

    width: 100px;
    height: 50px;
    border: 2px solid #FF7008;
    border-radius: 5px;

    font-family: 'Lexend Deca';
    font-style: normal;
    font-weight: 500;
    font-size: 16px;
    line-height: 20px;
    text-align: center;
    color: #FF7008;
    background-color: #FFFFFF;

    

    &:hover {
        background-color: #FF7008;
        color: #FFFFFF;
    }

    cursor: pointer;
`;


const RejectButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #FF0000;
  border-radius: 5px;

  font-family: 'Lexend Deca';
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #FF0000;
  background-color: #FFFFFF;

  &:hover {
    background-color: #FF0000;
    color: #FFFFFF;
  }

    cursor: pointer;

`;

const ApproveButton = styled.button`
  width: 100px;
  height: 50px;
  border: 2px solid #FF7008;
  border-radius: 5px;

  font-family: 'Lexend Deca';
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;
  text-align: center;
  color: #FFFFFF;
  background-color: #FF7008;

  &:hover {
    background-color: #FF8009;
    border: 2px solid #FF8008;
    color: #FFFFFF;
  }

  cursor: pointer;
`;

