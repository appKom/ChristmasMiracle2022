import { atom, useRecoilState } from "recoil";
import { BackendUser } from "../types/api";

export const userState = atom<BackendUser | null>({
  key: "user",
  default: null,
});

export const useUser = () => {
  const [user, setUser] = useRecoilState(userState);
  return { user, setUser };
};

export default userState;
