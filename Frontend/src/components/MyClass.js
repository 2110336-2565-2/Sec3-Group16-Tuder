import React from "react";
import styled from "styled-components";

export default function MyClass() {

    return (
    <Request>
        <ClassSection>
          <GridSection>
            <ClassImg src="" alt="classImg" />
          </GridSection>
  
          <ClassInfoSection>
            <ClassFlex>
              <ClassInfo>
                <h3>Class Name</h3>
              </ClassInfo>
              <ClassInfo>
                <InfoTitle min_w="43px">Tutor :</InfoTitle>
                <InfoDesc>veveve</InfoDesc>
              </ClassInfo>
  
              <ClassInfo>
                <InfoTitle min_w="79px">Upcoming Class :</InfoTitle>
                <InfoDesc>vevev</InfoDesc>
              </ClassInfo>
              <ClassInfo>
                <InfoTitle min_w="105px">Remaining :</InfoTitle>
                <InfoDesc>
                  vwev
                </InfoDesc>
              </ClassInfo>
              <ClassInfoButton>
                <InfoTitle min_w="60px">Status :</InfoTitle>
                <InfoDesc>
                  ongoing
                </InfoDesc>
              </ClassInfoButton>
            </ClassFlex>
          </ClassInfoSection>
        </ClassSection>
        <GridSection>
            {/* put button in here */}
        </GridSection>
    </Request>
    );
}

const Request = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: center;
  background-color: white;
  border: 1px solid #000000;
  min-height: 167px;
  min-width: 1000px;
  gap: 20px;
  padding: 10px;
  border-radius: 10px;
`

const ClassSection = styled.div`
  display: grid;
  grid-template-columns: 30% 40% 30%;
  width: 100%;
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
  font-size: 14px;
  font-weight: 450;
  min-width: ${(props) => {
    return props.min_w;
  }};
`;

const InfoDesc = styled.div`
  display: flex;
  font-size: 14px;
  font-weight: 300;
`;
