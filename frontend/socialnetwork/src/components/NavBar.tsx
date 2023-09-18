import React, { useContext, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { UserContext } from "../context/AuthContext";
import { BiUserCircle } from "react-icons/bi";
import { IoShareSocial } from "react-icons/io5";
import { HiUserGroup } from 'react-icons/hi'
import LogoutButton from "./Auth/SignOut";
import { AiOutlinePlus } from "react-icons/ai";
import styles from "./Dashboard.module.css";
import Notification from "./Notifications/Notifications";
import { LuMessagesSquare } from "react-icons/lu";
import Notifications from "./Notifications/Notifications";

interface NavBarProps {
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setSideModalDisplay: React.Dispatch<React.SetStateAction<string | null>>;
}

export const NavBar: React.FC<NavBarProps> = ({
  setIsModalOpen,
  setSideModalDisplay
}) => {
  const userContext = useContext(UserContext);
  const navigate = useNavigate();

  return (
    <nav className={styles.navbar}>
      <div style={{ width: '50%' }}>
        <span className={styles.logo} style={{ marginLeft: '95.5%' }} onClick={() => { navigate("/dashboard") }}>
          <IoShareSocial />
        </span>
      </div>
      <div className={styles.navbarActions}>
        <Link to={"/dashboard/groups"}>
          <span>
            <HiUserGroup />
          </span>
        </Link>
        <div
          className={styles.navbarbutton}
          style={{ marginBottom: '2.5%' }}
          onClick={() => { setIsModalOpen(true); setSideModalDisplay("chats") }}>
          <LuMessagesSquare />
        </div>
        <Notifications setIsModalOpen={setIsModalOpen} setSideModalDisplay={setSideModalDisplay} />
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
