import React, { useState } from "react";
import styled, { css } from "styled-components";
import colors from "../colors";

interface Props {
  items: string[];
  inputValue: string;
  onClickItem: (id: number) => void;
  onMouseEnterItem: (id: number) => void;
  hoverIndex: number;
}

type Item = {
  id: number;
  label: string;
};

export const AutoSuggest = (props: Props) => {
  if (props.items.length < 1) return null;

  const handleOnClick = (e: React.MouseEvent<HTMLElement>) => {
    const targetIdx = e.currentTarget.dataset.idx;
    if (targetIdx) {
      props.onClickItem(Number(targetIdx));
    }
  };

  const handleOnMouseEnter = (e: React.MouseEvent<HTMLLIElement>) => {
    const targetIdx = e.currentTarget.dataset.idx;
    if (targetIdx) {
      props.onMouseEnterItem(Number(targetIdx));
    }
  };

  return (
    <Wrapper>
      <ul>
        {props.items.map((v, idx) => (
          <Item
            isSelect={idx === props.hoverIndex}
            key={idx}
            onClick={handleOnClick}
            data-idx={idx}
            onMouseEnter={handleOnMouseEnter}
          >
            {v}
          </Item>
        ))}
      </ul>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background: ${colors.mainWhite};
`;

const Hovered = css`
  background: ${colors.danger};
`;

const Item = styled.li<{ isSelect: boolean }>`
  ${props => props.isSelect && Hovered}
`;
