


import React, { useState} from 'react';
import styled from 'styled-components';
import { getRole, getUserId } from '../../utils/jwtGet';
import { useNavigate, useParams } from 'react-router-dom';
import { toast } from 'react-hot-toast';
import { submitCancelRequestHandler } from '../../handlers/cancelRequestHandler';


export default function UserCancelRequest() {
    const navigate = useNavigate();

    const [formData, setFormData] = useState({
      title: "",
      description: "",
      img_url: "",
    });

    const matchId = useParams().id;

    const handleChange = (event) => {
      setFormData({
        ...formData,
        [event.target.name]: event.target.value,
      });
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        const data = {
            matchId: matchId,
            userId: getUserId(),
            title: formData.title,
            description: formData.description,
            img_url: formData.img_url,
            reporter_role: getRole(),
        };
        console.log(data)
        submitCancelRequestHandler(data).then((res) => {
            if(res.data.success){
                toast.success('Success');
                navigate("/userrequest");
            }
        }).catch((err) => {
            console.log(err);
            toast.error(err.response.data.message);
        }
        )

    };


    var title = formData.title;
    var description = formData.description;
    var img_url = formData.img_url;


    return (

        <CancelRequest>
            <Title>
                Match Cancellation
            </Title>
            <SubTitle>
              course name
            </SubTitle>
            <SubTitle>
              ufirstname ulastname
            </SubTitle>

            <TopicSection>
              <Topic>
                Title
              </Topic>
              <RequestTitle type="text" name="title" value={title} onChange={handleChange} />
            </TopicSection>
            <TopicSection>
              <Topic>
                Picture
              </Topic>
              <RequestPicture name="img_url" src="/images/Union.png" onChange={handleChange}/>

            </TopicSection>

            <DescriptionSection>
              <Topic>
                Description
              </Topic>
              <Description name="description"  onChange={handleChange}/>
            </DescriptionSection>
            <ButtonSection>
              <CancelButton onClick={() => navigate("/userrequest")}> 
                Cancel
              </CancelButton>
              <SubmitButton onClick={handleSubmit}>
                Submit
              </SubmitButton>
            </ButtonSection>
        </CancelRequest>

    )

}






const CancelRequest = styled.div`

    box-sizing: border-box;
    width: 1153px;
    height: 1031px;
    margin-top: 50px;
    margin-bottom: 50px;
    padding: 50px;
    justify-content: center;

    background: #FFFFFF;
    border: 1px solid #858585;
    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
    border-radius: 10px;
`;

const Title = styled.div`
    width: 100%;
    text-align: center;
  
    font-family: 'Lexend Deca';
    font-style: normal;
    font-weight: 700;
    font-size: 35px;
    line-height: 44px;
`;

const SubTitle = styled.div`
    width: 100%;
    text-align: center;

    font-family: 'Lexend Deca';
    font-style: normal;
    font-weight: 300;
    font-size: 16px;
    line-height: 20px;
`;

const TopicSection = styled.div`
margin-top: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 400px;
  height: 76px;
  left: 220px;
  top: 311px;
`;

const Topic = styled.div`
  font-family: 'Lexend Deca';
  font-style: normal;
  font-weight: 500;
  font-size: 16px;
  line-height: 20px;

`;

const RequestTitle = styled.input`
  border: 2px solid #858585;
  border-radius: 5px;
  width: 400px;
  height: 76px;
  left: 220px;
  top: 311px;
  font-size: 20px;
  font-family: 'Lexend Deca';
  font-style: normal;
  font-weight: 500;
`;

const RequestPicture = styled.img`
  width: 40px;
  height: 40px;
`;

const DescriptionSection = styled.div`
  margin-top: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 400px;
`;

const Description = styled.textarea`
  border: 2px solid #858585;
  border-radius: 5px;
  width: 1000px;
  height: 488px;  
  font-family: 'Lexend Deca';
  font-style: normal;
  font-weight: 500;
  font-size: 20px;
  
`;

const ButtonSection = styled.div`
  display: flex;
  flex-direction: row;
  gap: 20px;
  width: 100%;
  margin-top: 20px;

`;

const CancelButton = styled.button`
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

const SubmitButton = styled.button`
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

