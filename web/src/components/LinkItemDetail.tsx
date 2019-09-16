import React from "react";
import { Link } from "../model/link";
import styled from "styled-components";

interface Props {
  item: Link;
}
export const LinkItemDetail = (props: Props) => {
  const { url, description } = props.item;
  const displayUrl = url.length > 100 ? `${url.slice(0, 100)}...` : url;

  return (
    <Wrapper>
      <Anchor href={props.item.url} target="_blank" rel="noopener noreferrer">
        {displayUrl}
      </Anchor>
      <Description>{description}</Description>
    </Wrapper>
  );
};

const Wrapper = styled.div``;

const Anchor = styled.a`
  margin-bottom: 8px;
  color: ${props => props.theme.primary};
`;
const Description = styled.p`
  word-break: break-word;
  margin: 0;
  margin-bottom: 8px;
`;
