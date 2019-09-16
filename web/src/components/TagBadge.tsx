import styled from "styled-components";

export const TagBadge = styled.span`
  color: ${props => props.theme.primary};
  border: 1px solid ${props => props.theme.primary};
  margin-right: 4px;
  padding: 4px;
  border-radius: 4px;
`;
