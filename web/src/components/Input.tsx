import styled from "styled-components";
import colors from "../colors";
import { FieldItemBase } from "./FieldItemBase";

export const Input = styled.input`
  ${FieldItemBase}
  border: 1px solid ${props => props.theme.border};
  background: ${colors.mainWhite};
`;
