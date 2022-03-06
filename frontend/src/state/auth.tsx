import { atom, useRecoilState } from "recoil";
import { User } from "oidc-client";

export const authState = atom<User | null>({
  key: "auth",
  default: null,
});

export default authState;
