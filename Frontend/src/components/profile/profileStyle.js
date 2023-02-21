import styled from "styled-components";

export const InfoContainter = styled.div`
    margin: 50px 0px 0px 20px;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    gap: 40px;
`;

export const InfoWrapper = styled.div`
    display: flex;
    flex-direction: column;
    align-items: flex-start;
`;

export const InfoTitle = styled.h2`
    font-size: 20px;
    font-weight: bold;
`;

export const InfoContent = styled.span`
    font-size: 20px;
    font-weight: regular;
`;