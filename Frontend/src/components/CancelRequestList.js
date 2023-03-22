import styled from 'styled-components';
import Course from '../components/Course';
import {useQuery} from 'react-query';
import { fetchCourseHandler } from '../handlers/searchCourseHandler';
import { useDataContext } from '../pages/Courses';
import CancelRequest from '../components/CancelRequest';
import React from 'react';


export default function CancelRequestList(){ 
    // const {data, setData} = useDataContext();

    // const { isLoading, error} = useQuery(
    //     'courses',() =>
    //     {
    //         fetchCourseHandler().then((res) => {
                
    //             if(res.data.success){
                    
    //                 setData({data:res.data.result});
    //             }
    //         }).catch((err) => {
    //             console.log(err);
    //         }
    //         )
    //     },
    //     {
    //         refetchOnWindowFocus: false,
    //     }
    //     );


    // if (isLoading) {
    //     return <div>Loading...</div>
    // }

    // if (error) {
    //     return <div>Error: {error.message}</div>
    // }
    
    // if (data === null) {
    //     return <div>Error</div>
    // }

    // if(data.data === []){
    //     return <div>Empty</div>
    // }
    

    return (
        <Container>
            <h1>Cancel-Class Requests</h1>
            <RequestListPage>
                
                    <RequestListcontent>
                        <FirstRow>
                            <FirstRowInfo>Class</FirstRowInfo>
                            <FirstRowInfo>Course</FirstRowInfo>
                            <FirstRowInfo>Tutor</FirstRowInfo>
                            <FirstRowInfo>Subject</FirstRowInfo>
                            <FirstRowInfo>Time</FirstRowInfo>
                            <FirstRowInfo>Price</FirstRowInfo>
                        </FirstRow>
                        <CancelRequest  classId='50323' course='Algorithm for daily life' tutor='John' subject='Math' time='2' price='20' img='https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.freepik.com%2Fpremium-vector%2Fmathematics-education-concept_1079551.htm&psig=AOvVaw3QZ-0-2Q8' /> 
                        
                    </RequestListcontent>

                
            </RequestListPage>
        </Container>
    )
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    gap: 20px;
    justify-content: center;
    align-items: center;
    border: 1px solid black;
    padding: 20px;
    margin-top: 20px;
    margin-bottom: 20px;
`

const RequestListPage = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: center;
    padding: 20px;
    
`

const RequestListcontent = styled.div`
    display: flex;
    flex-direction: column;
    gap: 20px;
    
`

const FirstRow = styled.div`
    display: flex;
    flex-direction: row;
    gap: 20px;
    justify-content: left;
    margin-left: 10px;
    
`

const FirstRowInfo = styled.div`
    display: flex;
    flex-direction: column;
    gap: 10px;
    font-size: 20px;
    margin-top: 15px;
    margin-right: 10px;
`


