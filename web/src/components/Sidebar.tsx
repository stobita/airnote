import React, { useState } from "react";
import styled from "styled-components";
import { Tag } from "../model/link";
import { ReactComponent as SettingIconImage } from "../assets/setting.svg";
import { Switch } from "./Switch";

interface Props {
  tags: Tag[];
  onClickTitle: () => void;
  onClickTag: (id: number) => void;
  isPanelView: boolean;
  setIsPanelView: () => void;
}

export const Sidebar = (props: Props) => {
  const [settingActive, setSettingActive] = useState(false);
  const onClickConfig = () => {
    setSettingActive(prev => !prev);
  };
  const handleOnClickTag = (e: React.MouseEvent<HTMLLIElement>) => {
    const id = e.currentTarget.dataset.id;
    if (id) {
      props.onClickTag(Number(id));
    }
  };
  return (
    <Wrapper>
      <Title onClick={props.onClickTitle}>AirNote</Title>
      <ul>
        {props.tags.map(tag => (
          <Item key={tag.id} data-id={tag.id} onClick={handleOnClickTag}>
            {tag.text}
          </Item>
        ))}
      </ul>
      <Bottom>
        {settingActive && (
          <Setting>
            <Field>
              <label>Panel View</label>
              <Switch
                value={props.isPanelView}
                onChange={props.setIsPanelView}
              />
            </Field>
          </Setting>
        )}
        <SettingIcon onClick={onClickConfig} />
      </Bottom>
    </Wrapper>
  );
};

const Setting = styled.div``;
const Field = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const Wrapper = styled.div`
  background: ${props => props.theme.bg};
  color: ${props => props.theme.text};
  border-right: 1px solid ${props => props.theme.border};
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  height: 100%;
  padding: 24px;
  box-sizing: border-box;
`;

const Title = styled.h2`
  cursor: pointer;
`;

const Item = styled.li`
  cursor: pointer;
  font-size: 1.1rem;
  margin: 8px 0;
`;

const Bottom = styled.div`
  margin-top: auto;
`;

const SettingIcon = styled(SettingIconImage)`
  cursor: pointer;
  height: 36px;
`;
