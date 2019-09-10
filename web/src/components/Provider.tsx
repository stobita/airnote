import React, { ReactNode } from "react";
import { DataContextProvider } from "../DataContextProvider";
import { ViewContextProvider } from "./ViewContextProvider";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  return (
    <DataContextProvider>
      <ViewContextProvider>{props.children}</ViewContextProvider>
    </DataContextProvider>
  );
};
