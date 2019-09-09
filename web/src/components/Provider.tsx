import React, { ReactNode } from "react";
import { DataContextProvider } from "../DataContextProvider";

interface Props {
  children: ReactNode;
}

export default (props: Props) => {
  return <DataContextProvider>{props.children}</DataContextProvider>;
};
