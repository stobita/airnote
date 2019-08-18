import React, { ReactNode } from "react";
import styled from "styled-components";
import { Transition } from "react-transition-group";
import { TransitionStatus } from "react-transition-group/Transition";
import colors from "../colors";
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
`;

const slideWidth = 500;

const Inner = styled.div<{ state: TransitionStatus }>`
  padding: 16px 24px;
  box-sizing: border-box;
  position: absolute;
  right: -${slideWidth}px;
  transition: 0.5s;
  width: ${slideWidth}px;
  transform: translateX(
    ${({ state }) => (state === "entered" ? -slideWidth : 0)}px
  );
  min-height: 100vh;
  background: ${colors.mainBlack};
`;
