import React from "react";
import styled from "styled-components";
import Modal from "react-modal";
import QRPayment from "./QRPayment";
import { CloseOutlined } from "@ant-design/icons";
import { confirm } from "./customToast";

export default function PaymentModal({
  isOpen,
  setIsOpen,
  title,
  courseID,
  tutorID,
  amount,
}) {
  function CloseHandler() {
    confirm("Are you sure you want to discard this payment?", () =>
      setIsOpen(false)
    );
  }
  return (
    <Modal isOpen={isOpen} onRequestClose={CloseHandler} style={modalStyle}>
      <Wrapper>
        <Top>
          <p></p>
          <Title>{title}</Title>
          <CloseIcon onClick={CloseHandler} />
        </Top>
        <QRPayment courseID={courseID} tutorID={tutorID} amount={amount} />
      </Wrapper>
    </Modal>
  );
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

const Top = styled.div`
  width: 100%;
  height: 70px;
  display: flex;
  justify-content: space-between;
  padding: 20px;
`;

const CloseIcon = styled(CloseOutlined)`
  font-size: 20px;
  color: #e2e6ea;
  align-self: flex-start;
  cursor: pointer;
`;

const Title = styled.h1`
  font-size: 30px;
  font-weight: bold;
  color: #113566;
`;
