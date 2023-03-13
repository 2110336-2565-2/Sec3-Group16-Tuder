import React from "react";
import { FormP } from "./ProfileStyle";

export default function TextInput({
  label,
  id,
  isRequired,
  name,
  value,
  options,
  onChange,
  width,
}) {
  return (
    <FormP.InputComponent width={width}>
      <FormP.Label htmlFor={id}>{label}</FormP.Label>
      <FormP.Select
        id={id}
        name={name}
        value={value}
        required={isRequired}
        onChange={onChange}
      >
        {options.map((option) => (
          <option
            value={option.value}
            key={option.label}
            disabled={option.disabled}
            hidden={option.hidden}
          >
            {option.label}
          </option>
        ))}
      </FormP.Select>
    </FormP.InputComponent>
  );
}
