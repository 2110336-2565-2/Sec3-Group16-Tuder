import React, { useState } from 'react';
import { toast } from 'react-hot-toast';
import Footer from "../components/global/Footer.js";
import submitIssueReportHandler from '../handlers/submitIssueReportHandler';
import styled from 'styled-components';
import { Outlet, Link } from 'react-router-dom';

export default function Report(){

    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');
    const [contact, setContact] = useState('');

    async function submitHandler(e){
        e.preventDefault();
        let date = Date.now()
        const reportData = {
            title,
            description,
            contact,
            date
        }
        try{
            await submitIssueReportHandler(reportData)
            toast.success('Thank you for your attention.')
        } catch (error){
            toast.error(error.message)
        }
    }

    return (
        <>
            <Container>
                <ProblemContainer>
                    <form onSubmit={submitHandler} >
                        <ReportForm>
                            <HeaderReport>Problem report</HeaderReport>
                            <HeaderContent>If you have any problem or question, let me know here</HeaderContent>
                            <LabelContent>Title :</LabelContent>
                            <InputContent>
                                <InputTopic BoxSize='300px' name='title' type='text' placeholder='' value={title} onChange={(e) => setTitle(e.target.value)}/>
                            </InputContent>
                            <LabelContent>Description :</LabelContent>
                            <InputContent>
                                <InputAreaTopic  BoxSize='600px' name='description' type='description' placeholder='' value={description} onChange={(e) => setDescription(e.target.value)}/>
                            </InputContent>
                            <LabelContent>Contact :</LabelContent>
                            <InputContent>
                                <InputTopic  BoxSize='300px' name='contact' type='contact' placeholder='Email or Phonenumber' value={contact} onChange={(e) => setContact(e.target.value)}/>
                            </InputContent>
                            <ButtonContent>
                                <ButtonFlex>
                                    <ButtonSubmit type='submit'>Submit</ButtonSubmit>
                                </ButtonFlex>
                            </ButtonContent>
                        </ReportForm>
                    </form>
                </ProblemContainer>
            </Container>
            <Footer />
            <Outlet/>
        </>
        
    )
}

const Container = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  background-color: #fdedeb;
`;


const ProblemContainer = styled.div`
    width: 75%;
    min-width: 650px;
    border: 0.5px solid black;
    border-radius: 10px;
    background-color: white;
    margin-top: 3rem;
    margin-bottom: 3rem;
`
const ReportForm = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap:10px;
    padding-top:20px;
    padding-bottom:20px;
`

const HeaderReport = styled.div`
    display: flex;
    font-size: 30px;
    font-weight: 600;
`

const HeaderContent = styled.div`
    display: flex;
`
const InputContent = styled.div`
    display: flex;
    width: 600px;
`

const ButtonContent = styled.div`
    display: flex;
    margin-top: 10px;
    width: 600px;
    flex-direction: row;
    justify-content: flex-end;
`

const ButtonFlex = styled.div`
    display: flex;
`

const ButtonSubmit = styled.button`
  width: 80px;
  height: 35px;
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

const InputTopic = styled.input`
    width: ${(props) => {
        return props.BoxSize
    }};
    padding: 10px 13px;
    border: 1px solid black;
    border-radius: 6px;
    background-color: white;
    color: black;
    &::placeholder{
        color:black;
    }
`

const LabelContent = styled.div`
    margin-top: 10px;
    display: flex;
    width: 600px;
`

const InputAreaTopic = styled.textarea`
    width: ${(props) => {
        return props.BoxSize
    }};
    height: 200px;
    padding: 10px 13px;
    border: 1px solid black;
    border-radius: 6px;
    background-color: white;
    color: black;
    &::placeholder{
        color:black;
    }
`