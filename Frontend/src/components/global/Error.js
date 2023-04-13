import React from 'react'
import styled from 'styled-components'
import { FrownOutlined } from '@ant-design/icons'

export default function Error({error_msg}) {
  return (
    <Container>
        <FrownOutlined style={iconStyle} />
        <ErrorText>{error_msg}</ErrorText>
    </Container>
  )
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    min-height: 70vh;
    gap: 20px;
`;

const iconStyle = {
    fontSize: '80px',
    color: '#EB7B42'
}

const ErrorText = styled.p`
    font-size: 30px;
`;
