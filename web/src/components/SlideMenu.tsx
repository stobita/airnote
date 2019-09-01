import React, { ReactNode } from "react";
import styled from "styled-components";
import { Transition } from "react-transition-group";
import { TransitionStatus } from "react-transition-group/Transition";
import { CloseButton } from "./CloseButton";

interface Props {
  children: ReactNode;
  open: boolean;
  onClose: () => void;
}

export const SlideMenu = (props: Props) => {
  return (
    <Transition
      in={props.open}
      timeout={{ appear: 0, exit: 500 }}
      unmountOnExit
    >
      {state => (
        <Inner state={state}>
          <Head>
            <CloseButton onClick={props.onClose} />
          </Head>
          {props.children}
        </Inner>
      )}
    </Transition>
  );
};

const Head = styled.div`
  display: flex;
  justify-content: flex-end;
  margin-bottom: 8px;
`;

const slideWidth = 500;

const Inner = styled.div<{ state: TransitionStatus }>`
  background: ${props => props.theme.bg};
  color: ${props => props.theme.text};
  border-left: 1px solid ${props => props.theme.border};
  padding: 16px 24px;
  box-sizing: border-box;
  position: fixed;
  overflow-y: auto;
  right: -${slideWidth}px;
  transition: 0.5s;
  width: ${slideWidth}px;
  transform: translateX(
    ${({ state }) => (state === "entered" ? -slideWidth : 0)}px
  );
  min-height: 100vh;
  height: 100%;
`;
