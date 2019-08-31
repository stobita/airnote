import React, { useState, useEffect } from "react";
import styled, { css } from "styled-components";
import colors from "../colors";
import { FieldItemBase } from "./FieldItemBase";
import { AutoSuggest } from "./AutoSuggest";
import { Tag } from "../model/link";

interface Props {
  name: string;
  placeholder: string;
  value: string[];
  tags: Tag[];
  onChange: (items: string[]) => void;
}

const useTagInput = (
  value: string[],
  onChange: (item: string[]) => void,
  tags: Tag[]
) => {
  const [items, setItems] = useState(value);
  const [inputValue, setInputValue] = useState("");
  const [inputError, setInputError] = useState(false);
  const [willRemove, setWillRemove] = useState(false);

  const [selectSuggestion, setSelectSuggestion] = useState("");
  const [selectSuggestionIndex, setSelectSuggestionIndex] = useState(0);

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

  const suggestions = tags
    .filter(
      v =>
        inputValue.length > 0 &&
        v.text.toLowerCase().startsWith(inputValue.toLowerCase())
    )
    .map(v => v.text);

  useEffect(() => {
    setSelectSuggestion(suggestions[selectSuggestionIndex]);
  }, [selectSuggestionIndex, suggestions]);

  const addItemFromInput = () => {
    const value = inputValue.trim();
    addItem(value);
  };

  const addItem = (value: string) => {
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
      if (selectSuggestion) {
        addItem(selectSuggestion);
      } else {
        addItemFromInput();
      }
    } else if (e.keyCode === 8 && inputValue.length === 0 && items.length > 0) {
      if (willRemove) {
        setItems(prev => prev.slice(0, prev.length - 1));
      } else {
        setWillRemove(true);
      }
    } else if (e.keyCode === 38) {
      if (selectSuggestionIndex < 1) {
        setSelectSuggestionIndex(suggestions.length - 1);
      } else {
        setSelectSuggestionIndex(prev => --prev);
      }
    } else if (e.keyCode === 40) {
      if (selectSuggestionIndex > suggestions.length - 2) {
        setSelectSuggestionIndex(0);
      } else {
        setSelectSuggestionIndex(prev => ++prev);
      }
    }
  };

  const handleOnBlurInput = (e: React.FocusEvent<HTMLElement>) => {
    if (suggestions.length < 1) {
      addItemFromInput();
    }
  };

  const handleOnClickSuggest = (idx: number) => {
    addItem(suggestions[idx]);
  };

  const handleOnMouseEnterSuggest = (idx: number) => {
    setSelectSuggestionIndex(idx);
  };

  return {
    items,
    inputValue,
    inputError,
    willRemove,
    handleOnClickClose,
    handleOnBlurInput,
    handleOnChangeInput,
    handleOnKeyDown,
    setSelectSuggestionIndex,
    selectSuggestionIndex,
    suggestions,
    handleOnClickSuggest,
    handleOnMouseEnterSuggest
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
    handleOnClickClose,
    selectSuggestionIndex,
    suggestions,
    handleOnClickSuggest,
    handleOnMouseEnterSuggest
  } = useTagInput(props.value, props.onChange, props.tags);

  return (
    <Wrapper>
      <InputArea>
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
      </InputArea>
      <AutoSuggest
        items={suggestions}
        inputValue={inputValue}
        onMouseEnterItem={handleOnMouseEnterSuggest}
        hoverIndex={selectSuggestionIndex}
        onClickItem={handleOnClickSuggest}
      ></AutoSuggest>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  width: 100%;
`;

const InputArea = styled.label`
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
