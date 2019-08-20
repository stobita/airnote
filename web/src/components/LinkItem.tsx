import React, { useCallback } from "react";
import { Link } from "../model/link";
import { ReactComponent as DefaultImage } from "../assets/default.svg";
import styled from "styled-components";
import colors from "../colors";

interface Props {
  item: Link;
  onClick: (l: Link) => void;
}

export const LinkItem = (props: Props) => {
  const handleOnClick = useCallback(() => {
    props.onClick(props.item);
  }, [props]);

  return (
    <Wrapper onClick={handleOnClick}>
      <DefaultImage />
      <Bottom>
        <p>{props.item.description}</p>
      </Bottom>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  flex: 0 1 25%;
  border: 1px solid ${colors.borderGray};
  border-radius: 4px;
  box-sizing: border-box;
  margin: 8px;
  cursor: pointer;
`;

const Bottom = styled.div`
  padding: 8px;
`;
