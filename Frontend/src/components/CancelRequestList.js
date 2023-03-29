import styled from 'styled-components';
import {useQuery} from 'react-query';
import { fetchCancellingClassHandler } from '../handlers/cancellingClassHandler';
import { useDataContext } from '../pages/CancelRequestList';
import CancelRequest from '../components/CancelRequest';
import React,{useNavigate} from 'react';


export default function CancelRequestList(){ 
    const {data, setData} = useDataContext();

    const { isLoading, error} = useQuery(
        'cancellingClass',() =>
        {
            fetchCancellingClassHandler().then((res) => {
                
                if(res.data.success){
                    if(res.data.data !== null)
                        setData(res.data.data);
                }
            }).catch((err) => {
                console.log(err);
                alert("Unauthorized, please login again");
                window.location.href = "/sign-in";
            }
            )
        },
        {
            refetchOnWindowFocus: false,
        }
        );

    if (isLoading) {
        return <div>Loading...</div>
    }

    if (error) {
        return <div>Error: {error.message}</div>
    }
    
    if (data === null) {
        return <div>Error</div>
    }

    if(data === []){
        return <div>Empty</div>
    }

    const removeItem = (classId) => {
        setData(data.filter(item => item.classId !== classId));
    }
    

    return (
        <Container>
            <h1>Cancel-Class Requests</h1>
            <RequestListPage>
                
                    <RequestListcontent>
                        <CancelRequest type="Topic" classId='ClassID' course='Course' tutor='Tutor' student='Student' subject='Subject' total_hour='Total Hours' success_hour='Success Hour'   price='Price/Hour' /> 
                        {data.map(item => (  

                            <CancelRequest key={item.classId} removeItem={removeItem} classId={item.classId} course={item.course} tutor={item.tutor_name} student={item.student_name} subject={item.subject} total_hour={item.total_hours} success_hour={item.success_hours}   price={item.price} /> 
                        ))}
                        
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


