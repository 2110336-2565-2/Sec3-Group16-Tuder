import styled from "styled-components";
import FormT from "../FormStyle";
//icon
import { CameraFilled } from "@ant-design/icons";

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
  width: ${(props) => (props.isSpan ? "100%" : "")};
  gap: 10px;
`;

export const InfoTitle = styled.h2`
  font-size: 16px;
  font-weight: bold;
`;

export const InfoContent = styled.span`
  font-size: 16px;
  font-weight: regular;
`;

// Form
export const FormP = {
  ProfilePictureWrapper: styled.div`
    width: 90px;
    height: 90px;
    border-radius: 50%;
    background-color: #ff7008;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    cursor: pointer;
    margin-top: 20px;
  `,

  ProfilePicture: styled.img`
    border-radius: 50%;
    width: 85%;
    height: 85%;
    object-fit: cover;
  `,

  CameraIconWrapper: styled.div`
    width: 30px;
    height: 30px;
    border-radius: 50%;
    background-color: #ff7008;
    display: flex;
    align-items: center;
    justify-content: center;
    overflow: hidden;
    position: absolute;
    bottom: 0;
    right: 0;
  `,

  CameraIcon: styled(CameraFilled)`
    font-size: 20px;
    color: white;
  `,

  FormContainer: styled.div`
    width: 100%;
    padding: 0px 40px;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    gap: 25px 0px;
    margin-top: 20px;
  `,

  InputComponent: styled(FormT.Component)`
    width: ${(props) => props.width || "48%"};
  `,

  Label: styled(FormT.Label)`
    font-size: 16px;
  `,

  TextInput: styled(FormT.TextInput)`
    width: 100%;
  `,

  TextArea: styled.textarea`
    font-size: 13px;
    padding: 10px 10px;
    border: 1px solid black;
    border-radius: 6px;
    background-color: white;
    color: black;

    // scrollbar
    scrollbar-width: thin;
    ::-webkit-scrollbar {
      width: 8px;
    }

    ::-webkit-scrollbar-track {
      background: transparent;
    }

    ::-webkit-scrollbar-thumb {
      background-color: #d9d9d9;
      border-radius: 20px;
      border: 3px solid transparent;
    }
  `,

  Select: styled(FormT.Select)`
    width: 100%;
  `,

  Option: styled(FormT.Option)`

  `,
  DateInput: styled(FormT.DateInput)`
    width: 100%;
  `,
};
