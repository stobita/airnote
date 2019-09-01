import React, { useCallback } from "react";
import { Link } from "../model/link";
import { LinkItem } from "./LinkItem";
import styled, { css } from "styled-components";
import { LinkItemPanel } from "./LinkItemPanel";

interface Props {
  items: Link[];
  onSelectItem: (l: Link) => void;
  onClickTag: (id: number) => void;
  isPanelView: boolean;
}

export const LinkIndex = (props: Props) => {
  const onClickItem = useCallback(
    (l: Link) => {
      props.onSelectItem(l);
    },
    [props]
  );

  return (
    <Wrapper>
      {props.items.map(link => (
        <Column isPanel={props.isPanelView} key={link.id}>
          {props.isPanelView ? (
            <LinkItemPanel
              item={link}
              onClick={onClickItem}
              onClickTag={props.onClickTag}
            />
          ) : (
            <LinkItem
              item={link}
              onClick={onClickItem}
              onClickTag={props.onClickTag}
            />
          )}
        </Column>
      ))}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-wrap: wrap;
  padding: 16px;
`;

const Column = styled.div<{ isPanel: boolean }>`
  display: flex;
  ${props =>
    props.isPanel
      ? css`
          flex: 0 1 20%;
          max-width: 20%;
        `
      : css`
          flex: 0 1 100%;
        `}
  box-sizing: border-box;
`;
