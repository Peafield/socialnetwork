import React, { ReactNode, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../context/AuthContext';
import { getCookie } from '../controllers/SetUserContextAndCookie';
import { useSetUserContextAndCookie } from '../controllers/SetUserContextAndCookie'

interface ProtectedRouteProps {
    element: ReactNode
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
    element
}) => {
    const navigate = useNavigate();
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const setUserContextAndCookie = useSetUserContextAndCookie();

    const checkUserToken = () => {
        const userToken = getCookie("sessionToken");
        if (!userToken || userToken === 'undefined') {
            setIsLoggedIn(false);
            return navigate('/signin');
        }
        setUserContextAndCookie(userToken)
        setIsLoggedIn(true);
    }
    useEffect(() => {
        checkUserToken();
    }, [isLoggedIn]);
    return (
        <React.Fragment>
            {
                isLoggedIn ? element : null
            }
        </React.Fragment>
    );
}

export default ProtectedRoute