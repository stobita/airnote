import React, { useState, useContext } from "react";
import { Link } from "../model/link";
import styled from "styled-components";
import colors from "../colors";
import { LinkItemDetail } from "./LinkItemDetail";
import { LinkItemTags } from "./LinkItemTags";
import { ViewContext } from "../context/viewContext";
import { ReactComponent as EditIconImage } from "../assets/edit.svg";

interface Props {
  item: Link;
}

export const LinkItem = (props: Props) => {
  const { item } = props;
  const [expand, setExpand] = useState(false);
  const { setSlideTargetLinkId } = useContext(ViewContext);
  const handleOnClick = () => {
    setExpand(prev => !prev);
  };

  const handleOnClickEdit = (e: React.MouseEvent) => {
    e.stopPropagation();
    setSlideTargetLinkId(item.id);
  };

  return (
    <Wrapper onClick={handleOnClick}>
      <Title>{item.title && item.title}</Title>
      {expand && <LinkItemDetail item={item}></LinkItemDetail>}
      <Bottom>
        <LinkItemTags items={item.tags}></LinkItemTags>
        {expand && (
          <Operator>
            <EditIcon onClick={handleOnClickEdit} />
          </Operator>
        )}
      </Bottom>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background: ${props => props.theme.bg};
  border: 1px solid ${colors.borderGray};
  border-radius: 4px;
  box-sizing: border-box;
  margin: 8px;
  cursor: pointer;
  overflow: hidden;
  width: 100%;
  padding: 8px;
`;

const Bottom = styled.div`
  display: flex;
  align-items: center;
`;

const Title = styled.p`
  font-weight: bold;
  margin: 0;
  margin-bottom: 8px;
`;

const Operator = styled.div`
  flex: 2;
  display: flex;
  justify-content: flex-end;
`;

const EditIcon = styled(EditIconImage)`
  fill: ${props => props.theme.text};
  height: 24px;
`;
