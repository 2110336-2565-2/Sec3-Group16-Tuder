import React from "react";
import styled from "styled-components";

export default function QRPayment({total_price}) {
  return (
    <Container>
      <Image src="https://api.omise.co/charges/chrg_test_5vg0c1ujdv90959eu0z/documents/docu_test_5vg0c1wyz7wk5rcs8he/downloads/B1C7932DBA80CC4E" />
      <TotalPriceWrapper>
        <TotalPrice>Scan QR To Pay {total_price} Baht</TotalPrice>
      </TotalPriceWrapper>
    </Container>
  );
}

const Container = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
`;


const Image = styled.img`
    width: 400px;
`;

const TotalPriceWrapper = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #113566;
    width: 400px;
`;

const TotalPrice = styled.p`
    font-size: 20px;
    font-weight: bold;
    color: white;
    padding:30px 20px;
    text-align: center;
`;
