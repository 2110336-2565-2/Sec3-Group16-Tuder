import React, { useState } from "react";
import styled from "styled-components";
import Modal from "react-modal";
import { CloseOutlined } from "@ant-design/icons";

export default function FileUploader({ isOpen, setIsOpen, handleChange }) {
  const [file, setFile] = useState(null);
  const handleFileChange = (e) => {
    const addedFile = e.target.files[0];
    setFile(addedFile);
    const reader = new FileReader();
    reader.readAsDataURL(addedFile);
    reader.onloadend = () => {
      handleChange({
        target: {
          name: "newProfilePicture",
          value: reader.result,
        },
      });
    };
    setIsOpen(false);
  };

  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={() => setIsOpen(false)}
      style={modalStyle}
    >
      <Container>
        <Top>
          <Title>Upload Picture</Title>
          <CloseIcon onClick={() => setIsOpen(false)} />
        </Top>
        <InputContainer>
          <Input
            type="file"
            id="file"
            name="file"
            accept="image/png, image/jpeg"
            onChange={handleFileChange}
          />
          <InputLabel htmlFor="file">Click to browse your files</InputLabel>
        </InputContainer>
      </Container>
    </Modal>
  );
}

const modalStyle = {
  overlay: {
    backgroundColor: "rgba(0, 0, 0, 0.5)",
    zIndex: 1000,
  },
  content: {
    width: "380px",
    height: "400px",
    margin: "auto",
    padding: "0",
    border: "none",
    borderRadius: "10px",
  },
};

const Container = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
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

const Title = styled.h2`
  font-size: 16px;
  font-weight: 400px;
  align-self: center;
  margin-left: 35%;
`;

const InputContainer = styled.div`
  position: relative;
  width: 100%;
  height: 330px;
  background-color: #F7F9FB;
`;

const InputLabel = styled.label`
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  font-weight: regular;
  text-align: center;
  color: #444;
  bottom: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 1;
  cursor: pointer;
`;

const Input = styled.input`
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  z-index: 2;
`;