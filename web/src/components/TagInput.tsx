import React, { useState, useEffect } from "react";
import styled, { css } from "styled-components";
import { FieldItemBase } from "./FieldItemBase";
import { AutoSuggest, useSuggest } from "./AutoSuggest";
import { Tag } from "../model/link";

interface Props {
  name: string;
  placeholder: string;
  value: string[];
  tags: Tag[];
  onChange: (items: string[]) => void;
}

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
    selectedSuggestionIndex,
    suggestions,
    handleOnClickSuggest,
    handleOnMouseEnterSuggest,
    handleOnClickRecommentded
  } = useTagInput(props.value, props.onChange, props.tags);

  return (
    <Wrapper>
      <InputArea hasTag={items.length > 0}>
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
        <Input
          error={inputError}
          name={props.name}
          value={inputValue}
          placeholder={props.placeholder}
          onChange={handleOnChangeInput}
          onKeyDown={handleOnKeyDown}
          onBlur={handleOnBlurInput}
          hasTag={items.length > 0}
        />
      </InputArea>
      <AutoSuggest
        items={suggestions}
        inputValue={inputValue}
        onMouseEnterItem={handleOnMouseEnterSuggest}
        hoverIndex={selectedSuggestionIndex}
        onClickItem={handleOnClickSuggest}
      ></AutoSuggest>
      {props.tags.length > 0 && (
        <Recommended>
          <RecommendedTitle>Recommended:</RecommendedTitle>
          {props.tags
            .filter(v => !items.includes(v.text))
            .slice(0, 5)
            .map(v => (
              <RecommendItem
                key={v.id}
                data-text={v.text}
                onClick={handleOnClickRecommentded}
              >
                #{v.text}
              </RecommendItem>
            ))}
        </Recommended>
      )}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  width: 100%;
`;

const InputArea = styled.div<{ hasTag: boolean }>`
  ${FieldItemBase}
  border: 1px solid ${props => props.theme.border};
  height: auto;
  min-height: 36px;
  display: flex;
  flex-wrap: wrap;
  background: ${props => props.theme.bg};
  padding: 4px 8px;
  ${props =>
    props.hasTag &&
    css`
      padding-bottom: 0px;
    `}
`;

const Input = styled.input<{ hasTag: boolean; error: boolean }>`
  ${FieldItemBase};
  color: ${props => props.theme.solid};
  padding: 4px 0;
  margin: 0;
  height: auto;
  width: 50%;
  ${props =>
    props.error &&
    css`
      color: ${props => props.theme.danger};
    `}
  ${props =>
    props.hasTag &&
    css`
      margin-bottom: 4px;
    `}
`;

const Item = styled.li<{ willRemove: boolean }>`
  position: relative;
  color: ${props => props.theme.solid};
  border: 1px solid ${props => props.theme.solid};
  border-radius: 4px;
  margin-right: 4px;
  margin-bottom: 4px;
  box-sizing: border-box;
  padding: 4px 20px 4px 4px;
  ${props =>
    props.willRemove &&
    css`
      background: ${props => props.theme.weak};
      color: ${props => props.theme.solid};
      border: 1px solid ${props => props.theme.border};
    `}
`;

const Close = styled.span`
  position: absolute;
  cursor: pointer;
  top: 3px;
  right: 5px;
`;

const RecommendItem = styled.span`
  margin-right: 8px;
  cursor: pointer;
`;

const Recommended = styled.div`
  color: ${props => props.theme.solid};
  margin-top: 8px;
  display: flex;
  flex-wrap: wrap;
`;

const RecommendedTitle = styled.span`
  margin-right: 8px;
  font-weight: bold;
`;

const useTagInput = (
  value: string[],
  onChange: (item: string[]) => void,
  tags: Tag[]
) => {
  const [items, setItems] = useState(value);
  const [inputValue, setInputValue] = useState("");
  const [inputError, setInputError] = useState(false);
  const [willRemove, setWillRemove] = useState(false);

  const suggestionBase = tags.map(v => v.text).filter(v => !items.includes(v));

  const {
    suggestions,
    selectedSuggestion,
    selectedSuggestionIndex,
    setSelectedSuggestionIndex,
    suggestUp,
    suggestDown
  } = useSuggest(suggestionBase, inputValue);

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
      e.preventDefault();
      if (selectedSuggestion) {
        addItem(selectedSuggestion);
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
      suggestUp();
    } else if (e.keyCode === 40) {
      suggestDown();
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
    setSelectedSuggestionIndex(idx);
  };

  const handleOnClickRecommentded = (e: React.MouseEvent<HTMLElement>) => {
    const selectedText = e.currentTarget.dataset.text;
    if (selectedText) {
      addItem(selectedText);
    }
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
    selectedSuggestionIndex,
    suggestions,
    handleOnClickSuggest,
    handleOnMouseEnterSuggest,
    handleOnClickRecommentded
  };
};
