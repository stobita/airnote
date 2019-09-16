import React, { useState, useContext } from "react";
import styled from "styled-components";
import { ReactComponent as SettingIconImage } from "../assets/setting.svg";
import { Switch } from "./Switch";
import { ViewContext } from "../context/viewContext";
import { TagList } from "./TagList";

interface Props {
  onClickTitle: () => void;
}

export const Sidebar = (props: Props) => {
  const [settingActive, setSettingActive] = useState(false);
  const { themeName, setThemeName } = useContext(ViewContext);
  const isDarkTheme = themeName === "dark";
  const onClickConfig = () => {
    setSettingActive(prev => !prev);
  };
  const handleOnClickThemeSwitch = () => {
    setThemeName(prev => (prev === "light" ? "dark" : "light"));
  };
  return (
    <Wrapper>
      <Title onClick={props.onClickTitle}>AirNote</Title>
      <TagList></TagList>
      <Bottom>
        {settingActive && (
          <Setting>
            <Field>
              <label>Dark Theme</label>
              <Switch value={isDarkTheme} onChange={handleOnClickThemeSwitch} />
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
  background: ${props => props.theme.main};
  color: ${props => props.theme.solid};
  border-right: 1px solid ${props => props.theme.border};
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  height: 100%;
  flex: 1;
  padding: 16px 24px;
  box-sizing: border-box;
`;

const Title = styled.h2`
  cursor: pointer;
  margin: 8px 0;
`;

const Bottom = styled.div`
  margin-top: auto;
`;

const SettingIcon = styled(SettingIconImage)`
  fill: ${props => props.theme.solid};
  cursor: pointer;
  height: 36px;
`;
