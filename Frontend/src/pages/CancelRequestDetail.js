import React from "react";
import styled from "styled-components"
import Footer from "../components/global/Footer.js";
import { IsAdmin } from "../components/IsAuth.js";

export default function CancelRequestDetail(){

    return(
        <IsAdmin>
            <Container>
                <ContainerWithHeight margintop='100px'>
                    <Title>
                        Chi tiết yêu cầu hủy
                    </Title>
                    <Image src="https://i.imgur.com/4ZQ9Z9C.png" />
                    <ReportInfoSection>
                        <Reporter>
                            Reporter: Nguyễn Văn A
                        </Reporter>
                        <DescriptionSection>

                            <DescriptionTopic>
                                Description: 
                            </DescriptionTopic>
                            <Description>
                                 Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                            </Description>
                        </DescriptionSection>

                        <Reporter>
                            Report Date: 20 March 2021
                        </Reporter>
                        <ButtonSection>
                            <RejectButton > 
                                Reject
                            </RejectButton>
                            <ApproveButton >
                                Approve
                            </ApproveButton>
                        </ButtonSection>
                        
                    </ReportInfoSection>

                </ContainerWithHeight>
                <Footer />   
            </Container>
        </IsAdmin>
        
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

const Container = styled.div`
    position: relative;
    min-height: 100vh;
    height: 100%;
    width: 100%;
    
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
    width: 150px;
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
  justify-content: right;

`;

const RejectButton = styled.button`
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
    background-color: #FF9009;
    border: 2px solid #FF9008;
    color: #FFFFFF;
  }
`;

