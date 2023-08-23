import React, { useState, useEffect, CSSProperties, useContext } from 'react';
import { Link } from 'react-router-dom';
import { UserContext } from '../context/AuthContext';
import { BiUserCircle } from 'react-icons/bi'
import { IoShareSocial } from 'react-icons/io5';

//style properties for the navbar
const navbarStyles: CSSProperties = {
    position: 'sticky',
    display: 'flex',
    top: 0,
    height: '100%',
    width: '100%',
    backgroundColor: '#fa4d6a',
    alignItems: 'center',
    justifyContent: 'space-around'
}

const logoStyles: CSSProperties = {
    position: 'absolute',
    display: 'flex',
    height: '80%',
    alignItems: 'center',
    justifyContent: 'center',
    fontSize: '300%',
    color: 'white',
    border: '2px black solid',
    borderRadius: '20%'
}


export const NavBar: React.FC = () => {

    const userContext = useContext(UserContext)
    console.log(userContext);


    return (
        <nav style={navbarStyles}>
            <span style={logoStyles}>
                <IoShareSocial />
            </span>
            <h1>Social Network</h1>
            <div>
                <Link to={userContext.user ? "/dashboard/user/" + userContext.user.displayName : ""}>
                    <BiUserCircle />
                </Link>
            </div>
        </nav>
    );

};