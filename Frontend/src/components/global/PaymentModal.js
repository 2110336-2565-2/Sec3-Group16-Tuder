import React from 'react'
import styled from 'styled-components'
import Modal from "react-modal";
import QRPayment from './QRPayment';

export default function PaymentModal({ isOpen, setIsOpen, title, courseID, tutorID, amount }) {
  return (
    <Modal
        isOpen={isOpen}
        onRequestClose={() => setIsOpen(false)}
        style={modalStyle}
        >
        <Wrapper>
            <Title>{title}</Title>
            <QRPayment courseID={courseID} tutorID={tutorID} amount={amount} />
        </Wrapper>
    </Modal>
  )
}

const modalStyle = {
    overlay: {
      backgroundColor: "rgba(0, 0, 0, 0.5)",
      zIndex: 1000,
    },
    content: {
      margin: "auto",
      padding: "20px 5px",
      border: "none",
      borderRadius: "10px",
    },
  };

const Wrapper = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;
`;

const Title = styled.h1`
    font-size: 30px;
    font-weight: bold;
    color: #113566;
`;