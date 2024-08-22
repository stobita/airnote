import colors from "./colors";

export const themeNames = ["light", "dark"] as const;
export type ThemeType = typeof themeNames[number];

const lightTheme = {
  primary: colors.primary,
  lightPrimary: colors.lightPrimary,
  primaryInner: colors.white,
  main: colors.mainWhite,
  bg: colors.paleGray,
  solid: colors.mainBlack,
  weak: colors.lightGray,
  border: colors.borderGray,
  danger: colors.danger,
  dangerInner: colors.white
};

const darkTheme = {
  primary: colors.primary,
  lightPrimary: colors.lightPrimary,
  primaryInner: colors.white,
  main: colors.mainBlack,
  bg: colors.black,
  solid: colors.mainWhite,
  weak: colors.mainGray,
  border: colors.mainGray,
  danger: colors.danger,
  dangerInner: colors.white
};

export default {
  light: lightTheme,
  dark: darkTheme
};
