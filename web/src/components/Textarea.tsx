import styled from "styled-components";
import colors from "../colors";
import { FieldItemBase } from "./FieldItemBase";

export const Textarea = styled.textarea`
  ${FieldItemBase}
  border: 1px solid ${props => props.theme.border};
  background: ${colors.white};
  resize: none;
  box-sizing: border-box;
  height: 64px;
`;
