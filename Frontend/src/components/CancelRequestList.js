import styled from 'styled-components';
import {useQuery} from 'react-query';
import { fetchCancellingClassHandler } from '../handlers/cancellingClassHandler';
import { useDataContext } from '../pages/CancelRequestList';
import CancelRequest from '../components/CancelRequest';
import React from 'react';
import {useNavigate} from 'react-router-dom';


export default function CancelRequestList(){ 
    const {data, setData} = useDataContext();
    const navigate = useNavigate();

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
            <h1>Class Cancellation Requests</h1>
            <RequestListPage>
                
                    <RequestListcontent>
                        {data.map(item => (
                            <div onClick={(e) =>navigate('/cancel-request-detail')}>

                            <CancelRequest  title="big hee" img="/images/index.png" course="hee" reporter="Name of hee" report_date="Today" status="ongoing"/>
                            </div>
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


