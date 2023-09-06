import React, { useContext } from "react";
import { Link } from "react-router-dom";
import { UserContext } from "../context/AuthContext";
import { BiUserCircle } from "react-icons/bi";
import { IoShareSocial } from "react-icons/io5";
import { IoMdNotifications } from 'react-icons/io'
import LogoutButton from "./Auth/SignOut";
import { AiOutlinePlus } from "react-icons/ai";
import styles from "./Dashboard.module.css";

export const NavBar: React.FC = () => {
  const userContext = useContext(UserContext);

  return (
    <nav className={styles.navbar}>
      <div style={{ width: '50%' }}>
        <Link to={"/dashboard"}>
          <span className={styles.logo} style={{ marginLeft: '95.5%' }}>
            <IoShareSocial />
          </span>
        </Link>
      </div>
      <div className={styles.navbarActions}>
        <div
          className={styles.navbarbutton}>
          <IoMdNotifications />
        </div>
        <Link
          to={
            userContext.user
              ? "/dashboard/user/" + userContext.user.displayName
              : ""
          }
        >
          <div style={{ display: 'flex', alignItems: 'center' }}>
            <div>
              <BiUserCircle />{" "}
            </div>
            <div style={{ paddingBottom: '10%', marginLeft: '2%' }}>
              {userContext.user ? userContext.user.displayName : " "}
            </div>
          </div>
        </Link>
        <div>
          <Link to={"/dashboard/createpost"}>
            <AiOutlinePlus />
          </Link>
        </div>
        <div>
          <LogoutButton />
        </div>
      </div>
    </nav>
  );
};
