import React, { useEffect, useState } from "react";
import styled from "styled-components";
import toast from "react-hot-toast";
import {
  fetchQRCodeForCoursePayment,
  fetchQRCodeForTuitionFee,
  changePaymentStatus,
} from "../../handlers/payment/paymentHandler";
import { paymentData } from "../../datas/Payment";

export default function QRPayment({
  courseID,
  tutorID,
  appointmentID,
  amount,
  callback,
  callbackData,
}) {
  const [qrCode, setQrCode] = useState("");
  const [paymentInfo, setPaymentInfo] = useState(paymentData);
  const [chargeID, setChargeID] = useState("");
  const [paymentID, setPaymentID] = useState("");
  useEffect(() => {
    // POST api/v1/payment/qrCode
    if (courseID) {
      const data = {
        course_id: courseID,
        amount: amount * 100,
      };
      fetchQRCodeForCoursePayment(data)
        .then((res) => {
          setQrCode(res.data.data.qr_code_url);
          setChargeID(res.data.data.payment.charge_id);
          setPaymentID(res.data.data.payment.id);
        })
        .catch((err) => {
          toast.error("Something went wrong");
        });
    } else if (tutorID) {
      const data = {
        tutor_id: tutorID,
        appointment_id: appointmentID,
        amount: amount * 100,
      };
      fetchQRCodeForTuitionFee(data)
        .then((res) => {
          setQrCode(res.data.data.qr_code_url);
          setChargeID(res.data.data.payment.charge_id);
          setPaymentID(res.data.data.payment.id);
        })
        .catch((err) => {
          toast.error("Something went wrong");
        });
    }
    //---------------------------
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
      const data = {
        ...paymentInfo,
        data: {
          ...paymentInfo.data,
          id: chargeID,
          status: "failed",
          source: {
            ...paymentInfo.data.source,
            charge_status: "failed",
          },
        },
      };
      changePaymentStatus(data)
        .then((res) => {
          toast.error("Payment failed");
        })
        .catch((err) => {
          toast.error("Something went wrong");
        });
      // ---------------------------
    } else if (remainingTime === 590) {
      // This is for demo only we will change payment status to successful
      const data = {
        ...paymentInfo,
        data: {
          ...paymentInfo.data,
          id: chargeID,
          status: "successful",
          source: {
            ...paymentInfo.data.source,
            charge_status: "successful",
          },
        },
      };
      changePaymentStatus(data)
        .then((res) => {
          toast.success("Payment successful");
          // if this payment is for enroll
          if(courseID){
            callback({...callbackData,
              payment_id: paymentID,});
          } else if(tutorID){
            callback(callbackData);
          }
        })
        .catch((err) => {
          console.log(err);
          toast.error("Something went wrong");
        });
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
