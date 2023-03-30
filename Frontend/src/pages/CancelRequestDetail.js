import React, {useState} from "react";
import styled from "styled-components"
import Footer from "../components/global/Footer.js";
import { IsAdmin } from "../components/IsAuth.js";
import CancelRequestInfo from "../components/cancelRequestInfo.js";
import {useQuery} from 'react-query';
import { fetchCancellingRequestHandler } from '../handlers/cancelRequestHandler';
import { useParams } from "react-router-dom";

export default function CancelRequestDetailPage(){
    const [data, setData] = useState({});

    const {id} = useParams();



    const { isLoading, error} = useQuery(
        'cancelRequest',() =>
        {
            fetchCancellingRequestHandler(id).then((res) => {
                
                if(res.data.success){
                    if(res.data.data !== null)
                        setData(res.data.data);
                }
            }).catch((err) => {
                console.log(err);
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

    return(
        <IsAdmin>
            <Container>
                <CancelRequestInfo cancelRequestId={data.cancelRequestId} title={data.title} reporter={data.reporter} img={data.img_url} description={data.description} report_date={data.report_date} status={data.status} />
                <Footer />   
            </Container>
        </IsAdmin>
        
    )
}


const Container = styled.div`
    position: relative;
    min-height: 100vh;
    height: 100%;
    width: 100%;
    
`;
