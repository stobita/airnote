import styled from "styled-components";
import { FieldItemBase } from "./FieldItemBase";

export const Input = styled.input`
  ${FieldItemBase}
  color: ${props => props.theme.solid};
  border: 1px solid ${props => props.theme.border};
  background: ${props => props.theme.bg};
`;
