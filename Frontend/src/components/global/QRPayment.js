import React, { useEffect, useState } from "react";
import styled from "styled-components";

export default function QRPayment({ courseID, tutorID, amount }) {
  const [qrCode, setQrCode] = useState("");
  useEffect(() => {
    // POST api/v1/payment/qrCode

    //---------------------------
    setQrCode(
      "https://api.omise.co/charges/chrg_test_5vg0c1ujdv90959eu0z/documents/docu_test_5vg0c1wyz7wk5rcs8he/downloads/B1C7932DBA80CC4E"
    );
  }, [courseID]);

  const [remainingTime, setRemainingTime] = useState(600); // 10 minutes in seconds
  useEffect(() => {
    const timer = setInterval(() => {
      setRemainingTime((prevTime) => prevTime - 1);
    }, 1000);

    return () => clearInterval(timer);
  }, []);

  useEffect(() => {
    if (remainingTime === 0) {
      // Call API to cancel payment
      // ---------------------------
    } else if (remainingTime === 580) {
      // This is for demo only we will change payment status to completed
      // ---------------------------
    }
  }, [remainingTime]);

  return (
    <Container>
      {remainingTime > 0 ? (
        <>
          <Image src={qrCode} />
          <Countdown>
            Please pay with in {Math.floor(remainingTime / 60)} minutes{" "}
            {remainingTime % 60} seconds
          </Countdown>
        </>
      ) : (
        <Countdown>
          Time is up, please contact us for further assistance
        </Countdown>
      )}
      <ContactInfo>
        <Bold>For any payment-related issues</Bold>
        <br />
        <Bold>Tel :</Bold> 02-123-4567 ( Everyday 8:00 - 20:00 )
      </ContactInfo>
      {remainingTime > 0 ? (
        <TotalPriceWrapper>
          <TotalPrice>Scan QR To Pay {amount} Baht</TotalPrice>
        </TotalPriceWrapper>
      ) : null}
    </Container>
  );
}

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  gap: 20px;
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
  font-weight: 500;
  color: white;
  padding: 30px 20px;
  text-align: center;
`;

const Countdown = styled.p`
  font-size: 20px;
  font-weight: 500;
  color: #ff0000;
`;

const ContactInfo = styled.p`
  font-size: 20px;
`;

const Bold = styled.span`
  font-weight: 500;
`;
