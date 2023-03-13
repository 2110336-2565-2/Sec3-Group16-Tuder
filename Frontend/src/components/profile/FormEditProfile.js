import React, { useState } from "react";
import styled from "styled-components";
import { useNavigate } from "react-router-dom";
import { FormP } from "./ProfileStyle";
import { studentFields, tutorFields } from "../../datas/Profile.role";
import FileUploader from "../global/FileUploader";
import TextInput from "./TextInput";
import SelectInput from "./SelectInput";
import DateInput from "./DateInput";
import TimeSelector from "./TimeSelector";

export default function FormEditProfile({ user }) {
  const navigate = useNavigate();
  const fields = user.role === "student" ? studentFields : tutorFields;
  const [isFileUploaderOpen, setIsFileUploaderOpen] = useState(false);
  const [formData, setFormData] = useState({...user, new_profile_picture: user.profile_picture_URL});

  // EDIT PROFILE HANDLER CHANGE THIS TO SEND DATA TO BACKEND
  const editProfileHandler = () => {
    console.log("edit profile");
  };

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    })
    // console.log("e.target.name: ", e.target.name);
    // console.log("e.target.value: ", e.target.value);
    console.log("formData: ", formData)
  };

  return (
    <Container>
      <FileUploader
        isOpen={isFileUploaderOpen}
        setIsOpen={setIsFileUploaderOpen}
        handleChange={handleChange}
      />
      <Title>Edit Profile</Title>
      <FormP.ProfilePictureWrapper onClick={()=>setIsFileUploaderOpen(true)}>
        <FormP.ProfilePicture src={formData.new_profile_picture} />
        <FormP.CameraIconWrapper>
          <FormP.CameraIcon />
        </FormP.CameraIconWrapper>
      </FormP.ProfilePictureWrapper>
      <FormP.FormContainer>
        {fields.map((field) => {
          if (field.type === "text" || field.type === "textArea") {
            return (
              <TextInput
                key={field.id}
                type={field.type}
                label={field.label}
                id={field.id}
                name={field.name}
                value={formData[field.name]}
                onChange={handleChange}
                width={field.width}
              />
            );
          } else if (field.type === "select") {
            return (
              <SelectInput
                key={field.id}
                label={field.label}
                id={field.id}
                isRequired={field.isRequired}
                name={field.name}
                value={formData[field.name]}
                onChange={handleChange}
                options={field.options}
                width={field.width}
              />
            );
          } else if (field.type === "date") {
            return (
              <DateInput
                key={field.id}
                label={field.label}
                id={field.id}
                name={field.name}
                value={formData[field.name]}
                onChange={handleChange}
                width={field.width}
              />
            );
          } else if (field.type === "time") {
            return (
              <TimeSelector
                key={field.id}
                label={field.label}
                id={field.id}
                name={field.name}
                value={formData[field.name]}
                onChange={handleChange}
                width={field.width}
              />
            );
          }
          return null;
        })}
      </FormP.FormContainer>
      <ButtonSection>
        <Button type="cancel" onClick={() => navigate("/profile")}>Cancel</Button>
        <Button type="save" onClick={editProfileHandler}>Save</Button>
      </ButtonSection>
    </Container>
  );
}

const Container = styled.div`
  width: 50%;
  background-color: white;
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 10px;
  padding: 30px 60px;
`;

const Title = styled.h1``;

const ButtonSection = styled.div`
  display: flex;
  justify-content: flex-end;
  align-items: center;
  width: 100%;
  margin-top: 35px;
  gap: 25px;
`;

const Button = styled.button`
  background-color: ${(props) => (props.type === "cancel" ? "#FFFFFF" : "#FF7008")};
  color: ${(props) => (props.type === "cancel" ? "#FF7008" : "#FFFFFF")};
  font-size: 13px;
  width: 100px;
  padding: 10px;
  border-radius: 10px;
  border: 2px solid #ff7008;

  cursor: pointer;
`;