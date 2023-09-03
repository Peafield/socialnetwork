import React, { useState, useEffect, CSSProperties, useContext } from 'react';
import { Link } from 'react-router-dom';
import { UserContext } from '../context/AuthContext';
import { BiUserCircle } from 'react-icons/bi'
import { IoShareSocial } from 'react-icons/io5';
import { MdOutlineLogout } from 'react-icons/md'
import LogoutButton from './Auth/SignOut';
import { AiOutlinePlus } from 'react-icons/ai';
import styles from './Dashboard.module.css'

export const NavBar: React.FC = () => {

    const userContext = useContext(UserContext)

    return (
        <nav className={styles.navbar}>
            <span className={styles.logo}>
                <IoShareSocial />
            </span>
            <div className={styles.navbarActions}>
                <div>
                    <Link to={userContext.user ? "/dashboard/user/" + userContext.user.displayName : ""}>
                        <BiUserCircle />
                    </Link>
                </div>
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