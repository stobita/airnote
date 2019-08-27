import styled from "styled-components";
import colors from "../colors";
import { FieldItemBase } from "./FieldItemBase";

export const Textarea = styled.textarea`
  ${FieldItemBase}
  background: ${colors.mainWhite};
  resize: none;
  box-sizing: border-box;
  height: 64px;
`;
