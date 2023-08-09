import { createContext } from "react";
import { User } from "../hooks/useUser";

interface AuthContext {
  user: User;
  setUser: (user: User) => void;
}

export const AuthContext = createContext<AuthContext>({
  user: {} as User,
  setUser: () => {},
});