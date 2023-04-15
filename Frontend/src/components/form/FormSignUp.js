import FormT from "./FormStyle.js";
import React, { useState } from "react";
import signUpHandler from "../../handlers/signUpHandler.js";
import signupContent, {
  tutorSpecificContent,
} from "../../datas/SignUp.role.js";
import styled from "styled-components";
import { toast } from "react-hot-toast";

import { useNavigate } from "react-router-dom";

export default function FormSignUp() {
  // Input State
  const [firstname, setFirstName] = useState("");
  const [lastname, setLastName] = useState("");
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmpassword, setConfirmPassword] = useState("");
  const [address, setAddress] = useState("");
  const [contactnumber, setContactNumber] = useState("");
  const [gender, setGender] = useState("");
  const [birthdate, setBirthDate] = useState("");
  const [role, setRole] = useState("");
  const [citizenID, setCitizenID] = useState("");
  const [description, setDescription] = useState("");
  const [cardHolderName, setCardHolderName] = useState("");
  const [cardNumber, setCardNumber] = useState("");
  const [formattedCardNumber, setFormattedCardNumber] = useState(""); // use for display only [1234 5678 9012 3456]
  const [expiryDate, setExpiryDate] = useState("");
  const [cvv, setCVV] = useState("");

  const valueSetter = {
    "First Name": setFirstName,
    "Last Name": setLastName,
    Username: setUsername,
    Email: setEmail,
    Password: setPassword,
    "Confirm Password": setConfirmPassword,
    Address: setAddress,
    "Contact Number": setContactNumber,
    Gender: setGender,
    "Birth Date": setBirthDate,
    Role: setRole,
  };
  const setStatus = useState("waiting")[1];
  const navigate = useNavigate();
  const [isSubmitting, setIsSubmitting] = useState(false);
  async function submitHandler(e) {
    e.preventDefault();
    if (isSubmitting) return;
    setIsSubmitting(true);
    const loadingToast = toast.loading("Signing up...");
    let birthdateISO = new Date(birthdate).toISOString();
    const omiseData =
      role === "tutor"
        ? {
            expiration_month: parseInt(expiryDate.split("/")[0].trim()),
            expiration_year: parseInt(expiryDate.split("/")[1].trim()),
            name: cardHolderName,
            number: cardNumber,
            security_code: cvv,
          }
        : null;
    const studentSignUpData = {
      username: username,
      password: password,
      email: email,
      confirm_password: confirmpassword,
      firstname: firstname,
      lastname: lastname,
      address: address,
      phone: contactnumber,
      birthdate: birthdateISO,
      gender: gender,
      role: role,
      citizen_id: role === "tutor" ? citizenID : "",
      description: role === "tutor" ? description : "",
    };
    setStatus("submitting");
    if (role === "tutor") {
      Omise.setPublicKey(process.env.REACT_APP_OMISE_PUBLIC_KEY);
      Omise.createToken("card", omiseData, (statusCode, response) => {
        if (statusCode === 200) {
          const tutorSignUpData = {
            ...studentSignUpData,
            omise_bank_token: response.id,
          };
          signUpHandler(tutorSignUpData, navigate)
            .then(() => {
              toast.dismiss(loadingToast);
              toast.success("Sign up successfully");
              setStatus("success");
              setIsSubmitting(false);
            })
            .catch((error) => {
              toast.dismiss(loadingToast);
              toast.error(error.message);
              setStatus("error");
              setIsSubmitting(false);
            });
        } else {
          toast.dismiss(loadingToast);
          toast.error(response.message);
          setStatus("error");
          setIsSubmitting(false);
        }
      });
    } else {
      signUpHandler(studentSignUpData, navigate)
        .then(() => {
          toast.dismiss(loadingToast);
          toast.success("Sign up successfully");
          setStatus("success");
          setIsSubmitting(false);
        })
        .catch((error) => {
          toast.dismiss(loadingToast);
          toast.error(error.message);
          setStatus("error");
          setIsSubmitting(false);
        });
    }
  }

  const signupContents = signupContent.contents;
  const tutorSpecificContents = tutorSpecificContent.contents;
  const signupContentElement = signupContents.map((content, index) => {
    return (
      <FormT.Content key={index}>
        {content.map((element, elementindex) => {
          const pholder = "Enter your " + element;
          const name = element.replace(" ", "").toLowerCase();
          let type = "";
          let boxsize = "";
          let value = "";
          let onChange = (e) => valueSetter[element](e.target.value);

          if (element === "Username") {
            type = "text";
            boxsize = "315px";
            value = username;
          } else if (element === "Email") {
            type = "email";
            boxsize = "315px";
            value = email;
          } else if (element === "Password") {
            type = "password";
            boxsize = "150px";
            value = password;
          } else if (element === "Confirm Password") {
            type = "password";
            boxsize = "150px";
            value = confirmpassword;
          } else if (element === "Birth Date") {
            type = "date";
            boxsize = "150px";
            value = birthdate;
          } else if (element === "Contact Number") {
            type = "tel";
            boxsize = "315px";
            value = contactnumber;
          } else if (element === "First Name") {
            type = "text";
            boxsize = "150px";
            value = firstname;
          } else if (element === "Last Name") {
            type = "text";
            boxsize = "150px";
            value = lastname;
          } else if (element === "Gender") {
            type = "text";
            boxsize = "150px";
            value = gender;
          } else if (element === "Address") {
            type = "text";
            boxsize = "315px";
            value = address;
          }

          if (element === "Birth Date") {
            return (
              <FormT.Component key={elementindex}>
                <FormT.Label>{element} :</FormT.Label>
                <FormT.DateInput
                  BoxSize={boxsize}
                  name={name}
                  type={type}
                  placeholder={pholder}
                  value={value}
                  onChange={onChange}
                  required
                />
              </FormT.Component>
            );
          } else if (element === "Gender") {
            return (
              <FormT.Component key={elementindex}>
                <FormT.Label>{element} :</FormT.Label>
                <FormT.Select
                  BoxSize="150px"
                  name={name}
                  value={value}
                  onChange={onChange}
                  required
                >
                  <FormT.Option value="" disabled hidden>
                    Please select
                  </FormT.Option>
                  <FormT.Option value="male">male</FormT.Option>
                  <FormT.Option value="female">female</FormT.Option>
                </FormT.Select>
              </FormT.Component>
            );
          } else {
            return (
              <FormT.Component key={elementindex}>
                <FormT.Label>{element} :</FormT.Label>
                <FormT.TextInput
                  BoxSize={boxsize}
                  name={name}
                  type={type}
                  placeholder={pholder}
                  value={value}
                  onChange={onChange}
                  required
                />
              </FormT.Component>
            );
          }
        })}
      </FormT.Content>
    );
  });
  const tutorSpecificContentElement = tutorSpecificContents.map(
    (content, index) => {
      return (
        <FormT.Content key={"tutor" + index}>
          {content.map((element, elementindex) => {
            const pholder = "Enter your " + element;
            const name = element.replace(" ", "_").toLowerCase();
            if (element === "Citizen ID") {
              return (
                <FormT.Component key={elementindex}>
                  <FormT.Label>{element} :</FormT.Label>
                  <FormT.TextInput
                    BoxSize="315px"
                    name={name}
                    type="text"
                    value={citizenID}
                    placeholder={pholder}
                    onChange={(e) => setCitizenID(e.target.value)}
                    required
                  />
                </FormT.Component>
              );
            } else if (element === "Description") {
              return (
                <FormT.Component key={elementindex}>
                  <FormT.Label>{element} :</FormT.Label>
                  <FormT.TextInput
                    BoxSize="315px"
                    name={name}
                    type="text"
                    value={description}
                    placeholder={pholder}
                    onChange={(e) => setDescription(e.target.value)}
                    required
                  />
                </FormT.Component>
              );
            } else if (element === "Card Holder Name") {
              return (
                <FormT.Component key={elementindex}>
                  <FormT.Label>{element} :</FormT.Label>
                  <FormT.TextInput
                    BoxSize="315px"
                    name={name}
                    type="text"
                    value={cardHolderName}
                    placeholder={pholder}
                    onChange={(e) => setCardHolderName(e.target.value)}
                    required
                  />
                </FormT.Component>
              );
            } else if (element === "Card Number") {
              return (
                <FormT.Component key={elementindex}>
                  <FormT.Label>{element} :</FormT.Label>
                  <FormT.TextInput
                    BoxSize="315px"
                    name={name}
                    type="tel"
                    value={formattedCardNumber}
                    placeholder="1234 5678 9012 3456"
                    inputmode="numeric"
                    pattern="[0-9\s]{13,19}"
                    autocomplete="cc-number"
                    maxLength="19"
                    minLength="19"
                    onChange={(e) => {
                      const formattedValue = e.target.value
                        .replace(/\D/g, "") // Remove non-numeric characters
                        .replace(/(.{4})/g, "$1 ") // Add a space after every 4 digits
                        .trim(); // Remove any leading/trailing spaces
                      setFormattedCardNumber(formattedValue);
                      setCardNumber(formattedValue.replace(/\s/g, ""));
                    }}
                    required
                  />
                </FormT.Component>
              );
            } else if (element === "Expiry Date") {
              return (
                <FormT.Component key={elementindex}>
                  <FormT.Label>{element} :</FormT.Label>
                  <FormT.DateInput
                    BoxSize="150px"
                    name={name}
                    type="text"
                    value={expiryDate}
                    placeholder="mm / yyyy"
                    inputmode="numeric"
                    autocomplete="cc-exp"
                    pattern="^(0[1-9]|1[0-2])\s\/\s[0-9]{4}$"
                    maxLength="9"
                    minLength="9"
                    onChange={(e) => {
                      const formattedValue = e.target.value
                        .replace(/\D/g, "") // Remove non-numeric characters
                        .replace(/^(\d{2})(\d{0,4})$/, "$1 / $2") // Format date as MM / YYYY
                        .trim() // Remove any leading/trailing spaces
                        .slice(0, 9); // Limit to 9 characters
                      setExpiryDate(formattedValue);
                    }}
                    required
                  />
                </FormT.Component>
              );
            } else if (element === "CVV") {
              return (
                <FormT.Component key={elementindex}>
                  <FormT.Label>{element} :</FormT.Label>
                  <FormT.TextInput
                    BoxSize="150px"
                    name={name}
                    type="tel"
                    value={cvv}
                    placeholder={123}
                    inputmode="numeric"
                    pattern="[0-9\s]{3}"
                    autocomplete="cc-csc"
                    minLength="3"
                    maxlength="3"
                    onChange={(e) => {
                      const formattedValue = e.target.value
                        .replace(/\D/g, "") // Remove non-numeric characters
                        .trim(); // Remove any leading/trailing spaces
                      setCVV(formattedValue);
                    }}
                    required
                  />
                </FormT.Component>
              );
            }
          })}
        </FormT.Content>
      );
    }
  );
  return (
    <SignUpForm onSubmit={submitHandler}>
      <FormT.Div FormW="500px">
        <FormT.Header>create account</FormT.Header>
        <FormT.Content>
          Sign in and start managing your candidates!
        </FormT.Content>
        {signupContentElement}
        <FormT.Content>
          <FormT.Component>
            <FormT.Label>As :</FormT.Label>
            <FormT.Select
              BoxSize="315px"
              name="role"
              value={role}
              onChange={(e) => setRole(e.target.value)}
              required
            >
              <FormT.Option value="" disabled hidden>
                Please select
              </FormT.Option>
              <FormT.Option value="student">Student</FormT.Option>
              <FormT.Option value="tutor">Tutor</FormT.Option>
            </FormT.Select>
          </FormT.Component>
        </FormT.Content>
        {role === "tutor" ? tutorSpecificContentElement : null}
        <FormT.Content>
          <FormT.Button type="submit">Sign Up</FormT.Button>
        </FormT.Content>
        <FormT.Content>
          <FormT.ContentSmall>Already have an account ?</FormT.ContentSmall>
          <FormT.Link to="/sign-in" underline="none">
            Sign In
          </FormT.Link>
        </FormT.Content>
      </FormT.Div>
    </SignUpForm>
  );
}

const SignUpForm = styled.form`
  display: flex;
  justify-content: center;
`;
