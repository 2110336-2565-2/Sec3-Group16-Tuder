import React from "react";
import { FormP } from "./ProfileStyle";

export default function TextInput({ type, label, id, name, value, onChange ,width}) {
  if (type === "text") {
  return (
    <FormP.InputComponent width={width}>
      <FormP.Label htmlFor={id} >{label}</FormP.Label>
      <FormP.TextInput
        type="text"
        id={id}
        name={name}
        value={value}
        onChange={onChange}
      ></FormP.TextInput>
    </FormP.InputComponent>
  )} else if (type === "textArea") {
    return (
      <FormP.InputComponent width={width}>
        <FormP.Label htmlFor={id} >{label}</FormP.Label>
        <FormP.TextArea
          type="text"
  
          id={id}
          name={name}
          value={value}
          onChange={onChange}
        ></FormP.TextArea>
      </FormP.InputComponent>
    )
  }
}
