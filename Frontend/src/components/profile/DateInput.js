import React from "react";
import { FormP } from "./ProfileStyle";

export default function DateInput({ label, id, name, value, onChange ,width}) {

  return (
    <FormP.InputComponent width={width}>
      <FormP.Label htmlFor={id} >{label}</FormP.Label>
      <FormP.DateInput
        type="date"
        id={id}
        name={name}
        value={value}
        onChange={onChange}
      ></FormP.DateInput>
    </FormP.InputComponent>
  );
}
