import React, { useState, createContext, useContext} from "react";
import styled from "styled-components"
import WaveFooter from "../components/global/WaveFooter.js";

const DataContext = createContext({
    data : {
        data: []
    },
    setData : () => {}
});

export const useDataContext = () => useContext(DataContext);
export default function CancelRequestList(){
    const [data, setData] = useState({data:[]});

    return(
        <Container>
            <DataContext.Provider value={{data, setData}}>
            <ContainerWithHeight margintop='100px'>
                


            </ContainerWithHeight>
            <WaveFooterWrapper>
                <WaveFooter />
            </WaveFooterWrapper>
            </DataContext.Provider>
        </Container>
        
    )
}

const ContainerWithHeight = styled.div`
    display: flex;
    flex-direction: column
    padding: 0px 30px;
    margin-top: ${(props) => {
        return props.margintop
    }};
    justify-content: center;
`;

const WaveFooterWrapper = styled.div`
    position: absolute;
    bottom: 0;
    width: 100%;
`;

const Container = styled.div`
    position: relative;
    min-height: 100vh;
    width: 100%;
`;
