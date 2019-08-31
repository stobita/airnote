import React, { useState, useEffect } from "react";
import styled, { css } from "styled-components";
import colors from "../colors";
import { FieldItemBase } from "./FieldItemBase";

interface Props {
  name: string;
  placeholder: string;
  value: string[];
  onChange: (items: string[]) => void;
}

const useTagInput = (value: string[], onChange: (item: string[]) => void) => {
  const [items, setItems] = useState(value);
  const [inputValue, setInputValue] = useState("");
  const [inputError, setInputError] = useState(false);
  const [willRemove, setWillRemove] = useState(false);

  useEffect(() => {
    setInputError(false);
    setWillRemove(false);
  }, [inputValue, items]);

  useEffect(() => {
    onChange(items);
  }, [items, inputValue, onChange]);

  useEffect(() => {
    setItems(value);
  }, [value]);

  const addItem = () => {
    const value = inputValue.trim();
    if (items.find(v => v === value)) {
      setInputValue(value);
      setInputError(true);
      return;
    }
    if (value.length === 0) {
      setInputValue("");
      return;
    }
    setItems(prev => [...prev, value]);
    setInputValue("");
  };
  const handleOnChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputError(false);
    setInputValue(e.target.value);
  };
  const handleOnClickClose = (e: React.MouseEvent<HTMLElement>) => {
    setInputError(false);
    setItems(items.filter(v => v !== e.currentTarget.dataset.item));
  };
  const handleOnKeyDown = (e: React.KeyboardEvent) => {
    if (e.keyCode === 13) {
      addItem();
    } else if (e.keyCode === 8 && inputValue.length === 0 && items.length > 0) {
      if (willRemove) {
        setItems(prev => prev.slice(0, prev.length - 1));
      } else {
        setWillRemove(true);
      }
    }
  };
  const handleOnBlurInput = () => {
    addItem();
  };

  return {
    items,
    inputValue,
    inputError,
    willRemove,
    handleOnClickClose,
    handleOnBlurInput,
    handleOnChangeInput,
    handleOnKeyDown
  };
};

export const TagInput = (props: Props) => {
  const {
    items,
    inputValue,
    inputError,
    willRemove,
    handleOnBlurInput,
    handleOnChangeInput,
    handleOnKeyDown,
    handleOnClickClose
  } = useTagInput(props.value, props.onChange);

  return (
    <Wrapper>
      <List>
        {items.map((item, index) => (
          <Item
            key={item}
            willRemove={willRemove && index === items.length - 1}
          >
            <span>{item}</span>
            <Close data-item={item} onClick={handleOnClickClose}>
              ×
            </Close>
          </Item>
        ))}
      </List>
      <Input
        error={inputError}
        name={props.name}
        value={inputValue}
        placeholder={props.placeholder}
        onChange={handleOnChangeInput}
        onKeyDown={handleOnKeyDown}
        onBlur={handleOnBlurInput}
      />
    </Wrapper>
  );
};

const Wrapper = styled.label`
  ${FieldItemBase}
  display: flex;
  background: ${colors.mainWhite};
  padding: 4px 8px;
`;

const Input = styled.input<{ error: boolean }>`
  ${FieldItemBase}
  padding:4px 0;
  margin: 0;
  height: auto;
  ${props =>
    props.error &&
    css`
      color: ${colors.danger};
    `}
`;

const List = styled.ul`
  display: flex;
`;

const Item = styled.li<{ willRemove: boolean }>`
  position: relative;
  color: ${colors.mainGray};
  border: 1px solid ${colors.mainGray};
  border-radius: 4px;
  margin-right: 4px;
  box-sizing: border-box;
  padding: 4px 20px 4px 4px;
  ${props =>
    props.willRemove &&
    css`
      background: ${colors.thinGray};
      color: ${colors.mainWhite};
      border: 1px solid ${colors.thinGray};
    `}
`;

const Close = styled.span`
  position: absolute;
  cursor: pointer;
  top: 3px;
  right: 5px;
`;
