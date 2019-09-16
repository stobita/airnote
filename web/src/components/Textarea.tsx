import styled from "styled-components";
import { FieldItemBase } from "./FieldItemBase";

export const Textarea = styled.textarea`
  ${FieldItemBase}
  color: ${props => props.theme.solid};
  border: 1px solid ${props => props.theme.border};
  background: ${props => props.theme.bg};
  resize: none;
  box-sizing: border-box;
  height: 64px;
`;
