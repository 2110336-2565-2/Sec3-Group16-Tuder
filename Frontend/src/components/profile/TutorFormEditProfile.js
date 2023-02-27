import React, { useState } from "react";
import styled from "styled-components";
import { FormP } from "./ProfileStyle";
import { tutorFields } from "../../datas/Profile.role";
import TextInput from "./TextInput";
import SelectInput from "./SelectInput";
import DateInput from "./DateInput";
import TimeSelector from "./TimeSelector";

export default function TutorFormEditProfile({ user }) {
  const [formData, setFormData] = useState(user);

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  return (
    <Container>
      <Title>Edit Profile</Title>
      <FormP.ProfilePictureWrapper>
        <FormP.ProfilePicture src={formData.profile_picture_URL} />
        <FormP.CameraIconWrapper>
          <FormP.CameraIcon />
        </FormP.CameraIconWrapper>
      </FormP.ProfilePictureWrapper>
      <FormP.FormContainer>
        {tutorFields.map((field) => {
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
  padding: 30px;
`;

const Title = styled.h1``;
