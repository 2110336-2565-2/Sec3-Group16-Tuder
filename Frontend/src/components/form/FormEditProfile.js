import React, { useState } from "react";
import styled from "styled-components";
import { toast } from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { FormP } from "../profile/ProfileStyle";
import { convertFrontendSchedulesToBackend } from "../../utils/profile/scheduleConverter";
import { studentFields, tutorFields } from "../../datas/Profile.role";
import { updateStudent, updateTutor } from "../../handlers/profile/updateUser";
import FileUploader from "../global/FileUploader";
import TextInput from "../profile/TextInput";
import SelectInput from "../profile/SelectInput";
import DateInput from "../profile/DateInput";
import TimeSelector from "../profile/TimeSelector";

export default function FormEditProfile({ user }) {
  const navigate = useNavigate();
  const fields = user.role === "student" ? studentFields : tutorFields;
  const [isFileUploaderOpen, setIsFileUploaderOpen] = useState(false);
  const [formData, setFormData] = useState({ ...user });
  const [isSubmitting, setIsSubmitting] = useState(false);

  const submitHandler = (e) => {
    e.preventDefault();
    if (isSubmitting) return;
    setIsSubmitting(true);
    const loadingToast = toast.loading("Submitting...");
    let compatibleFormData = { ...formData };
    // Convert some fields format to match the backend
    if (compatibleFormData.new_profile_picture) {
      compatibleFormData.new_profile_picture = compatibleFormData.new_profile_picture.split(',')[1];
    }
    try {
      if (user.role === "student") {
        updateStudent(compatibleFormData).then((res) => {
          navigate("/profile");
          toast.success("Profile updated successfully");
          toast.dismiss(loadingToast);
          setIsSubmitting(false);
        });
      } else if (user.role === "tutor") {
        updateTutor(compatibleFormData).then((res) => {
          navigate("/profile");
          toast.success("Profile updated successfully");
          toast.dismiss(loadingToast);
          setIsSubmitting(false);
        });
      }
    } catch (error) {
        console.log(error);
        toast.error("Profile update failed");
        toast.dismiss(loadingToast);
        setIsSubmitting(false);
    }
  };

  const handleChange = (e) => {
    let value = e.target.value;
    // Convert some fields format to match the backend
    if (e.target.name === "birthdate") {
      value = new Date(e.target.value).toISOString();
    } else if (e.target.name === "schedule") {
      value = convertFrontendSchedulesToBackend(e.target.value);
    }
    
    setFormData({
      ...formData,
      [e.target.name]: value,
    });
  };

  return (
    <Form onSubmit={submitHandler}>
      <FileUploader
        isOpen={isFileUploaderOpen}
        setIsOpen={setIsFileUploaderOpen}
        handleChange={handleChange}
      />
      <Title>Edit Profile</Title>
      <FormP.ProfilePictureWrapper onClick={() => setIsFileUploaderOpen(true)}>
        <FormP.ProfilePicture
          src={
            formData.new_profile_picture
              ? formData.new_profile_picture
              : formData.profile_picture_URL
          }
        />
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
          } else if (field.type === "email") {
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
            const displayedDate = formData[field.name].split("T")[0];
            return (
              <DateInput
                key={field.id}
                label={field.label}
                id={field.id}
                name={field.name}
                value={displayedDate}
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
                value={formData["raw_schedule"]}
                onChange={handleChange}
                width={field.width}
              />
            );
          }
          return null;
        })}
      </FormP.FormContainer>
      <ButtonSection>
        <Button type="cancel" onClick={() => {
          toast.dismiss()
          navigate("/profile")
          }}>
          Cancel
        </Button>
        <Button type="submit">Save</Button>
      </ButtonSection>
    </Form>
  );
}

const Form = styled.form`
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
  background-color: ${(props) =>
    props.type === "cancel" ? "#FFFFFF" : "#FF7008"};
  color: ${(props) => (props.type === "cancel" ? "#FF7008" : "#FFFFFF")};
  font-size: 13px;
  width: 100px;
  padding: 10px;
  border-radius: 10px;
  border: 2px solid #ff7008;

  cursor: pointer;
`;
